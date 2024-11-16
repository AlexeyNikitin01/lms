package httpgin_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/lms-user/internal/testutil"
)

type respRegister struct {
	User user `json:"user"`
}

type user struct {
	UUID  string `json:"uuid"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func BenchmarkRegister(b *testing.B) {
	tc := testutil.CreateTestServer()

	b.Run("register", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data := createDataRegister()

			req, err := http.NewRequest(
				http.MethodPost,
				tc.BaseURL+"/user/register",
				bytes.NewReader(data),
			)
			require.NoError(b, err)

			req.Header.Add("Content-Type", "application/json")

			var resp respRegister
			err = tc.GetResponse(req, &resp)
			require.NoError(b, err)
			require.NotZero(b, resp.User.UUID)
		}
	})
}

func createDataRegister() []byte {
	id, _ := uuid.NewUUID()

	body := map[string]any{
		"login":    id.String(),
		"password": id.String(),
		"email":    id.String(),
	}

	data, _ := json.Marshal(body)

	return data
}
