package storage

import "github.com/Dorrrke/shusi_api/internal/models"

type UserStorage struct {
	users map[string]models.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{users: make(map[string]models.User)}
}

func (u *UserStorage) AddUser(user models.User) {
	u.users[user.ID] = user
}
