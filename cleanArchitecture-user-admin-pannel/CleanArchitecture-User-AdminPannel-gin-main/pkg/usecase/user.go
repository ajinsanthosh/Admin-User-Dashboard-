package usecase

import (
	"errors"
	interfaces "example/pkg/repository/interface"

	interfacesUseCase "example/pkg/usecase/interface"
	"example/pkg/utils/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) interfacesUseCase.UserUseCase {
	return &userUseCase{userRepo: repo}
}

// user signup

func (c *userUseCase) UseUserSignup(userData models.UserDetails) error {
	if userData.ConfirmPassword == userData.Password {

		HashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err, "problem at hashing signup")
		}

		userData.Password = string(HashedPassword)
		exist := c.userRepo.SaveUserData(userData)
		if exist != nil {
			fmt.Println(exist, "at in usecase exist")
			return exist
		}
	} else {
		return errors.New("confirm password is not match")
	}
	return nil
}

//user login

func (c *userUseCase) UseUserLogin(LoginData models.UserLoginDetails) error {
	LoginFeatchData, err := c.userRepo.GetUserData(LoginData)

	if err != nil {
		return errors.New("no user exist")
	} else {

		err := bcrypt.CompareHashAndPassword([]byte(LoginFeatchData.Password), []byte(LoginData.Password))
		if err != nil {
			return errors.New("password is not match")
		} else {
			return nil
		}

	}

}

// user home

func (c *userUseCase) UseUserName(UserId string) string {
	email := c.userRepo.RepoGetUserName(UserId)
	return email
}
