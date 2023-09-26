package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"log"
	valid "w8protector/validation"
)

func (u *User) ValidateUser() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(valid.RequiredIf(u.EncryptedPassword == "")), validation.Length(5, 64)),
	)
}

func (u *User) BeforeUserCreation() error {
	if len(u.Password) > 0 {
		enc, err := encryptPassword(u.Password)
		if err != nil {
			log.Fatalf("failed to encrypt user password, err:%s", err.Error())
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(passwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(passwd)) == nil
}

func encryptPassword(passwd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(passwd), 14)
	return string(b), err
}
