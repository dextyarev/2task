package repositories

import (
	"2task/internal/database/models"

	"gorm.io/gorm"
)

type User = models.User

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id int, user User) (User, error)
	DeleteUserByID(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepository) UpdateUserByID(id int, user User) (User, error) {
	var res *gorm.DB

	switch {
	case user.Email != "" && user.Password != "":
		res = r.db.Model(&user).Select("Email", "Password").Updates(User{Email: user.Email, Password: user.Password})
	case user.Email != "":
		res = r.db.Model(&user).Select("Email").Updates(User{Email: user.Email})
	case user.Password != "":
		res = r.db.Model(&user).Select("Password").Updates(User{Password: user.Password})
	}

	if err := res.Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(id int) error {
	res := r.db.Delete(&User{}, id)

	if err := res.Error; err != nil {
		return err
	}
	return nil
}
