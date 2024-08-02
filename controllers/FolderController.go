package Controller

import (
	Consts "mypackages/consts"
	"mypackages/controllers/actions"
	"mypackages/db"
	Model "mypackages/models"
	"mypackages/policy"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FolderRequest struct {
	NameFolder string `form:"name_folder" json:"name_folder" xml:"name_folder"  binding:"required"` 
	FolderID int `form:"folder_id" json:"folder_id" xml:"folder_id"` 
}

type FolderDelete struct {
	FolderID int `form:"folder_id" json:"folder_id" xml:"folder_id"  binding:"required"` 
}

type MoveFolderRequest struct {
	FolderID int `form:"folder_id" json:"folder_id" xml:"folder_id"  binding:"required"` 
	FolderToID int `form:"folder_to_id" json:"folder_to_id" xml:"folder_to_id"  binding:"required"` 
}

type ChangeFolderAccessRequest struct{
	FolderID int `form:"folder_id" json:"folder_id" xml:"folder_id" binding:"required"`
	AccessID int `form:"access_id" json:"access_id" xml:"access_id" bindings:"required"`
}

func CreateFolder(c *gin.Context){
	var jsonValid FolderRequest
	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}
	
	if jsonValid.FolderID!=0 && !policy.FolderPolicyID(c, strconv.Itoa(jsonValid.FolderID)){
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}

	user, _ := actions.GetUser(c)

	var folder Model.Folder;

	if jsonValid.FolderID!=0 {
		folder = Model.Folder{
			FolderID: jsonValid.FolderID,
			UserID: int(user.ID),
			NameFolder: jsonValid.NameFolder,
			AccessId: Consts.CLOSE,
		};
	} else{
		folder = Model.Folder{
			UserID: int(user.ID),
			NameFolder: jsonValid.NameFolder,
			AccessId: Consts.CLOSE,
		};
	}


	result := db.DB.Create(&folder)

	if result.RowsAffected==0{
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	os.Mkdir("files/"+strconv.Itoa(int(user.ID))+"/"+strconv.Itoa(int(folder.ID)), os.ModePerm)

	c.JSON(201, gin.H{
		"message": "success",
	})
}

func DeleteFolder(c *gin.Context){

	var folder Model.Folder;
	if !policy.FolderPolicyID(c, c.Param("id")){
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	user, _ := actions.GetUser(c)
	
	actions.RecursiveDeleteFiles(c.Param("id"), user)

	result := db.DB.Unscoped().Where("id = ? AND user_id=?", c.Param("id"), user.ID).Delete(&folder)

	if result.RowsAffected==0{
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}


	c.JSON(201, gin.H{
		"message": "success",
	})
}

func RenameFolder(c *gin.Context){
	var jsonValid RenameRequest
	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}
	
	if !policy.FolderPolicyID(c, c.Param("id")){
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}
	user, _ := actions.GetUser(c)
	
	var folder Model.Folder 

	actions.RecursiveDeleteFiles(c.Param("id"), user)

	result:=db.DB.Where("id = ? AND user_id=?", c.Param("id"), user.ID).First(&folder)
	if result.RowsAffected==0{
		c.JSON(404, gin.H{
			"message": "error",
		})
		return
	}
	db.DB.Model(&folder).Update("name_folder", jsonValid.Name)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func MoveFolder(c *gin.Context) {
	var jsonValid MoveFolderRequest

	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}

	if !policy.FolderPolicyID(c, strconv.Itoa(jsonValid.FolderToID)){
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}


	if !policy.FolderPolicyID(c, strconv.Itoa(jsonValid.FolderID)){
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}

	user, _ := actions.GetUser(c)

	var folder *Model.Folder

	result:=db.DB.Where("id = ? AND user_id=?", jsonValid.FolderID, user.ID).First(&folder)

	if result.RowsAffected==0{
		c.JSON(404, gin.H{
			"message": "error",
		})
		return
	}

	db.DB.Model(&folder).Update("folder_id", jsonValid.FolderToID)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func ChangeFolderAccess(c *gin.Context){
	var jsonValid ChangeFolderAccessRequest;

	if err := c.ShouldBind(&jsonValid); err != nil {
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}

	var folder Model.Folder

	user, _ := actions.GetUser(c)

	result:=db.DB.Where("id = ? AND user_id=?", jsonValid.FolderID, user.ID).First(&folder)

	if !policy.FolderPolicyID(c, strconv.Itoa(jsonValid.FolderID)) && jsonValid.FolderID!=0{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}

	if result.RowsAffected==0{
		c.JSON(404, gin.H{
			"message": "error",
		})
		return
	}

	db.DB.Model(&folder).Update("access_id", jsonValid.AccessID)

	c.JSON(200, gin.H{
		"message": "success",
	})
}