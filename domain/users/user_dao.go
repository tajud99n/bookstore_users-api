package users

import (
	"fmt"

	"github.com/tajud99n/bookstore_users-api/database/mysql/users"
	"github.com/tajud99n/bookstore_users-api/utils/date"
	"github.com/tajud99n/bookstore_users-api/utils/errors"
	mysqlUtils "github.com/tajud99n/bookstore_users-api/utils/mysql"
)

const (
	queryInsertUser = "INSERT INTO users(firstname, lastname, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, firstname, lastname, email, date_created FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlUtils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlUtils.ParseError(saveErr)
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := users.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		fmt.Println(getErr)
		return mysqlUtils.ParseError(getErr)
	}

	return nil
}
