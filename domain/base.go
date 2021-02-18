package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID       string    `json:"id" gorm:"type:uuid;primary_key"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	DeleteAt time.Time `json:"deleted_at" sql:"index"`
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
