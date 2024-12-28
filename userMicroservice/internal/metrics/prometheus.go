package metrics

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type Telemetry struct {
	*UserTelemetry
}

func NewUserOpenTelemetryMetric() (ITelemetry, error) {
	metricExporter, err := prometheus.New()
	if err != nil {
		return nil, errors.Wrap(err, "don`t create prometheus")
	}

	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(metricExporter))
	otel.SetMeterProvider(meterProvider)

	meter := otel.Meter("user")

	u, err := newUserTelemetry(meter)
	if err != nil {
		return nil, errors.Wrap(err, "don`t create metric user telemetry")
	}

	return &Telemetry{
		UserTelemetry: u,
	}, nil
}

type UserTelemetry struct {
	SingIn metric.Int64Counter
	SingUp metric.Int64Counter
}

func newUserTelemetry(meter metric.Meter) (*UserTelemetry, error) {
	var (
		err error
		ut  = &UserTelemetry{}
	)

	ut.SingIn, err = meter.Int64Counter(
		"user_singIn",
		metric.WithDescription("Количество посещений пользователей"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "не создан счётчик для посещений")
	}

	ut.SingUp, err = meter.Int64Counter(
		"user_singUp",
		metric.WithDescription("Количество уникальных пользователей"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "не создан счётчик для уникальных пользователей")
	}

	return ut, nil
}

func (b *UserTelemetry) IncSingIn(ctx context.Context) {
	b.SingIn.Add(ctx, 1)
	log.Println("add sing in")
}

func (b *UserTelemetry) IncSingUp(ctx context.Context) {
	b.SingUp.Add(ctx, 1)
	log.Println("add sing up")
}
