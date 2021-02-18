package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index"`
}

//associar a funcao ao metodo. Olha o ponteiro * usado no User
//Funcao começando com Maiuscula é PUBLICA e minuscula Privada

func NewUser() *User {
	return &User{}
}
func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error during the pass generation:m %v", err)
		return err
	}
	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	if err != nil {
		log.Fatalf("Error during the user generation:m %v", err)
		return err
	}

	return nil
}

func (user *User) validate() error {
	return nil
}
