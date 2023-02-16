package repository

import (
	"errors"
	"fmt"

	"github.com/Karibu/api-go-human/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetById(string) (model.User, error)
	GetAll() ([]model.User, error)
	Save(model.User) (model.User, error)
}

type SqlUserRepository struct {
	db *gorm.DB
}

func (r *SqlUserRepository) Save(m model.User) (model.User, error) {
	result := r.db.Create(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

func (r *SqlUserRepository) GetById(id string) (model.User, error) {
	user := model.User{}
	result := r.db.Find(&user, "id = ?", id)

	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, fmt.Errorf("no se encontraron usuarios con este ID: '%s'", id)
	}

	return user, nil
}

func (r *SqlUserRepository) GetAll() ([]model.User, error) {
	users := []model.User{}
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, errors.New("error al obtener lista de usuarios")
	}
	return users, nil
}

func NewUserRepository() (IUserRepository, error) {
	db, _ := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	return &SqlUserRepository{db: db}, nil
}
