package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/raismaulana/blogP/application/apperror"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type User struct {
	ID          int64          `gorm:"primary_key:auto_increment;column:id_user"` //
	Username    string         `gorm:"type:varchar(12) unique not null"`          //
	Name        string         `gorm:"type:varchar(20) not null"`                 //
	Email       string         `gorm:"type:varchar(45) unique not null"`          //
	Password    string         `gorm:"type:varchar(255) not null"`                //
	City        string         `gorm:"type:varchar(50) not null"`                 //
	Country     string         `gorm:"type:varchar(50) not null"`                 //
	Birthday    time.Time      `gorm:"type:date not null"`                        //
	WebProfile  null.String    `gorm:"type:text null"`                            //
	ActivatedAt null.Time      `gorm:"default:null"`                              //
	CreatedAt   time.Time      ``                                                 //
	UpdatedAt   time.Time      ``                                                 //
	DeletedAt   gorm.DeletedAt `gorm:"index"`                                     //
}

type UserRequest struct {
	Username   string      `` //
	Name       string      `` //
	Email      string      `` //
	Password   string      `` //
	City       string      `` //
	Country    string      `` //
	Birthday   time.Time   `` //
	WebProfile null.String `` //
}

type UserUpdateRequest struct {
	Name       string      `` //
	City       string      `` //
	Country    string      `` //
	Birthday   time.Time   `` //
	WebProfile null.String `` //
}

func NewUser(req UserRequest) (*User, error) {

	//validate
	if strings.TrimSpace(req.Username) == "" {
		return nil, apperror.UsernameMustNotEmpty
	}
	if strings.TrimSpace(req.Name) == "" {
		return nil, apperror.NameMustNotEmpty
	}
	if strings.TrimSpace(req.Email) == "" {
		return nil, apperror.EmailMustNotEmpty
	}
	if strings.TrimSpace(req.Password) == "" {
		return nil, apperror.PasswordMustNotEmpty
	}
	if strings.TrimSpace(req.City) == "" {
		return nil, apperror.CityMustNotEmpty
	}
	if strings.TrimSpace(req.Country) == "" {
		return nil, apperror.CountryMustNotEmpty
	}
	if strings.TrimSpace(req.Birthday.String()) == "" {
		return nil, apperror.BirthdayMustNotEmpty
	}

	var obj User

	obj = User{
		Username:    req.Username,
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		City:        req.City,
		Country:     req.Country,
		Birthday:    req.Birthday,
		WebProfile:  req.WebProfile,
		ActivatedAt: null.Time{},
	}

	return &obj, nil
}

func (r *User) UpdateUser(req UserUpdateRequest) error {

	//validate
	if strings.TrimSpace(req.Name) == "" {
		return apperror.NameMustNotEmpty
	}
	if strings.TrimSpace(req.City) == "" {
		return apperror.CityMustNotEmpty
	}
	if strings.TrimSpace(req.Country) == "" {
		return apperror.CountryMustNotEmpty
	}
	if strings.TrimSpace(req.Birthday.String()) == "" {
		return apperror.BirthdayMustNotEmpty
	}

	r.Name = req.Name
	r.City = req.City
	r.Country = req.Country
	r.Birthday = req.Birthday
	r.WebProfile = req.WebProfile

	return nil
}

func (r *User) ValidateActivation(email string, activationCode string) error {
	if r.ActivatedAt.Valid {
		return apperror.UserIsAlreadyActivated
	}

	if r.Email != email {
		return apperror.InvalidEmail
	}

	if activationCode == "" {
		return apperror.InvalidActivationCode
	}

	return nil
}

func (r *User) RDBKeyActivationUser() string {
	return fmt.Sprintf("users:%v:%s:activation-code", r.ID, r.Email)
}
