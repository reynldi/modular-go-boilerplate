package seeds

import (
	"log"

	"github.com/google/uuid"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/entity"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

type seedAuthUser struct {
	db *gorm.DB
}

func SeedAuthUsers(db *gorm.DB) *seedAuthUser {
	return &seedAuthUser{db}
}

func (s *seedAuthUser) CreateOne() entity.AuthUser {
	user := entity.AuthUser{
		UserId:          uuid.New().String(),
		Email:           faker.Internet().Email(),
		IsEmailVerified: true,
		PasswordHash:    faker.RandomString(128),
		Status:          1,
	}

	err := s.db.Create(&user).Error

	if err != nil {
		log.Println(err.Error())
	}

	return user
}

func (s *seedAuthUser) CreateMany(count int) []entity.AuthUser {
	var users []entity.AuthUser

	for i := 1; i <= count; i++ {
		user := s.CreateOne()
		users = append(users, user)
	}

	return users
}

func (s *seedAuthUser) DeleteOne(data entity.AuthUser) {
	s.db.Where("email = ?", data.Email).Unscoped().Delete(&data)
}
