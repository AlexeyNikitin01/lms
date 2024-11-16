package grpc_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/AlexeyNikitin01/lms-user/internal/ports/grpc"
	"github.com/AlexeyNikitin01/lms-user/internal/testutil"
)

func BenchmarkRegisterBenchmark(b *testing.B) {
	c, ctx := testutil.Client(b)

	b.Run("register_gRPC", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id, _ := uuid.NewUUID()
			resp, err := c.RegisterUser(ctx, &grpc.UserRegisterRequest{
				Login:    id.String(),
				Password: id.String(),
				Email:    id.String(),
			})
			require.NoError(b, err)
			require.NotZero(b, resp.Uuid)
		}
	})
}
