package repository

import (
	"errors"

	interfaces "example/pkg/repository/interface"
	"example/pkg/utils/models"
	"fmt"

	"gorm.io/gorm"
)

type UserDataBase struct {
	DB *gorm.DB
}

func NewUserDataBase(db *gorm.DB) interfaces.UserRepository {
	return &UserDataBase{DB: db}
}

func (c *UserDataBase) SaveUserData(userData models.UserDetails) error {

	var name string

	query1 := "SELECT name FROM users WHERE email=$1"
	row := c.DB.Raw(query1, userData.Email).Row()
	err := row.Scan(&name)

	if err != nil {
		fmt.Println(err, "error at inserting of data to database `SaveUserData`")
	}

	if name != "" {
		fmt.Println("already account exist from gmail")
		return errors.New("email contain a account")
	} else {
		query := `INSERT INTO Users (name, email, phone, password) VALUES($1, $2, $3, $4)`
		result := c.DB.Exec(query, userData.Name, userData.Email, userData.Phone, userData.Password)

		if result != nil {
			fmt.Println(result, "error at inseting of data to database ``saveUserData")
		}
	}
	return nil
}

//

func (c *UserDataBase) GetUserData(LoginData models.UserLoginDetails) (models.UserFeatchData, error) {
	var UserFeatchDetails models.UserFeatchData
	query := `SELECT email,password FROM users WHERE email=$1`

	row := c.DB.Raw(query, LoginData.Email).Row()
	err := row.Scan(&UserFeatchDetails.Email, &UserFeatchDetails.Password)
	if err != nil {
		fmt.Println(err, "error at featching data from database `GetUserData`")
	}

	if UserFeatchDetails.Email == "" {
		return UserFeatchDetails, errors.New("no user")
	}

	UserFeatchDetails.Email = LoginData.Email
	return UserFeatchDetails, nil
}

func (c *UserDataBase) RepoGetUserName(UserId string) string {
	var name string
	query := `SELECT name FROM users WHERE email=$1`
	row := c.DB.Raw(query, UserId).Row()
	err := row.Scan(&name)
	if err != nil {
		fmt.Println(err, "error at featching  data from database `GetUserData`")
	}
	fmt.Println(name, "---------")
	return name
}
