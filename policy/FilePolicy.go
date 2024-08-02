package policy

import (
	"mypackages/controllers/actions"
	"mypackages/db"
	Model "mypackages/models"

	"github.com/gin-gonic/gin"
)

func FileCreate(c *gin.Context, folder_id string) bool {
	var user Model.User;
	var file Model.Folder;

	user, exist := actions.GetUser(c)

	if !exist {
		c.JSON(403, gin.H{
			"message": "error",
		})
		return false
	}

	if len(folder_id) > 0 {
		r := db.DB.Model(&file).First(&file, "user_id=? AND id=?", user.ID, folder_id)

		if r.RowsAffected==0{
			c.JSON(403, gin.H{
				"message": "error",
			})
			return false
		}
	}
	return true
}

func FilePolicyID(c *gin.Context, file_id string) bool{
	var user Model.User;
	var file Model.File;

	user, exist := actions.GetUser(c)

	if !exist {
		c.JSON(403, gin.H{
			"message": "error",
		})
		return false
	}
	r := db.DB.Model(&file).First(&file, "user_id=? AND id=?", user.ID, file_id)


	if r.RowsAffected==0{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return false
	}
	
	return true
}

func FreeStorage(c *gin.Context) bool {
	var user Model.User;
	var file Model.File;

	user, exist := actions.GetUser(c)

	if !exist {
		c.JSON(403, gin.H{
			"message": "error",
		})
		return false
	}

	db.DB.Model(&file).Select("sum(size) as size").Where("user_id=?", user.ID).Group("user_id").Scan(&file)
	filesSize, _ := c.FormFile("file")
	
	if file.Size+int(filesSize.Size)> 536870912*2{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return false
	}
	return true
}

func DeleteFile(c *gin.Context, file_id string) bool{
	var user Model.User;
	var file Model.File;

	user, exist := actions.GetUser(c)

	if !exist {
		c.JSON(403, gin.H{
			"message": "error",
		})
		return false

	}

	r := db.DB.Model(&file).First(&file, "user_id=? AND id=?", user.ID, file_id)
	
	if r.RowsAffected==0{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return false
	}
	return true

}