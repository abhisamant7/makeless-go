package saas_model

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type Token struct {
	gorm.Model
	Token *string `gorm:"unique;not null" json:"token" binding:"required"`
	Read  *bool   `gorm:"not null" json:"read" binding:"required"`
	Write *bool   `gorm:"not null" json:"write" binding:"required"`

	UserId *uint `gorm:"not null" json:"userId"`
	User   *User `json:"user"`

	*sync.RWMutex `json:"-"`
}

func (token *Token) GetId() uint {
	token.RLock()
	defer token.RUnlock()

	return token.ID
}

func (token *Token) GetToken() *string {
	token.RLock()
	defer token.RUnlock()

	return token.Token
}

func (token *Token) GetUserId() *uint {
	token.RLock()
	defer token.RUnlock()

	return token.UserId
}
