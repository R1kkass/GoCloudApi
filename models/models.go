package Model

import (
	"time"

	"gorm.io/gorm"
)

type DefaultModel struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Folder struct {
	DefaultModel
	UserID int `json:"user_id"`
	NameFolder  string `json:"name_folder"`
	FolderID int `json:"folder_id" gorm:"default:null;"`
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Folder *Folder `json:"folder" gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AccessId int `json:"access_id"`
	Access *Accesses `json:"access" gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	DefaultModel
	ID uint `gorm:"primary_key" json:"id"`
	Email string `json:"email"`
	Name string `json:"Name"`
	Password  string `json:"-"`
}

type File struct {
	DefaultModel
	UserID int `json:"user_id"`
	FolderID int `json:"folder_id" gorm:"default:null;"`
	Size int `json:"size"`
	FileName string `json:"file_name"`
	FileNameHash string `json:"file_name_hash"`
	User *User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Folder *Folder `json:"folder" gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccessId int `json:"access_id"`
	Access *Accesses `json:"access" gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Permissions struct {
	ID int `json:"id"`
	Name string `json:"name"`
	GuardName string `json:"guard_name"`
	PermissionGroupId int `json:"permission_group_id"`
	Type string `json:"type"`
	CreatedAt []uint8 `json:"created_at"`
	UpdatedAt []uint8 `json:"updated_at"`
}

type PermissionGroups struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	CreatedAt []uint8 `json:"created_at"`
	UpdatedAt []uint8 `json:"updated_at"`
}

type RoleHasPermissions struct {
	PermissionId int `json:"permission_id"`
	RoleId int `json:"role_id"`
}

type ModelHasRoles struct {
	RoleId int `json:"role_id"`
	ModelType string `json:"model_type"`
	ModelId int `json:"model_id"`
}

type Accesses struct {
	DefaultModel
	Name string `json:"name"`
}

type RequestAccess struct {
	DefaultModel
	UserID int `json:"user_id"`
	CurrentUserID int `json:"current_user_id"`
	StatusID int `json:"status_id"`
	FolderID int `json:"folder_id" gorm:"default:null;"`
	FileID int `json:"file_id" gorm:"default:null;"`
	Folder *Folder `json:"folder" gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	File *File `json:"file" gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User *User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CurrentUser *User `json:"current_user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status *Status `json:"status" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Status struct {
	DefaultModel
	Name string `json:"name"`
}