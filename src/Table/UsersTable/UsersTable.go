package UsersTable

import (
	"../../Model"
	"../../DataBase"
	// "gorm.io/gorm"
	"errors"
	"regexp"
)

func isEmail(email string) bool {
	r := regexp.MustCompile(`^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$`)
	return r.MatchString(email)
}

func Create(user *Model.User) (*Model.User, error) {
	db := DataBase.New()

	if len(user.Username) < 3 {
		return nil, errors.New("too short username")
	}

	if len(user.Name) < 3 {
		return nil, errors.New("too short name")
	}

	if !isEmail(user.Email) {
		return nil, errors.New("invalid email")
	}

	user.Status = 1
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

	if !isEmail(user.Email) {
		return nil, errors.New("invalid email")
	}

	db.Save(&user)

	return user, nil
}