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
	r.POST("/favorite", a.AddFavorite)
	r.GET("/favorite/:project_id", a.GetFavoriteProject)
	r.DELETE("/favorite/:project_id", a.DeleteFavorite)
	r.GET("/projects", a.GetAllProjects)
	r.GET("/projects/owned", a.GetAllOwnedProjects)
	r.PUT("/email", a.ChangeEmail)
	r.GET("/projects/latest", a.LastThreeProjects)

	//project
	r.PUT("/project/:project", a.UpdateProject)
	r.DELETE("/project/:project_id", a.DeleteProject)
	r.GET("/project/:project_id/collaborators", a.GetCollaborators)
	r.GET("/project/:project_id/sections", a.GetAllSections)
	r.POST("/project/:project_id/collaborator", a.AddCollaborator)
	r.DELETE("/project/:project_id/collaborator", a.DeleteCollaborator)

	//section
	r.PUT("/project/section/:section_id", a.UpdateSection)
	r.DELETE("/project/section/:section_id", a.DeleteSection)
	r.POST("/project/:project_id/section", a.CreateSection)
	r.GET("/project/section/:section_id/notifications", a.GetAllNotifications)
	r.GET("/project/section/notification/:notification_id", a.GetNotification)

	//notification
	r.PUT("/project/section/notification/:notification_id", a.UpdateNotification)
	r.DELETE("/project/section/notification/:notification_id", a.DeleteNotification)
	r.POST("/project/section/:section_id/notification", a.CreateNotification)

	//moderator
	r.GET("/undelivered_notifications", a.GetUndeliviredNotifications)
	r.Run()
}
