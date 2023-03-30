package app

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") //localhost:3000
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (a *Application) StartServer() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	//authorize
	r.POST("/login", a.Login)
	r.GET("/logout", a.Logout)
	r.POST("/signup", a.SignUp)

	//user
	r.POST("/project", a.CreateProject)
	r.GET("/upcoming", a.GetUpcomingNotifications)
	r.GET("/favorites", a.GetFavoriteProjects)
	r.GET("/projects", a.GetAllProjects)
	r.GET("/owned_projects", a.GetAllOwnedProjects)

	//project
	r.PUT("/:project", a.UpdateProject)
	r.DELETE("/:project", a.DeleteProject)
	r.GET("/:project/collaborators", a.GetCollaborators)
	r.GET("/:project/sections", a.GetAllSections)
	r.POST("/:project/collaborator", a.AddCollaborator)
	r.DELETE("/:project/collaborator", a.DeleteCollaborator)

	//section
	r.PUT("/:project/:section", a.UpdateSection)
	r.DELETE("/:project/:section", a.DeleteSection)
	r.POST("/:project/section", a.CreateSection)
	r.GET("/:project/:section/notifications", a.GetAllNotifications)
	r.GET("/:project/:section/:notification", a.GetNotification)

	//notification
	r.PUT("/:project/:section/:notification", a.UpdateNotification)
	r.DELETE("/:project/:section/:notification", a.DeleteNotification)
	r.POST("/:project/:section/notification", a.CreateNotification)

	//moderator
	r.GET("/undelivered_notifications", a.GetUndeliviredNotifications)
	r.Run()
}
