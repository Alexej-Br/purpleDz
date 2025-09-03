package auth

import (
	"go-adv/4-order-api/pkg/db"
	"log"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB *db.DB
}

func NewUserRepository(database *db.DB) *UserRepository {
	return &UserRepository{DB: database}
}

func (user *UserRepository) FindByPhone(phone string) (string, error) {
	var u User
	tx := user.DB.First(&u, "phone=?", phone)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return "", tx.Error
	}
	return u.Phone, nil
}

func (user *UserRepository) FindBySession(session string) (string, error) {
	var u User
	tx := user.DB.First(&u, "session_id=?", session)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return "", tx.Error
	}
	return u.Phone, nil
}

func (user *UserRepository) NewUser(new *User) error {
	tx := user.DB.Create(&new)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return tx.Error
	}
	return nil
}

func (user *UserRepository) Update(u *User) error {
	result := user.DB.Where("phone=?", u.Phone).Clauses(clause.Returning{}).Updates(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
