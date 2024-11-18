package main
//
//import (
//	"bytes"
//	"encoding/json"
//	"io/ioutil"
//	"log"
//	"net/http"
//
//	"github.com/google/uuid"
//	"github.com/spf13/afero"
//	"github.com/yandex/pandora/cli"
//	phttp "github.com/yandex/pandora/components/phttp/import"
//	"github.com/yandex/pandora/core"
//	"github.com/yandex/pandora/core/aggregator/netsample"
//	coreimport "github.com/yandex/pandora/core/import"
//	"github.com/yandex/pandora/core/register"
//)
//
//type respUsers struct {
//	Users []user `json:"user"`
//}
//
//type user struct {
//	UUID  string `json:"uuid"`
//	Login string `json:"login"`
//	Email string `json:"email"`
//}
//
//type Ammo struct {
//	Tag    string
//	Param1 string
//	Param2 string
//	Param3 string
//}
//
//type Sample struct {
//	URL              string
//	ShootTimeSeconds float64
//}
//
//type GunConfig struct {
//	Target string `validate:"required"` // Configuration will fail, without target defined
//}
//
//type Gun struct {
//	// Configured on construction.
//	client *http.Client
//	conf   GunConfig
//	// Configured on Bind, before shooting
//	aggr core.Aggregator // May be your custom Aggregator.
//	core.GunDeps
//}
//
//func NewGun(conf GunConfig) *Gun {
//	return &Gun{conf: conf}
//}
//
//func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
//	g.client = &http.Client{}
//	g.aggr = aggr
//	g.GunDeps = deps
//	return nil
//}
//
//func (g *Gun) Shoot(ammo core.Ammo) {
//	customAmmo := ammo.(*Ammo)
//	g.shoot(customAmmo)
//}
//
//func (g *Gun) case1_method(ammo *Ammo) int {
//	code := 0
//
//	id, _ := uuid.NewUUID()
//
//	b := map[string]any{
//		"login":    id.String(),
//		"password": id.String(),
//		"email":    id.String(),
//	}
//
//	data, _ := json.Marshal(b)
//
//	req, err := http.NewRequest(
//		http.MethodPost,
//		"http://localhost:18080/user/register",
//		bytes.NewReader(data),
//	)
//	if err != nil {
//		return 404
//	}
//
//	resp, err := g.client.Do(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	// Читаем ответ от сервера
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if body != nil {
//		code = 200
//	}
//
//	return code
//}
//
//func (g *Gun) shoot(ammo *Ammo) {
//	code := 0
//	sample := netsample.Acquire(ammo.Tag)
//
//	switch ammo.Tag {
//	case "/MyCase1":
//		code = g.case1_method(ammo)
//	default:
//		code = 404
//	}
//
//	defer func() {
//		sample.SetProtoCode(code)
//		g.aggr.Report(sample)
//	}()
//}
//
//func main() {
//	//debug.SetGCPercent(-1)
//	// Standard imports.
//	fs := afero.NewOsFs()
//	coreimport.Import(fs)
//	// May not be imported, if you don't need http guns and etc.
//	phttp.Import(fs)
//
//	// Custom imports. Integrate your custom types into configuration system.
//	coreimport.RegisterCustomJSONProvider("custom_provider", func() core.Ammo { return &Ammo{} })
//
//	register.Gun("my_custom_gun_name", NewGun, func() GunConfig {
//		return GunConfig{
//			Target: "localhost:18080",
//		}
//	})
//
//	cli.Run()
//}
