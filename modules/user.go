package modules

import (
	"github.com/containerops/sails/models"
)

func SaveUser(username, passwd, email string) (string, error) {
	return "", nil
}

func GetUser(username, passwd string) (models.User, error) {

	return models.User{}, nil
}
