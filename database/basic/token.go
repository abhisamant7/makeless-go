package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
)

func (database *Database) GetToken(token *go_saas_model.Token, value string) (*go_saas_model.Token, error) {
	return token, database.GetConnection().
		Preload("Team").
		Preload("User").
		Where("tokens.token = ?", value).
		First(&token).
		Error
}

func (database *Database) GetTokens(user *go_saas_model.User, tokens []*go_saas_model.Token) ([]*go_saas_model.Token, error) {
	return tokens, database.GetConnection().
		Select([]string{
			"tokens.id",
			"tokens.note",
			"tokens.user_id",
			"tokens.team_id",
			"CONCAT(REPEAT('X', CHAR_LENGTH(tokens.token) - 4),SUBSTRING(tokens.token, -4)) as token",
		}).
		Where("tokens.user_id = ? AND tokens.team_id IS NULL", user.GetId()).
		Order("tokens.id DESC").
		Find(&tokens).
		Error
}

func (database *Database) CreateToken(token *go_saas_model.Token) (*go_saas_model.Token, error) {
	return token, database.GetConnection().
		Create(&token).
		Error
}

func (database *Database) DeleteToken(token *go_saas_model.Token) error {
	return database.GetConnection().
		Unscoped().
		Where("tokens.id = ? AND tokens.user_id = ? AND tokens.team_id IS NULL", token.GetId(), token.GetUserId()).
		Delete(&token).
		Error
}
