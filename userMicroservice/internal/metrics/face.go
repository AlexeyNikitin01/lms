package metrics

import (
	"context"
)

type ITelemetry interface {
	IncSingIn(ctx context.Context)
	IncSingUp(ctx context.Context)
}
