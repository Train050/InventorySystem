package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("secret_key_yall")

type Register struct {
	FirstName string `json: firstName`
	LastName  string `json: lastName`
	Email     string `json: email`
	Password  string `json: password`
}
type Login struct {
	Email    string `json: email`
	Password string `json: password`
}
type User struct {
	ID        int64
	FirstName string `gorm:"column:firstName"`
	LastName  string `gorm:"column:lastName"`
	Email     string
	Password  string
}
type Tabler interface {
	TableName() string
}

//sqlite3 start of code
//opens file if it exists

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "user"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		var authHeader = c.Request.Header.Get("Authorization")
		substrings := strings.Split(authHeader, " ")
		tokenFromHeader := substrings[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenFromHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")

		} else {
			c.Next()
		}

	}
}
func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/rasyuenet?charset=utf8mb4&parseTime=True&loc=Local"
	r := gin.Default()

	config := cors.DefaultConfig()

	config.AllowHeaders = []string{"Authorization", "content-type"}
	config.AllowOrigins = []string{"http://localhost:4200"}
	r.Use(cors.New(config))
	r.POST("/register", func(c *gin.Context) {
		var registerData Register

		// Bind JSON Data to Object
		err := c.BindJSON(&registerData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
		}
		// hash the password
		hashPass, err := HashPassword(registerData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
		}
		var user User
		// db connection
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		user = User{FirstName: registerData.FirstName, LastName: registerData.LastName,
			Email: registerData.Email, Password: hashPass}
		db.Create(&user) // pass pointer of data to Create

		// create jwt to login
		expirationTime := time.Now().Add(30000 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Email: registerData.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"id":      user.ID,
			"jwt":     tokenString,
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var loginData Login
		var user User
		// Bind JSON Data to Object
		err := c.BindJSON(&loginData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
		}
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		db.First(&user, "email = ?", loginData.Email)
		checkHash := CheckPasswordHash(loginData.Password, user.Password)
		if checkHash == true {

			expirationTime := time.Now().Add(5 * time.Minute)
			// Create the JWT claims, which includes the username and expiry time
			claims := &Claims{
				Email: loginData.Email,
				RegisteredClaims: jwt.RegisteredClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: jwt.NewNumericDate(expirationTime),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Create the JWT string
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "")
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"jwt":     tokenString,
			})
		} else {
			c.JSON(http.StatusInternalServerError, "")
		}
	})

	r.GET("/user-session", auth(), func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
