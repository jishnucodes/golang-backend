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

type PatientHandler struct {
	groupName string
	patientManager managers.PatientManager
}

func NewPatientHandler (patientManager managers.PatientManager) *PatientHandler {
	return &PatientHandler {
		"api/patient",
		patientManager,
	}
}

func (handler *PatientHandler) RegisterApis(r *gin.Engine) {
	patientGroup := r.Group(handler.groupName);
	// patientGroup.GET("/list", handler.UserList);
	patientGroup.POST("/create", handler.InsertPatient);
}


// func (handler *UserHandler) UserList(ctx *gin.Context) {

// 	users, err := handler.userManager.GetUsers()

// 	if err!= nil {
// 		log.Println("user fetching failed", err)
// 		common.SendError(ctx, http.StatusInternalServerError, err )
// 		return
// 	}

// 	common.SendSuccess(ctx, http.StatusOK, users)
// 	log.Println("user fetched successfully")


// }

func (handler *PatientHandler) InsertPatient(ctx *gin.Context) {
	patientData := common.NewPatientCreationInput()

	err := ctx.BindJSON(&patientData)

	if err!= nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("binding patient details from json is failed:", err)
		common.SendError(nil, http.StatusBadRequest, err )
		return
	}

	newPatient, err := handler.patientManager.CreatePatient(patientData)

	fmt.Println("error", err)

	if err!= nil {
		log.Println("patient details insertion failed", err)
		common.SendError(ctx, http.StatusInternalServerError, err )
		return
	}

	common.SendSuccess(ctx, http.StatusOK, newPatient)
	log.Println("patient created successfully")
}

