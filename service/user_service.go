package service

import (
	"corelyst-api/model"
	"corelyst-api/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user model.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("username & password wajib diisi")
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("gagal hash password")
	}
	user.Password = string(hashedPassword)

	return repository.CreateUser(user)
}

func LoginUser(username, password string) (model.User, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return model.User{}, errors.New("user tidak ditemukan")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, errors.New("password salah")
	}
	return user, nil
}

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}
