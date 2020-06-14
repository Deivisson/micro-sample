package models

import (
	"encoding/json"
	"html"
	"net/http"
	"strings"
	"time"

	// "github.com/deivisson/apstore/api/utils"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name     string         `gorm:"size:50;not null" json:"name"`
	LastName string         `gorm:"size:50;not null" json:"lastName"`
	Email    string         `gorm:"size:100;not null;unique" json:"email"`
	Password string         `gorm:"size:100;not null;" json:"password"`
	Settings postgres.Jsonb `gorm:"not null;default: '{}'" json:"settings"`
}

type UserErrors struct {
	Name     []string `json:"name,omitempty"`
	LastName []string `json:"lastName,omitempty"`
	Email    []string `json:"email,omitempty"`
	Password []string `json:"password,omitempty"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// FindAll return all users
func (u *User) FindAll(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// FindByID return the user by given id
func (u *User) FindByID(db *gorm.DB, id uint64) error {
	var err error
	err = db.Model(User{}).Where("id = ?", id).Take(&u).Error

	if err != nil && gorm.IsRecordNotFoundError(err) {
		u.Errors.RecordNotFound = err
	} else if err != nil {
		u.Errors.Exception = err
	}
	return err
}

// FindByEmail return the user by given email
func (u *User) FindByEmail(db *gorm.DB, email string) (finded bool, err error) {
	err = db.Model(User{}).Where("email = ?", email).Take(&u).Error

	if err != nil && gorm.IsRecordNotFoundError(err) {
		u.Errors.RecordNotFound = err
	} else if err != nil {
		u.Errors.Exception = err
	}
	finded = err == nil
	return
}

// Create an new user when signUp
func (u *User) Create(db *gorm.DB, params []byte) error {
	var err error
	if err = json.Unmarshal(params, &u); err != nil {
		u.Errors.Exception = err
		return err
	}

	u.prepare()
	if errors := u.validateSignUp(); errors != nil {
		u.Errors.Business = errors
		return err
	}

	if err = db.Create(&u).Error; err != nil {
		u.Errors.Exception = err
		return err
	}
	return nil
}

// UpdateUser the user
// func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {

// 	// To hash the password
// 	err := u.BeforeSave()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
// 		map[string]interface{}{
// 			"password":  u.Password,
// 			"nickname":  u.Name,
// 			"email":     u.Email,
// 			"update_at": time.Now(),
// 		},
// 	)
// 	if db.Error != nil {
// 		return &User{}, db.Error
// 	}
// 	// This is the display the updated user
// 	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }


// ValidateSignIn validate credentials of user
func (u *User) ValidateSignIn() interface{} {
	errors := UserErrors{}
	validatePresenceOf(u.Password, &errors.Password)
	validatePresenceOf(u.Email, &errors.Email)
	validateEmail(u.Email, &errors.Email)

	if len(errors.Password) > 0 || len(errors.Email) > 0 {
		return errors
	}
	return nil
}

func (u *User) prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) validateSignUp() interface{} {
	errors := UserErrors{}
	validates(u.Name, &errors.Name, validatesForStringFields(true, 3, 50)...)
	validates(u.LastName, &errors.LastName, validatesForStringFields(true, 3, 50)...)
	validates(u.Password, &errors.Password, validatesForStringFields(true, 6, 20)...)
	validates(u.Email, &errors.Email, validatesForStringFields(true, nil, 100)...)
	validateEmail(u.Email, &errors.Email)

	if len(errors.Name) > 0 ||
		len(errors.Password) > 0 ||
		len(errors.Email) > 0 ||
		len(errors.LastName) > 0 {
		return errors
	}
	return nil
}

func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
