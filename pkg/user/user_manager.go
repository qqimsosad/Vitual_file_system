package user

import (
	"fmt"
	"strings"
	"virtual-file-system/pkg/models"
)

var users = make(map[string]*models.User)

func Register(username string) error {
	username = strings.ToLower(username)
	if err := models.ValidateName(username); err != nil {
		return err
	}
	if _, exists := users[username]; exists {
		return fmt.Errorf("the %s has already existed", username)
	}
	users[username] = &models.User{Username: username, Folders: make(map[string]*models.Folder)}
	return nil
}

func GetUser(username string) (*models.User, error) {
	username = strings.ToLower(username)
	user, exists := users[username]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", username)
	}
	return user, nil
}
