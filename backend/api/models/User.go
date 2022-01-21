package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint32 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Firstname string `gorm:"size:255" json:"firstName"`
	Lastname  string `gorm:"size:255" json:"lastName"`
	Position  string `gorm:"size:100;" json:"position"`
	Nation    string `gorm:"size:100;" json:"nation:abbrName"`
	Club      string `gorm:"size:100;" json:"club:abbrName"`
	Page      string `gorm:"size:100;" json:"page"`
}

func (u *User) Prepare() {
	u.ID = 0
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	u.Position = html.EscapeString(strings.TrimSpace(u.Position))
	u.Nation = html.EscapeString(strings.TrimSpace(u.Nation))
	u.Club = html.EscapeString(strings.TrimSpace(u.Club))
	u.Page = html.EscapeString(strings.TrimSpace(u.Page))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Firstname == "" {
			return errors.New("Required Firstname")
		}
		if u.Lastname == "" {
			return errors.New("Required Lastname")
		}
		if u.Position == "" {
			return errors.New("Required Position")
		}
		if u.Nation == "" {
			return errors.New("Required Nation")
		}
		if u.Club == "" {
			return errors.New("Required Club")
		}
		if u.Page == "" {
			return errors.New("Required Page")
		}

		return nil

	default:
		if u.Firstname == "" {
			return errors.New("Required Firstname")
		}
		if u.Lastname == "" {
			return errors.New("Required Lastname")
		}
		if u.Position == "" {
			return errors.New("Required Position")
		}
		if u.Nation == "" {
			return errors.New("Required Nation")
		}
		if u.Club == "" {
			return errors.New("Required Club")
		}
		if u.Page == "" {
			return errors.New("Required Page")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllplayers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
func (u *User) List(db *gorm.DB, pagination Pagination) (*Pagination, error) {
	users := []User{}
	db.Scopes(paginate(users, &pagination, db)).Find(&users)
	pagination.Rows = users

	return &pagination, nil
}
