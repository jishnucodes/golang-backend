package handlers

import (
	// "clinic-management/backend/builder"
	// "clinic-management/backend/builder"
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/managers"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	userGroup.GET("/:userId", handler.GetAUser)
	userGroup.POST("/create", handler.InsertUser);
	// userGroup.POST("/login", handler.Login)
	// userGroup.PUT("/update/:userId", handler.UpdateUser)
}

// func (handler *UserHandler) Login(ctx *gin.Context) {
	
// 	userData := common.NewUserOb()

// 	err := ctx.BindJSON(&userData)

// 	if err!= nil {
// 		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		log.Println("binding user details from json is failed:", err)
// 		common.SendError(nil, http.StatusBadRequest, err )
// 		return
// 	}

// 	result, err := handler.userManager.Login(userData)

// 	fmt.Println("error", err)

// 	if err!= nil {
// 		log.Println("user login failed", err)
// 		common.SendError(ctx, http.StatusInternalServerError, err )
// 		return
// 	}

// 	common.SendSuccess(ctx, http.StatusOK, result)
// 	log.Println("user loggined successfully")
// }

func (handler *UserHandler) UserList(ctx *gin.Context) {

	users, err := handler.userManager.GetUsers()

    response := builder.BuildUserDTOs(users.Result)


    fmt.Println("users", fmt.Sprintf("%T", users.Result))
    

    fmt.Println("data type:", fmt.Sprintf("%T", users)) 

	if err!= nil {
		log.Println("user fetching failed", err)
		common.SendError(ctx, http.StatusInternalServerError, users.Message, err )
		return
	}

	common.SendSuccess(ctx, http.StatusOK, users.Message, response)
	log.Println("user fetched successfully")


}

func (handler *UserHandler) GetAUser(ctx *gin.Context) {
    // Create a new user object
    userData := common.NewUserOb()

    // Bind the request body to the userData object
    // err := ctx.BindJSON(&userData)
    // if err != nil {
    //     log.Println("binding user details from JSON failed:", err)
    //     common.SendError(ctx, http.StatusBadRequest, err)
    //     return
    // }

    // Extract the user ID from the URL parameters
    userIdStr, ok := ctx.Params.Get("userId")
    if !ok {
        log.Println("user ID is missing in the request")
        common.SendError(ctx, http.StatusBadRequest,"", fmt.Errorf("user ID is required"))
        return
    }

	// Convert string userId to uint
    userId, err := strconv.ParseUint(userIdStr, 10, 32)
    if err != nil {
        log.Println("invalid user ID format:", err)
        common.SendError(ctx, http.StatusBadRequest, "",fmt.Errorf("invalid user ID format"))
        return
    }

    // Assign the user ID to the userData object
    userData.UserID = uint(userId)

    // Call the update method in the user manager
    user, err := handler.userManager.GetAUser(userData)

    response := builder.BuildUserDTOs(user.Result)

    if err != nil {
        log.Println("user update failed:", err)
        common.SendError(ctx, http.StatusInternalServerError, user.Message, err)
        return
    }

    // Send success response
    common.SendSuccess(ctx, http.StatusOK, user.Message, response)
    log.Println("user updated successfully")
}

func (handler *UserHandler) InsertUser(ctx *gin.Context) {
	userData := common.NewUserOb()

	err := ctx.BindJSON(&userData)

	if err!= nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("binding user details from json is failed:", err)
		common.SendError(nil, http.StatusBadRequest,"", err )
		return
	}

	newUser, err := handler.userManager.CreateUser(userData)

    // response := builder.BuildUserDTOs(newUser.Result)

	fmt.Println("error", err)

	if err!= nil {
		log.Println("user insertion failed", err)
		common.SendError(ctx, http.StatusInternalServerError, "",err )
		return
	}

	common.SendSuccess(ctx, http.StatusOK,newUser.Message, newUser.Result)
	log.Println("user created successfully")
}

// func (handler *UserHandler) UpdateUser(ctx *gin.Context) {
//     // Create a new user object
//     userData := common.NewUserOb()

//     // Bind the request body to the userData object
//     err := ctx.BindJSON(&userData)
//     if err != nil {
//         log.Println("binding user details from JSON failed:", err)
//         common.SendError(ctx, http.StatusBadRequest, err)
//         return
//     }

//     // Extract the user ID from the URL parameters
//     userIdStr, ok := ctx.Params.Get("userId")
//     if !ok {
//         log.Println("user ID is missing in the request")
//         common.SendError(ctx, http.StatusBadRequest, fmt.Errorf("user ID is required"))
//         return
//     }

// 	// Convert string userId to uint
//     userId, err := strconv.ParseUint(userIdStr, 10, 32)
//     if err != nil {
//         log.Println("invalid user ID format:", err)
//         common.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid user ID format"))
//         return
//     }

//     // Assign the user ID to the userData object
//     userData.UserID = uint(userId)

//     // Call the update method in the user manager
//     updatedUser, err := handler.userManager.UpdateUser(userData)

//     if err != nil {
//         log.Println("user update failed:", err)
//         common.SendError(ctx, http.StatusInternalServerError, err)
//         return
//     }

//     // Send success response
//     common.SendSuccess(ctx, http.StatusOK, updatedUser)
//     log.Println("user updated successfully")
// }

