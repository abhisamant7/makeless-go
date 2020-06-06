package go_saas_basic_database

import "github.com/go-saas/go-saas/model"

func (database *Database) UpdatePassword(user *go_saas_model.User, password string) error {
	return database.GetConnection().
		Model(user).
		Update(map[string]interface{}{
			"password": password,
		}).
		Error
}