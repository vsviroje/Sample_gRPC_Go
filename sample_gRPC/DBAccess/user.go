package DBAccess

import (
	"Sample_gRPC/models"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//init setUp

var userRecs []models.User
var userMapRecs map[int]models.User

func init() {
	userRecs = []models.User{
		{
			ID:      1,
			FName:   "ABC",
			City:    "Pune",
			Height:  5.7,
			Phone:   1234567890,
			Married: true,
		},
		{
			ID:      2,
			FName:   "PQR",
			City:    "Mumbai",
			Height:  6.1,
			Phone:   1234567890,
			Married: false,
		},
		{
			ID:      3,
			FName:   "XYZ",
			City:    "Pune",
			Height:  4.9,
			Phone:   1234567890,
			Married: true,
		},
		{
			ID:      4,
			FName:   "EFG",
			City:    "Banglore",
			Height:  5.5,
			Phone:   1234567890,
			Married: true,
		},
		{
			ID:      5,
			FName:   "FYI",
			City:    "Delhi",
			Height:  5.3,
			Phone:   1234567890,
			Married: false,
		},
	}
}

func GetUserList() []models.User {
	return userRecs
}

func GetUserMap() map[int]models.User {
	userMap := make(map[int]models.User)
	return userMap
}

func SetUserMap(userMap map[int]models.User) {
	userMapRecs = userMap
}

func GetUserByIdDB(id int) (*models.User, error) {
	v, isPresent := userMapRecs[id]

	if !isPresent {
		return nil, errors.New("no user found for " + fmt.Sprint(id))
	}

	return &v, nil
}

//required logical Func
