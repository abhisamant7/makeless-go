package go_saas_security

import (
	"github.com/go-saas/go-saas/database"
	"github.com/go-saas/go-saas/model"
)

type Security interface {
	GetDatabase() go_saas_database.Database
	GenerateToken(length int) (string, error)
	UserExists(field string, value string) (bool, error)
	Login(field string, value string, password string) (*go_saas_model.User, error)
	Register(user *go_saas_model.User) (*go_saas_model.User, error)
	EncryptPassword(password string) (string, error)
	ComparePassword(userPassword string, password string) error
	IsTeamUser(teamId uint, userId uint) (bool, error)
	IsTeamRole(role string, teamId uint, userId uint) (bool, error)
	IsTeamCreator(teamId uint, userId uint) (bool, error)
	IsModelUser(userId uint, model interface{}) (bool, error)
	IsTeamInvitation(teamInvitation *go_saas_model.TeamInvitation) (bool, error)
}
