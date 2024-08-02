package actions

import (
	"log"
	"mypackages/db"
	Model "mypackages/models"
	"os"
	"strconv"
)

func RecursiveDeleteFiles(folder_id string, user Model.User) {
	var folderNext []Model.Folder;

	log.Println("files/"+strconv.Itoa(int(user.ID))+"/"+folder_id)
	if folder_id!=""{
		path := "files/"+strconv.Itoa(int(user.ID))+"/"+folder_id
		os.RemoveAll(path)
	}

	r := db.DB.Where("folder_id=?", folder_id).Find(&folderNext)
	if r.RowsAffected!=0{
		for _, element := range folderNext{
			RecursiveDeleteFiles(strconv.Itoa(int(element.ID)), user)
		}
	}
}