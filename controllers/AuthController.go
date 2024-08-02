package Controller

import (
	"mypackages/db"
	Model "mypackages/models"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"` 
}

type RegistrationRequest struct {
	*LoginRequest
	Name    string `form:"name" json:"name" xml:"name"  binding:"required"`
}

var secretKey, _ = os.LookupEnv("SECRET_KEY")
var jwtSecretKey = []byte(secretKey)

func Login(c *gin.Context){
	var jsonValid LoginRequest

	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}

	var user *Model.User;

	r := db.DB.Unscoped().Model(&user).First(&user, "email = ?", jsonValid.Email)
	match := CheckPasswordHash(jsonValid.Password, user.Password)
	
	if r.RowsAffected == 0 || !match{
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}

    payload := jwt.MapClaims{
		"email": user.Email,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, _ := token.SignedString(jwtSecretKey)
	c.JSON(200, gin.H{
		"message": "success",
		"access_token": t,
	})
}

func Registration(c *gin.Context){
	var jsonValid RegistrationRequest

	if err := c.ShouldBind(&jsonValid); err != nil{
		c.JSON(422, gin.H{
			"message": "error",
		})
		return
	}

	var user *Model.User;

	r := db.DB.Unscoped().Model(&user).First(&user, "email = ?", jsonValid.Email)
	if r.RowsAffected > 0 {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}

	pass, _ := HashPassword(jsonValid.Password)

	newUser := Model.User{
		Email: jsonValid.Email,
		Password: pass,	
		Name: jsonValid.Name,	
	}

	r = db.DB.Create(&newUser)
	os.Mkdir("files/"+strconv.Itoa(int(user.ID)), os.ModePerm)

	if r.RowsAffected>0{
		c.JSON(201, gin.H{
			"message": "success",
		})
		return
	}

	c.JSON(500, gin.H{
		"message": "error",
	})
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}