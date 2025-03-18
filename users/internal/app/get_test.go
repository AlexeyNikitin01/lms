package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"lms-user/internal/app"
	"lms-user/internal/repository/pg/entity"
	"lms-user/mocks"
)

func TestUserGet(t *testing.T) {

	type args struct {
		userUUID string
	}

	tests := []struct {
		name      string
		args      args
		mockCalls func(mocks *mocks.MockIUserPostgres, prms args)
	}{
		{
			name: "success",
			mockCalls: func(mocks *mocks.MockIUserPostgres, prms args) {
				mocks.EXPECT().
					GetUserDB(gomock.Any(), prms.userUUID).
					Return(&entity.User{
						ID:   prms.userUUID,
						Name: "alex",
					}, nil)
			},
			args: args{
				userUUID: "testUser",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockIUserPostgres(ctrl)
			cloud := mocks.NewMockIFace(ctrl)
			metric := mocks.NewMockITelemetry(ctrl)

			tt.mockCalls(db, tt.args)

			a := app.CreateAppUser(db, cloud, metric)

			user, err := a.GetUser(context.Background(), tt.args.userUUID)
			if err != nil {
				t.Errorf("GetUser() error = %v, wantErr %v", err, nil)
			}

			fmt.Println(user)
		})
	}

}
