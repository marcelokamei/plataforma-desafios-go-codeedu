package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt time.Time `gorm:"type:datetime" sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())
	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
	}
	return nil
}
