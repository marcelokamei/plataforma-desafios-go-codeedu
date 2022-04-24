package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/marcelokamei/plataforma-desafios-go-codeedu/application/repositories"
	"github.com/marcelokamei/plataforma-desafios-go-codeedu/domain"
	"github.com/marcelokamei/plataforma-desafios-go-codeedu/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Marcelo",
		Email:    "marcelokamei@hotmail.com",
		Password: "1234",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}

	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}
