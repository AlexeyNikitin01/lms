package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/afero"
	"github.com/yandex/pandora/cli"
	phttp "github.com/yandex/pandora/components/phttp/import"
	"github.com/yandex/pandora/core"
	"github.com/yandex/pandora/core/aggregator/netsample"
	coreimport "github.com/yandex/pandora/core/import"
	"github.com/yandex/pandora/core/register"
	connect "google.golang.org/grpc"

	"lms-user/internal/ports/grpc"
)

type Ammo struct {
	Tag    string
	Param1 string
	Param2 string
	Param3 string
}

type Sample struct {
	URL              string
	ShootTimeSeconds float64
}

type GunConfig struct {
	Target string `validate:"required"` // Configuration will fail, without target defined
}

type Gun struct {
	// Configured on construction.
	client connect.ClientConnInterface
	conf   GunConfig
	// Configured on Bind, before shooting
	aggr core.Aggregator // May be your custom Aggregator.
	core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	var c connect.ClientConnInterface

	conn, err := connect.Dial(g.conf.Target, connect.WithInsecure())
	if err != nil {
		log.Fatalf("FATAL: %s", err)
	}
	c = conn
	g.client = c
	g.aggr = aggr
	g.GunDeps = deps
	return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo)
	g.shoot(customAmmo)
}

func (g *Gun) case1_method(client grpc.UserServiceClient, _ *Ammo) int {
	code := 0

	id, _ := uuid.NewUUID()
	out, err := client.RegisterUser(context.Background(), &grpc.UserRegisterRequest{
		Login:    id.String(),
		Password: id.String(),
		Email:    id.String(),
	})
	if err != nil {
		code = 404
		log.Println("FATAL: ", err)
	}

	if out != nil {
		code = 200
	}

	return code
}

func (g *Gun) shoot(ammo *Ammo) {
	code := 0
	sample := netsample.Acquire(ammo.Tag)

	client := grpc.NewUserServiceClient(g.client)

	switch ammo.Tag {
	case "/MyCase1":
		code = g.case1_method(client, ammo)
	default:
		code = 404
	}

	defer func() {
		sample.SetProtoCode(code)
		g.aggr.Report(sample)
	}()
}

func main() {
	//debug.SetGCPercent(-1)
	// Standard imports.
	fs := afero.NewOsFs()
	coreimport.Import(fs)
	// May not be imported, if you don't need http guns and etc.
	phttp.Import(fs)

	// Custom imports. Integrate your custom types into configuration system.
	coreimport.RegisterCustomJSONProvider("custom_provider", func() core.Ammo { return &Ammo{} })

	register.Gun("my_custom_gun_name", NewGun, func() GunConfig {
		return GunConfig{
			Target: "localhost:50054",
		}
	})

	cli.Run()
}
