package users

import (
	"fmt"
	"github.com/TriTranDev/bookstore_users_api/datasources/mysql/users_db"
	"github.com/TriTranDev/bookstore_users_api/utils/date_utils"
	"github.com/TriTranDev/bookstore_users_api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreatd = result.DateCreatd
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %d already registered", user.Id))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.Id))
	}

	user.DateCreatd = date_utils.GetNowString()

	userDB[user.Id] = user
	return nil
}
