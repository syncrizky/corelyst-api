package repository

import (
	"corelyst-api/config"
	"corelyst-api/model"
)

func CreateUser(user model.User) error {
	_, err := config.DB.Exec("INSERT INTO users(username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

func GetUserByUsername(username string) (model.User, error) {
	row := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username=?", username)
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func GetAllUsers() ([]model.User, error) {
	rows, err := config.DB.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Username, &u.Password)
		users = append(users, u)
	}
	return users, nil
}
