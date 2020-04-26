package saas_api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (api *Api) cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	log.Fatalf("%+v", api.getOrigins())
	config.AllowOrigins = api.getOrigins()
	config.AllowCredentials = true
	config.AddAllowHeaders("TEAM")

	return cors.New(config)
}
