package storage

import "app/models"

type StorageI interface {
	User() UserRepoI
}

type UserRepoI interface {
	Create(req *models.CreateUser) (*models.User, error)
}
