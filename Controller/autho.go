package Controller

import (
	"main/intializers"
	"main/models"
	"main/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var users struct {
		Email    string
		Password string
	}
	if err := c.ShouldBind(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(users.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{
		Email:    users.Email,
		Password: string(hashed),
	}
	if err := intializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})

}
func Login(c *gin.Context) {
	var users struct {
		Email    string
		Password string
	}
	if err := c.ShouldBind(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	var user models.User

	intializers.DB.First(&user, "Email = ?", users.Email)
	if user.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Invalid user/login credentials",
		})
		return
	}
	accessToken, err := utils.CreateJWT(user.ID, 15*time.Minute)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not create access token",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(users.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	refreshToken, err := utils.CreateRefreshJWT(user.ID, 7*24*time.Hour)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not create refresh token",
		})
		return
	}
	c.SetCookie("refresh_token", refreshToken, int((7 * 24 * time.Hour).Seconds()), "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"access token": accessToken,
	})

}
func Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "unauthorised refrsh token has been given by the cookies",
		})
		return

	}
	intializers.DB.Where("token = ?", refreshToken).Delete(&models.RefreshToken{})

	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"message": "logout has been successfully",
	})

}
func Validate(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "i am logged in",
	})

}
