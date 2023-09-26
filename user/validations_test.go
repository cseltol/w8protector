package user_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"w8protector/user"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *user.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *user.User {
				return testUser()
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *user.User {
				u := testUser()
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *user.User {
				u := testUser()
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "with encrypted password",
			u: func() *user.User {
				u := testUser()
				u.Password = ""
				u.EncryptedPassword = "encrypted_password"

				return u
			},
			isValid: true,
		},
		{
			name: "invalid email",
			u: func() *user.User {
				u := testUser()
				u.Email = "invalid email"

				return u
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *user.User {
				u := testUser()
				u.Password = "short"

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().ValidateUser())
			} else {
				assert.Error(t, tc.u().ValidateUser())
			}
		})
	}
}

func TestUser_BeforeUserCreation(t *testing.T) {
	u := testUser()

	assert.NoError(t, u.BeforeUserCreation())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func testUser() *user.User {
	return &user.User{
		Email:    "user@example.com",
		Password: "password",
	}
}
