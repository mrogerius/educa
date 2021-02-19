package main

import (
	"fmt"
	"log"

	"github.com/mrogerius/educa/application/repositories"
	"github.com/mrogerius/educa/domain"
	"github.com/mrogerius/educa/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Roger",
		Email:    "mrogerius@hotmail.com",
		Password: "123",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatalf("erro")
	}

	fmt.Println(result.Name)
}
