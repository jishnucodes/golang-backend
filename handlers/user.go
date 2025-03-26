package handlers

import (
	"clinic-management/backend/common"
	"clinic-management/backend/managers"
	"fmt"
	"log"
	"net/http"

	// "net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName string
	userManager managers.UserManager
}

func NewUserHandler (userManager managers.UserManager) *UserHandler {
	return &UserHandler {
		"api/user",
		userManager,
	}
}

func (handler *UserHandler) RegisterApis(r *gin.Engine) {
	userGroup := r.Group(handler.groupName);
	userGroup.GET("/list", handler.UserList);
	userGroup.POST("/create", handler.InsertUser);
	userGroup.POST("/login", handler.Login)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	
	userData := common.NewUserOb()

	err := ctx.BindJSON(&userData)

	if err!= nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("binding user details from json is failed:", err)
		common.SendError(nil, http.StatusBadRequest, err )
		return
	}

	result, err := handler.userManager.Login(userData)

	fmt.Println("error", err)

	if err!= nil {
		log.Println("user login failed", err)
		common.SendError(ctx, http.StatusInternalServerError, err )
		return
	}

	common.SendSuccess(ctx, http.StatusOK, result)
	log.Println("user loggined successfully")
}

func (handler *UserHandler) UserList(ctx *gin.Context) {

	users, err := handler.userManager.GetUsers()

	if err!= nil {
		log.Println("user fetching failed", err)
		common.SendError(ctx, http.StatusInternalServerError, err )
		return
	}

	common.SendSuccess(ctx, http.StatusOK, users)
	log.Println("user fetched successfully")


}

func (handler *UserHandler) InsertUser(ctx *gin.Context) {
	userData := common.NewUserOb()

	err := ctx.BindJSON(&userData)

	if err!= nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("binding user details from json is failed:", err)
		common.SendError(nil, http.StatusBadRequest, err )
		return
	}

	newUser, err := handler.userManager.CreateUser(userData)

	fmt.Println("error", err)

	if err!= nil {
		log.Println("user insertion failed", err)
		common.SendError(ctx, http.StatusInternalServerError, err )
		return
	}

	common.SendSuccess(ctx, http.StatusOK, newUser)
	log.Println("user created successfully")
}

