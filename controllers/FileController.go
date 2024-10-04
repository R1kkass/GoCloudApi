package Controller

import (
	"fmt"
	"io"
	"mime/multipart"
	Consts "mypackages/consts"
	"mypackages/controllers/actions"
	"mypackages/db"
	Model "mypackages/models"
	"mypackages/policy"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileRequest struct{
	FileName string `form:"file_name" json:"file_name" xml:"file_name"  binding:"required"` 
	FolderID int `form:"folder_id" json:"folder_id" xml:"folder_id"` 
	File *multipart.FileHeader `form:"file" json:"file" xml:"file"  binding:"required"` 
}

type RenameRequest struct{
	Name string `form:"name" json:"name" xml:"name"  binding:"required"` 
}

type MoveFileRequest struct{
	FileID string `form:"file_id" json:"file_id" xml:"file_id"  binding:"required"` 
	FolderID int `form:"folder_id" json:"folder_id" xml:"folder_id"` 
}


type FileDeleteRequest struct{
	FileID string `form:"file_id" json:"file_id" xml:"file_id"  binding:"required"`
}

type ChangeFileAccessRequest struct{
	FileID string `form:"file_id" json:"file_id" xml:"file_id" binding:"required"`
	FolderID int `form:"folder_id" json:"folder_id" xml:"folder_id" binding:"required"`
	AccessID string `form:"access_id" json:"access_id" xml:"access_id" bindings:"required"`
}


func FileCreate(c *gin.Context) {
	var pathFileFolder, _ = os.LookupEnv("PATH_FILES")
	var jsonValid FileRequest
	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}

	exist:=policy.FreeStorage(c)
	
	if !exist {
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}

	if(jsonValid.FolderID!=0){

		exist = policy.FileCreate(c, strconv.Itoa(jsonValid.FolderID))

		if !exist {
			c.JSON(403, gin.H{
				"message": "error",
			})
			return
		}
	}

	user, exist := actions.GetUser(c)

	if !exist {
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}
	files, _ := c.FormFile("file")

	
	var file Model.File;
	filesNameHash := uuid.New()
	filesNameHashFull := filesNameHash.String() + path.Ext(files.Filename)

	if jsonValid.FolderID!=0 {
		file = Model.File{
			FileName: files.Filename,
			UserRelation: Model.UserRelation{
				UserID: int(user.ID),
			},
			FolderRelation: Model.FolderRelation{
				FolderID: jsonValid.FolderID,
			},
			Size: int(files.Size),
			FileNameHash: filesNameHashFull,
			AccessId: Consts.CLOSE,
		};
	} else{
		file = Model.File{
			FileName: files.Filename,
			UserRelation: Model.UserRelation{
				UserID: int(user.ID),
			},
			Size: int(files.Size),
			FileNameHash: filesNameHashFull,
			AccessId: Consts.CLOSE,
		};
	}

	r := db.DB.Create(&file)

	if r.RowsAffected==0 {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}
	
	f, _ := files.Open()
	var path string = pathFileFolder+strconv.Itoa(int(user.ID))+"/"+filesNameHashFull

	if jsonValid.FolderID!=0{
		path = pathFileFolder+strconv.Itoa(int(user.ID))+"/"+strconv.Itoa(jsonValid.FolderID)+"/"+filesNameHashFull
	}

	defer f.Close()

	dst, errOs := os.Create(path)
	_, errIo := io.Copy(dst, f)

	if errIo!=nil || errOs!=nil{
		fmt.Println(errIo.Error(), errOs.Error())
		c.JSON(500, gin.H{
			"message": "Файл не сохранён",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "success",
	})

	defer dst.Close()
}

func DeleteFile(c *gin.Context) {
	
	policy.DeleteFile(c, c.Param("id"))
	
	var file *Model.File;

	r := db.DB.First(&file, c.Param("id"))

	if r.RowsAffected==0{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}
	user, _ := actions.GetUser(c)
	
	var path = getFilePath(file)

	errOs := os.Remove(path)

	if errOs!=nil{
		c.JSON(500, gin.H{
			"message": "Файл не удалён",
		})
		return
	}
	
	db.DB.Unscoped().Delete(&file, "id=? AND user_id=?", c.Param("id"), user.ID)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func GetAll(c *gin.Context){
	user, exist := actions.GetUser(c)
	
	if !exist{
		c.JSON(404, gin.H{
			"message": "error",
		})
	}
	var files []Model.File
	var folders []Model.Folder

	r:=db.DB.Preload("Folder").Preload("User").Where("user_id=?", user.ID)
	rFolder:=db.DB.Preload("User").Where("user_id=?", user.ID)
	
	if(c.Param("id")!=""){
		r.Where("folder_id=?", c.Param("id"))
		rFolder.Where("folder_id=?", c.Param("id"))	
	} else {
		r.Where("folder_id is null")
		rFolder.Where("folder_id is null")
	}
	r.Find(&files)
	rFolder.Find(&folders)

	c.JSON(200, gin.H{
		"folders": folders,
		"files": files,
	})
}

func OpenFile(c *gin.Context) {
	user, exist := actions.GetUser(c)
	var pathFileFolder, _ = os.LookupEnv("PATH_FILES")
	
	if !exist{
		c.JSON(404, gin.H{
			"message": "error",
		})
	}
	var file Model.File;

	r:=db.DB.Preload("Folder").Preload("User").Where("user_id=? and id=?", user.ID, c.Param("id")).First(&file)
	if r.RowsAffected==0{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}

	if file.FolderID==0{

		c.File(pathFileFolder+strconv.Itoa(file.UserID)+"/"+file.FileNameHash);	
		return
	}

	c.File(pathFileFolder+strconv.Itoa(file.UserID)+"/"+strconv.Itoa(file.FolderID)+"/"+file.FileNameHash);	
}

func GetSpace(c *gin.Context) {
	user, exist := actions.GetUser(c)
	
	if !exist{
		c.JSON(404, gin.H{
			"message": "error",
		})
	}
	var space int

	r:=db.DB.Table("files").Where("user_id=?", user.ID).Select("sum(size) as space").Scan(&space)
	
	if r.RowsAffected==0{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"space": space,
	})
}


func RenameFile(c *gin.Context){
	var jsonValid RenameRequest
	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}
	
	if !policy.FilePolicyID(c, c.Param("id")){
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}
	user, _ := actions.GetUser(c)
	
	var folder Model.File 

	result:=db.DB.Where("id = ? AND user_id=?", c.Param("id"), user.ID).First(&folder)
	if result.RowsAffected==0{
		c.JSON(404, gin.H{
			"message": "error",
		})
		return
	}
	db.DB.Model(&folder).Update("file_name", jsonValid.Name)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func MoveFile(c *gin.Context) {
	var jsonValid MoveFileRequest

	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}

	if !policy.FilePolicyID(c, jsonValid.FileID){
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}

	user, _ := actions.GetUser(c)

	if !policy.FolderPolicyID(c, strconv.Itoa(jsonValid.FolderID)) && jsonValid.FolderID!=0{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}
	var file *Model.File

	result:=db.DB.Where("id = ? AND user_id=?", jsonValid.FileID, user.ID).First(&file)

	if result.RowsAffected==0{
		c.JSON(404, gin.H{
			"message": "error",
		})
		return
	}
	var pathFileFolder, _ = os.LookupEnv("PATH_FILES")

	var oldPath = getFilePath(file)
	var err error;
	if jsonValid.FolderID!=0{
		err = os.Rename(oldPath, pathFileFolder+strconv.Itoa(file.UserID)+"/"+strconv.Itoa(jsonValid.FolderID)+"/"+file.FileNameHash)
	} else {
		err = os.Rename(oldPath, pathFileFolder+strconv.Itoa(file.UserID)+"/"+file.FileNameHash)
	}
	if err!=nil {		
		c.JSON(500, gin.H{
			"message": "success",
		})
		return
	}
	
	if jsonValid.FolderID!=0 {
		db.DB.Model(&file).Update("folder_id", jsonValid.FolderID)
	} else {
		db.DB.Model(&file).Update("folder_id", nil)
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func ChangeFileAccess(c *gin.Context) {
	var jsonValid ChangeFileAccessRequest;

	if err := c.ShouldBind(&jsonValid); err != nil {
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}

	user, _ := actions.GetUser(c)

	if !policy.FilePolicyID(c, jsonValid.FileID) && jsonValid.FileID!=""{
		c.JSON(403, gin.H{
			"message": "error",
		})
		return
	}

	var file Model.File

	result:=db.DB.Where("id = ? AND user_id=?", jsonValid.FileID, user.ID).First(&file)

	if result.RowsAffected==0{
		c.JSON(404, gin.H{
			"message": "error",
		})
		return
	}

	db.DB.Model(&file).Update("access_id", jsonValid.AccessID)

	c.JSON(200, gin.H{
		"message": "success",
	})
}



func getFilePath(file *Model.File) string{
	var pathFileFolder, _ = os.LookupEnv("PATH_FILES")

	var path string
	if file.FolderID==0{
		path = pathFileFolder+strconv.Itoa(file.UserID)+"/"+file.FileNameHash;	
	} else{
		path = pathFileFolder+strconv.Itoa(file.UserID)+"/"+strconv.Itoa(file.FolderID)+"/"+file.FileNameHash;	
	}
	return path;
}