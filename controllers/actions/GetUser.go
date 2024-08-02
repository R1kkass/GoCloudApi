package actions

import (
	"mypackages/db"
	Model "mypackages/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (Model.User, bool) {
	email:=ParseJWT(c.Request.Header["Access-Token"][0])

	var user Model.User;

	r := db.DB.Model(&user).First(&user, "email=?", email)
	
	if r.RowsAffected==0{
		c.JSON(401, gin.H{
			"message": "error",
		})
		return user, false
	}

	return user, true
}