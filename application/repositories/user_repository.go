package repositories

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/marcelokamei/plataforma-desafios-go-codeedu/domain"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()
	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
	}

	err = repo.Db.Create(user).Error
	if err != nil {
		log.Fatalf("error to persist user: %v", err)
	}
	return user, nil
}
