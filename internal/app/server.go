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
	r.POST("/login", a.Login)   //+
	r.GET("/logout", a.Logout)  //+
	r.POST("/signup", a.SignUp) //+

	r.GET("/projects", a.GetAllProjects) //+

	authorized := r.Group("/")

	authorized.Use(a.UserIdentity)
	{
		//user
		authorized.POST("/project", a.CreateProject)                  //+
		authorized.GET("/upcoming", a.GetUpcomingNotifications)       //-
		authorized.GET("/favorites", a.GetFavoriteProjects)           //+
		authorized.POST("/favorite", a.AddFavorite)                   //+
		authorized.GET("/favorite/:project_id", a.GetFavoriteProject) //+
		authorized.DELETE("/favorite/:project_id", a.DeleteFavorite)  //+
		authorized.GET("/projects/owned", a.GetAllOwnedProjects)      //+
		authorized.PUT("/email", a.ChangeEmail)                       //+
		authorized.GET("/projects/latest", a.LastSixProjects)         //+

		//project
		authorized.PUT("/project/:project_id", a.UpdateProject)                      //+
		authorized.DELETE("/project/:project_id", a.DeleteProject)                   //+
		authorized.GET("/project/:project_id/collaborators", a.GetCollaborators)     //+
		authorized.GET("/project/:project_id/sections", a.GetAllSections)            //+
		authorized.POST("/project/:project_id/collaborator", a.AddCollaborator)      //+
		authorized.DELETE("/project/:project_id/collaborator", a.DeleteCollaborator) //+

		//section
		authorized.PUT("/project/section/:section_id", a.UpdateSection)                     //+
		authorized.DELETE("/project/section/:section_id", a.DeleteSection)                  //+
		authorized.POST("/project/:project_id/section", a.CreateSection)                    //+
		authorized.GET("/project/section/:section_id/notifications", a.GetAllNotifications) //+
		authorized.GET("/project/section/notification/:notification_id", a.GetNotification) //+

		//notification
		authorized.PUT("/project/section/notification/:notification_id", a.UpdateNotification)        //+
		authorized.DELETE("/project/section/notification/:notification_id", a.DeleteNotification)     //+
		authorized.POST("/project/section/:section_id/notification", a.CreateNotification)            //+
		authorized.PUT("/project/section/notification/resend/:notification_id", a.ResendNotification) //+

		//moderator
		authorized.GET("/undelivered_notifications", a.GetUndeliviredNotifications) //-
	}

	r.Run()
}
