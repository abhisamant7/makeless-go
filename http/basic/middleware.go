package go_saas_http_basic

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/security"
	h "net/http"
	"strconv"
)

func (http *Http) CorsMiddleware(Origins []string, AllowHeaders []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = Origins
	config.AllowCredentials = true
	config.AddAllowHeaders(AllowHeaders...)

	return cors.New(config)
}

func (http *Http) TeamUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var teamUser bool
		var teamId int
		var userId = http.GetAuthenticator().GetAuthUserId(c)

		if c.GetHeader("Team") == "" {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(errors.New("no team header"), nil))
			return
		}

		if teamId, err = strconv.Atoi(c.GetHeader("Team")); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if teamUser, err = http.GetSecurity().IsTeamUser(http.GetSecurity().GetDatabase().GetConnection(), uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamUser {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(go_saas_security.NoTeamUserErr, nil))
			return
		}

		c.Next()
	}
}

func (http *Http) TeamRoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var teamRole bool
		var teamId int
		var userId = http.GetAuthenticator().GetAuthUserId(c)

		if c.GetHeader("Team") == "" {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(errors.New("no team header"), nil))
			return
		}

		if teamId, err = strconv.Atoi(c.GetHeader("Team")); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if teamRole, err = http.GetSecurity().IsTeamRole(http.GetSecurity().GetDatabase().GetConnection(), role, uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamRole {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(go_saas_security.NoTeamRoleError, nil))
			return
		}

		c.Next()
	}
}

func (http *Http) TeamCreatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var teamCreator bool
		var teamId int
		var userId = http.GetAuthenticator().GetAuthUserId(c)

		if c.GetHeader("Team") == "" {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(errors.New("no team header"), nil))
			return
		}

		if teamId, err = strconv.Atoi(c.GetHeader("Team")); err != nil {
			c.AbortWithStatusJSON(h.StatusBadRequest, http.Response(err, nil))
			return
		}

		if teamCreator, err = http.GetSecurity().IsTeamCreator(http.GetSecurity().GetDatabase().GetConnection(), uint(teamId), userId); err != nil {
			c.AbortWithStatusJSON(h.StatusInternalServerError, http.Response(err, nil))
			return
		}

		if !teamCreator {
			c.AbortWithStatusJSON(h.StatusUnauthorized, http.Response(go_saas_security.NoTeamCreatorError, nil))
			return
		}

		c.Next()
	}
}
