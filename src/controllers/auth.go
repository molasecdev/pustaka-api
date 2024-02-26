package controllers

import (
	"net/http"
	"os"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// register
func Register(c *gin.Context) {
	db := config.InitConfig()
	var body types.InputRegister
	var role models.Role
	var findEmail models.Auth
	var findNik models.User

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	db.Where("email = ?", body.Email).First(&findEmail)
	db.Where("nik = ?", body.Nik).First(&findNik)

	if body.Email == findEmail.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists!"})
		return
	}
	if body.Nik == findNik.Nik {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NIK already exists!"})
		return
	}

	result := db.Where("LOWER(role) = LOWER(?)", body.Role).First(&role)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found!"})
		return
	}

	password := body.Firstname + "123"
	hashPass, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)
	if errHash != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errHash.Error()})
		return
	}

	auth := models.Auth{
		Email:    body.Email,
		Password: string(hashPass),
	}
	createAuth := db.Create(&auth)
	if createAuth.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": createAuth.Error})
		return
	}

	user := models.User{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Birthday:  body.Birthday,
		Address:   body.Address,
		Nik:       body.Nik,
		Phone:     body.Phone,
		Role_id:   role.ID,
		Auth_id:   auth.ID,
	}
	createUser := db.Create(&user)
	if createUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": createUser.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Register successfully"})
}

// login
func Login(c *gin.Context) {
	db := config.InitConfig()
	var SECRET_KEY = os.Getenv("SECRET_KEY")
	var user models.User
	var auth models.Auth
	var body types.InputLogin
	var token types.JwtToken

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	db.Where("email = ?", body.Email).First(&auth)
	db.Where("auth_id = ?", auth.ID).Preload("Role").Preload("Auth").First(&user)

	errMatch := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(body.Password))

	if auth.Email != body.Email || errMatch != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email/password!"})
		return
	}

	lifetime := 1 // 1 day
	now := time.Now()

	token.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "Putaka-API",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * time.Duration(lifetime))),
	}

	// JWT contains
	token.Sub = user.ID
	token.Role = user.Role.Role

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	accessToken, err := _token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// set JWT
	c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", accessToken, 3600*24*lifetime, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "access_token": accessToken})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Valid"})
}
