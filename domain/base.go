package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID       string     `json:"id" gorm:"type:uuid;primary_key"`
	CreateAt time.Timer `json:"created_at" gorm:"type:datetime"`
	UpdateAt time.Timer `json:"updated_at" gorm:"type:datetime"`
	DeleteAt time.Timer `json:"deleted_at" gorm:"type:datetime"; sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
	}
	err := scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
	}

	return nil
}
