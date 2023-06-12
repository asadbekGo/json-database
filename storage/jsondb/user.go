package jsondb

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"app/models"

	"github.com/google/uuid"
)

type UserRepo struct {
	fileName string
	file     *os.File
}

func NewUserRepo(fileName string, file *os.File) *UserRepo {
	return &UserRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *UserRepo) Create(req *models.CreateUser) (*models.User, error) {

	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	var user = &models.User{
		Id:        id.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	users = append(users, user)

	body, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return user, nil
}
