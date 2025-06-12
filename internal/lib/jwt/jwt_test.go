package jwt_test

import (
	"testing"
	"time"

	"github.com/Braendie/sso/internal/domain/models"
	"github.com/Braendie/sso/internal/lib/jwt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewToken(t *testing.T) {
	testCases := []struct{
		name     string
		user     models.User
		app      models.App
		duration time.Duration
	}{
		{
			name: "Valid",
			user: models.User{
				ID: 1,
				Email: "test@mail.ru",
				PassHash: getPassHash(),
			},
			app: models.App{
				ID: 1,
				Name: "test",
				Secret: "testSecret",
			},
			duration: time.Hour,
		},
		{
			name: "empty user",
			user: models.User{},
			app: models.App{
				ID: 1,
				Name: "test",
				Secret: "testSecret",
			},
			duration: time.Hour,
		},
		{
			name: "empty app",
			user: models.User{
				ID: 1,
				Email: "test@mail.ru",
				PassHash: getPassHash(),
			},
			app: models.App{},
			duration: time.Hour,
		},
		{
			name: "So low time",
			user: models.User{
				ID: 1,
				Email: "test@mail.ru",
				PassHash: getPassHash(),
			},
			app: models.App{
				ID: 1,
				Name: "test",
				Secret: "testSecret",
			},
			duration: time.Microsecond,
		},
		{
			name: "Both empty",
			user: models.User{},
			app: models.App{},
			duration: time.Microsecond,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokenString, err := jwt.NewToken(tc.user, tc.app, tc.duration)
			assert.NoError(t, err)
			assert.NotEmpty(t, tokenString)
		})
	}
}


func getPassHash() []byte {
	pass := gofakeit.Password(true, true, true, true, false, 6)
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return passHash
}
