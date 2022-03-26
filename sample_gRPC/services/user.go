package services

import (
	dbAccess "Sample_gRPC/DBAccess"
	"Sample_gRPC/models"
	"errors"
	"fmt"
)

func GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("user id is zero")
	}
	return dbAccess.GetUserByIdDB(id)
}

func GetUserByIDs(ids []int32) (resp []models.User, err error) {

	if len(ids) == 0 {
		return nil, errors.New("user ids len is zero")
	}

	for i := range ids {
		user, err := dbAccess.GetUserByIdDB(int(ids[i]))
		if err != nil {
			fmt.Println("No user found")
			continue
		}
		resp = append(resp, *user)
	}

	return
}
