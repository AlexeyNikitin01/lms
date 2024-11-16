package app

import (
	"context"
	"time"

	"github.com/pascaldekloe/jwt"
	"github.com/thanhpk/randstr"

	"github.com/lms-user/internal/repository/pg/entity"
)

func (a appUser) RefreshToken(tokenUser *entity.Token) (string, error) {
	refresh := randstr.String(30)

	tokenUser.Refresh = refresh

	return refresh, nil
}

func (a appUser) AccessToken(user *entity.User, tokenUser *entity.Token) (string, error) {
	now := time.Now()
	accessExpires := now.Add(time.Hour * 1000)

	claims := jwt.Claims{
		Registered: jwt.Registered{
			Issued:  jwt.NewNumericTime(now.Truncate(time.Second)),
			Expires: jwt.NewNumericTime(accessExpires.Truncate(time.Second)),
			Subject: tokenUser.ID,
		},
		Set: map[string]interface{}{
			"user_id": user.ID,
		},
	}

	token, err := claims.HMACSign(jwt.HS512, []byte("00000000"))
	if err != nil {
		return "", err
	}

	tokenUser.Token = string(token)
	tokenUser.ExpiresAt = accessExpires

	return string(token), nil
}

func (a appUser) ParseToken(ctx context.Context, token string) (*entity.User, *entity.Token, error) {
	clm, err := jwt.HMACCheck([]byte(token), []byte("00000000"))
	if err != nil {
		return nil, nil, err
	}

	if !clm.Valid(time.Now()) {
		return nil, nil, err
	}

	tokenUserID := clm.Subject

	user, tokenUser, err := a.repo.GetUserByTokenDB(ctx, tokenUserID)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenUser, nil
}
