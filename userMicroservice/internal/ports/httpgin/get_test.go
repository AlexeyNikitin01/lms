package httpgin_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AlexeyNikitin01/lms-user/internal/testutil"
)

type respUsers struct {
	Users []user `json:"user"`
}

func BenchmarkGet(b *testing.B) {
	tc := testutil.CreateTestServer()

	b.Run("get all user", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(
				http.MethodPost,
				tc.BaseURL+"/user/get-users",
				nil,
			)
			require.NoError(b, err)

			req.Header.Add("Content-Type", "application/json")

			var resp respUsers
			err = tc.GetResponse(req, &resp)
			require.NoError(b, err)
			require.NotZero(b, resp.Users)
		}
	})
}
