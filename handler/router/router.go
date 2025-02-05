package router

import (
	"goCrud/di"

	"github.com/gin-gonic/gin"
)

func Setup(service *di.CrudService) *gin.Engine {
	r := gin.Default()

	//Register Database connection (will be asked promptly)
	r.Use(func(c *gin.Context) {
		c.Set("db", service.DB)
		c.Next()
	})

	//public endpoint (generally no auth)
	publicRoutes(r, service)

	//Privante endpoint (generally with auth, or is not publicly shared, grouped with versionings)
	privateRouteV1(r, service)

	return r
}

func publicRoutes(r *gin.Engine, service *di.CrudService) {
	//to be added with PING and HEALTHCHECK ROUTE
}

func privateRouteV1(r *gin.Engine, service *di.CrudService) {

	auth := r.Group("/v1")
	//auth = r.group (url/v1) //you can use { } to explicitly tell which group a group of route belongs to
	auth.GET("/users", service.CrudHandler.GetAll)
	auth.POST("/users/create", service.CrudHandler.CreateUser)
	auth.POST("/users/edit", service.CrudHandler.UpdateUser)
	auth.DELETE("/users/delete/:id", service.CrudHandler.DeleteUser)
}
