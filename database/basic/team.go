package go_saas_database_basic

import (
	"github.com/go-saas/go-saas/model"
	"github.com/jinzhu/gorm"
)

// CreateTeam creates team and retrieves the new team with all their users
func (database *Database) CreateTeam(connection *gorm.DB, team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, connection.
		Create(team).
		Preload("TeamUsers.Team").
		Preload("TeamUsers.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.email")
		}).
		Find(team).
		Error
}

// AddTeamInvitations appends team invitations to a given team
func (database *Database) AddTeamInvitations(connection *gorm.DB, team *go_saas_model.Team, teamInvitations []*go_saas_model.TeamInvitation) (*go_saas_model.Team, error) {
	return team, connection.
		Model(team).
		Association("TeamInvitations").
		Append(teamInvitations).
		Error
}

// DeleteTeamUsers deletes own teamUser
func (database *Database) DeleteTeamUser(connection *gorm.DB, user *go_saas_model.User, team *go_saas_model.Team) error {
	return connection.
		Exec("DELETE FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ?", team.GetId(), user.GetId()).
		Error
}

// DeleteTeam deletes team if user is team creator
func (database *Database) DeleteTeam(connection *gorm.DB, user *go_saas_model.User, team *go_saas_model.Team) error {
	return connection.
		Exec("DELETE FROM teams WHERE teams.id = ? AND teams.user_id = ?", team.GetId(), user.GetId()).
		Error
}

// IsTeamUser checks if user is part of team
func (database *Database) IsTeamUser(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

// IsTeamRole checks if user is part of team and has given role
func (database *Database) IsTeamRole(connection *gorm.DB, role string, team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? AND team_users.role = ? LIMIT 1", team.GetId(), user.GetId(), role).
		Count(&count).
		Error
}

// IsTeamCreator checks if user is team creator
func (database *Database) IsTeamCreator(connection *gorm.DB, team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, connection.
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}
