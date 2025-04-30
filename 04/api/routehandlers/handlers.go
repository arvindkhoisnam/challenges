package routehandlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/arvindkhoisnam/challenges/04/middleware"
	"github.com/arvindkhoisnam/challenges/04/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository struct{
	DbClient *gorm.DB
}
type SignupBody struct {
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	FirstName string  `json:"firstname" binding:"required"`
	LastName  string  `json:"lastname" binding:"required"`
}
type SigninBody struct {
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}


func (r *Repository)Routes(app *gin.Engine){
	api := app.Group("/api/v1")
	api.GET("/health",r.HealthEndpoint)
	api.POST("/signup",r.SignupEndpoint)
	api.POST("/signin",r.SigninEndpoint)
	api.GET("/users",middleware.AuthMiddleware(),r.AllUsers)
	api.GET("/user",middleware.AuthMiddleware(),r.GetUser)
}

func (r *Repository)HealthEndpoint(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{"message":"healthy server"})
}
func (r *Repository)SignupEndpoint(ctx *gin.Context){
	var newUser SignupBody
	if err := ctx.ShouldBindJSON(&newUser); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(newUser.Password),bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"Could not hash password"})
		return
	}
	user := &models.UserModel{
		Username: newUser.Username,
		Password: string(hashedPassword),
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		CreateAt: time.Now().Format(time.RFC3339),
		Role: "user",
	}

	err = r.DbClient.Create(user).Error
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,gin.H{"message":fmt.Sprintf("New user created with id %s",user.Id)})
}

func (r *Repository)SigninEndpoint(ctx *gin.Context){
	var reqBody SigninBody
	userModel := &models.UserModel{}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}

	if err := r.DbClient.Where("username = ?",reqBody.Username).First(userModel).Error; err != nil{
		ctx.JSON(http.StatusForbidden,gin.H{"message":"Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password),[]byte(reqBody.Password)); err != nil{
		ctx.JSON(http.StatusForbidden,gin.H{"message":"Invalid username or password"})
		return
	}
	jwtSec := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":userModel.Id,
		"role":userModel.Role,
	})
	tokenString,err := token.SignedString([]byte(jwtSec))
	if err != nil {
		ctx.JSON(http.StatusForbidden,gin.H{"message":"Could not sign jwt"})
		return
    }
	ctx.SetCookie(
		"auth_token",      // name
		tokenString,       // value
		3600,              // maxAge in seconds
		"/",               // path
		"",                // domain ("" = current domain)
		true,              // secure
		true,              // httpOnly
	)
	ctx.JSON(http.StatusOK,gin.H{"message":"successfully signed in."})
}

func (r *Repository)AllUsers(ctx *gin.Context){
	var usernames []string
	var users []models.UserModel
	r.DbClient.Find(&users)
	for _,user := range users{
		usernames = append(usernames, user.Username)
	}
	time.Sleep(5*time.Second)
	ctx.JSON(http.StatusOK,gin.H{"data":usernames})
}

func (r *Repository)GetUser(ctx *gin.Context){
	user := &models.UserModel{}
	userId,found := ctx.Get("user_id")
	if !found{
		ctx.JSON(http.StatusForbidden,gin.H{"message":"No user id found"})
		return
	}

	if err := r.DbClient.Where("id = ?", userId).Find(user).Error;err != nil{
		ctx.JSON(http.StatusForbidden,gin.H{"message":"No user found"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":fmt.Sprintf("Welcome back %s",user.Username)})
}