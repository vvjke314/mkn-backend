package app

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //localhost:3000
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
	r.POST("/api/login", a.Login)   //+
	r.GET("/api/logout", a.Logout)  //+
	r.POST("/api/signup", a.SignUp) //+

	r.GET("/api/projects", a.GetAllProjects) //+

	authorized := r.Group("/")

	authorized.Use(a.UserIdentity)
	{
		//user
		authorized.POST("/api/project", a.CreateProject)                 //+
		authorized.GET("/api/upcoming", a.GetUpcomingNotifications)      //-
		authorized.GET("/api/favorites", a.GetFavoriteProjects)          //+
		authorized.POST("/api/favorite", a.AddFavorite)                  //+
		authorized.DELETE("/api/favorite/:project_id", a.DeleteFavorite) //+
		authorized.GET("/api/projects/owned", a.GetAllOwnedProjects)     //+
		authorized.PUT("/api/email", a.ChangeEmail)                      //+
		authorized.GET("/api/projects/latest", a.LastSixProjects)        //+

		//project
		authorized.PUT("/api/project/:project_id", a.UpdateProject)                      //+
		authorized.DELETE("/api/project/:project_id", a.DeleteProject)                   //+
		authorized.GET("/api/project/:project_id/collaborators", a.GetCollaborators)     //+
		authorized.GET("/api/project/:project_id/sections", a.GetAllSections)            //+
		authorized.POST("/api/project/:project_id/collaborator", a.AddCollaborator)      //+
		authorized.DELETE("/api/project/:project_id/collaborator", a.DeleteCollaborator) //+

		//section
		authorized.PUT("/api/project/section/:section_id", a.UpdateSection)                     //+
		authorized.DELETE("/api/project/section/:section_id", a.DeleteSection)                  //+
		authorized.POST("/api/project/:project_id/section", a.CreateSection)                    //+
		authorized.GET("/api/project/section/:section_id/notifications", a.GetAllNotifications) //+
		authorized.GET("/api/project/section/notification/:notification_id", a.GetNotification) //+

		//notification
		authorized.PUT("/api/project/section/notification/:notification_id", a.UpdateNotification)        //+
		authorized.DELETE("/api/project/section/notification/:notification_id", a.DeleteNotification)     //+
		authorized.POST("/api/project/section/:section_id/notification", a.CreateNotification)            //+
		authorized.PUT("/api/project/section/notification/resend/:notification_id", a.ResendNotification) //+

		//moderator
		authorized.GET("/api/undelivered_notifications", a.GetUndeliviredNotifications) //-
	}

	r.Run()
}
