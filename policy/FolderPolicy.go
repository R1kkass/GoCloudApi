package policy

import (
	"fmt"
	"mypackages/controllers/actions"
	"mypackages/db"
	Model "mypackages/models"

	"github.com/gin-gonic/gin"
)

func FolderPolicy(c *gin.Context, folder_id string) bool{
	
	var user Model.User;
	var folder Model.Folder;

	user, exist := actions.GetUser(c)

	if !exist {

		return false
	}

	r := db.DB.Model(&folder).First(&folder, "user_id=? AND folder_id=?", user.ID, folder_id)

	if r.RowsAffected==0{

		return false
	}

	return true
}

func FolderPolicyID(c *gin.Context, folder_id string) bool{
	
	var user Model.User;
	var folder Model.Folder;

	user, exist := actions.GetUser(c)

	if !exist {
		return false
	}
	fmt.Println(folder_id)
	r := db.DB.Model(&folder).First(&folder, "user_id=? AND id=?", user.ID, folder_id)

	if r.RowsAffected==0{
		return false
	}

	return true
}