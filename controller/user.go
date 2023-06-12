package controller

import (
	"app/models"
	"errors"
	"log"
)

func (c *Controller) UserCreate(req *models.CreateUser) (*models.User, error) {

	log.Printf("User create req: %+v\n", req)

	resp, err := c.Strg.User().Create(req)
	if err != nil {
		log.Printf("error while user create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}
