package UsersTable

import (
	"../../Model"
	"../../DataBase"
	// "gorm.io/gorm"
	"errors"
)

func Create(user *Model.User) (*Model.User, error) {
	db := DataBase.New()

	if len(user.Username) < 3 {
		return nil, errors.New("too short username")
	}

	if len(user.Name) < 3 {
		return nil, errors.New("too short name")
	}

	db.Create(&user)

	return user, nil
}

func Save(user *Model.User) (*Model.User, error) {
	db := DataBase.New()

	if len(user.Username) < 3 {
		return nil, errors.New("too short username")
	}

	if len(user.Name) < 3 {
		return nil, errors.New("too short name")
	}

	db.Save(&user)

	return user, nil
}