package grpc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"lms-user/internal/testutil"
)

func BenchmarkGet(b *testing.B) {
	c, ctx := testutil.Client(b)

	b.Run("register_gRPC", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resp, err := c.GetAllUser(ctx, &emptypb.Empty{})
			require.NoError(b, err)
			require.NotZero(b, resp.Users)
		}
	})
}
