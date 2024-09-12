package main

import (
	"log"
	Controller "mypackages/controllers"
	"mypackages/db"
	"mypackages/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Message struct {
	Email string `json:"email"`
	Status string `json:"status"`
}



func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

const (
	defaultName = "world"
)

func main() {
	db.ConnectDatabase()
	db.Migration()
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))


	authorized := r.Group("/")

	authorized.Use(middleware.VerifyJWT())

	{
		folder := authorized.Group("folder")
		folder.POST("/create", Controller.CreateFolder)
		folder.DELETE("/delete/:id", Controller.DeleteFolder)
		folder.PATCH("/update/:id", Controller.RenameFolder)
		folder.PATCH("/move", Controller.MoveFolder)
		folder.PATCH("/changeaccess", Controller.ChangeFolderAccess)
	}

	{
		file := authorized.Group("file")
		file.POST("/create", Controller.FileCreate)
		file.DELETE("/delete/:id", Controller.DeleteFile)
		file.PATCH("/move", Controller.MoveFile)
		file.PATCH("/changeaccess", Controller.ChangeFileAccess)
		file.PATCH("/update/:id", Controller.RenameFile)
	}

	{
		get := authorized.Group("get")	
		get.GET("/:id", Controller.GetAll)
		get.GET("/", Controller.GetAll)
		get.GET("/open/:id", Controller.OpenFile)
		get.GET("/space", Controller.GetSpace)
	}

	r.POST("/login", Controller.Login)
	r.POST("/registration", Controller.Registration)
	
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
