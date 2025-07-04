package main

import (
	"clinic-management/backend/config"
	"clinic-management/backend/database"
	"clinic-management/backend/handlers"
	"clinic-management/backend/managers"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func fmtlog(params ...any) {
	fmt.Println(params...)
	log.Println(params...)
}

func init() {
	database.Initialize()
}

func main() {
	fmt.Println("Hello guys")

	//used to load the config
	config.ReadConfig()
	cfg := config.GetConfig()

	fmt.Println("Database Driver:", cfg.DbServer.Driver)
	fmt.Println("Database Connection String:", cfg.DbServer.ConnectionString)
	fmt.Println("SAP Username:", cfg.SapServer.Username)
	fmt.Println("SAP Service Layer URL:", cfg.SapServer.Servicelayerurl)
	fmt.Println("Port:", cfg.Port)
	fmt.Println("WebSocket Port:", cfg.WebSocketPort)

	//logging section
	execpath, _ := os.Getwd()
	logfile := "clientManagement_" + time.Now().Format("20060102") + ".log"
	logfile = filepath.Join(execpath, logfile)
	fmt.Println("Logging to", logfile)
	logFile, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	fmt.Println(logFile.Stat())
	log.SetOutput(logFile)

	// port := strconv.Itoa(cfg.Port)
	// envport := os.Getenv("ASPNETCORE_PORT")
	// if envport != "" { // get enviroment variable that set by ACNM
	// 	_, err := strconv.Atoi(envport)
	// 	if err != nil {
	// 		fmtlog("Unable to read port form environment variable ASPNETCORE_PORT", err, envport)
	// 	} else {
	// 		port = envport
	// 		fmtlog("Using port form environment variable ASPNETCORE_PORT", envport)
	// 	}
	// }
	// fmtlog("Running on port", port)

	errors := godotenv.Load()
	if errors != nil {
		log.Println("No .env file found or error loading it")
	}

	port := strconv.Itoa(cfg.Port) // default port from config

	// Override if PORT environment variable is set (used by Render and most platforms)
	if envport := os.Getenv("PORT"); envport != "" {
		if _, err := strconv.Atoi(envport); err == nil {
			port = envport
			fmtlog("Using port from environment variable PORT", envport)
		} else {
			fmtlog("Invalid port from environment variable PORT", err, envport)
		}
	} else if envport := os.Getenv("ASPNETCORE_PORT"); envport != "" {
		// Fallback for IIS or Azure
		if _, err := strconv.Atoi(envport); err == nil {
			port = envport
			fmtlog("Using port from ASPNETCORE_PORT", envport)
		} else {
			fmtlog("Invalid ASPNETCORE_PORT value", err, envport)
		}
	}

	fmtlog("Running on port", port)

	router := gin.Default()

	router.Use(cors.Default())

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandler(userManager)
	userHandler.RegisterApis(router)

	patientManager := managers.NewPatientManager()
	patientHandler := handlers.NewPatientHandler(patientManager)
	patientHandler.RegisterApis(router)

	doctorManager := managers.NewDoctorManager()
	doctorHandler := handlers.NewDoctorHandler(doctorManager)
	doctorHandler.RegisterApis(router)

	doctorAvailabilityManager := managers.NewDoctorAvailabilityManager()
	doctorAvailabilityHandler := handlers.NewDoctorAvailabilityHandler(doctorAvailabilityManager)
	doctorAvailabilityHandler.RegisterApis(router)

	roleManager := managers.NewRoleManager()
	roleHandler := handlers.NewRoleHandler(roleManager)
	roleHandler.RegisterApis(router)

	departmentManager := managers.NewDepartmentManager()
	departmentHandler := handlers.NewDepartmentHandler(departmentManager)
	departmentHandler.RegisterApis(router)

	employeeManager := managers.NewEmployeeManager()
	employeeHandler := handlers.NewEmployeeHandler(employeeManager)
	employeeHandler.RegisterApis(router)

	autoNumberManager := managers.NewAutoNumberManager()
	autoNumberHandler := handlers.NewAutoNumberHandler(autoNumberManager)
	autoNumberHandler.RegisterApis(router)

	employeeLeaveManager := managers.NewEmployeeLeaveManager()
	employeeLeaveHandler := handlers.NewEmployeeLeaveHandler(employeeLeaveManager)
	employeeLeaveHandler.RegisterApis(router)

	appointmentManager := managers.NewAppointmentManager()
	appointmentHandler := handlers.NewAppointmentHandler(appointmentManager)
	appointmentHandler.RegisterApis(router)

	paymentManager := managers.NewPaymentManager()
	paymentHandler := handlers.NewPaymentHandler(paymentManager)
	paymentHandler.RegisterApis(router)

	consultationMedicationManager := managers.NewConsultationMedicationManager()
	consultationMedicationHandler := handlers.NewConsultationMedicationHandler(consultationMedicationManager)
	consultationMedicationHandler.RegisterApis(router)

	medicineMasterManager := managers.NewMedicineMasterManager()
	medicineMasterHandler := handlers.NewMedicineMasterHandler(medicineMasterManager)
	medicineMasterHandler.RegisterApis(router)
	//api for adding chat
	whatsAppChatHandler := handlers.NewWhatsAppHandler()
	whatsAppChatHandler.RegisterApis(router)

	s := &http.Server{
		Addr:           ":" + port, //":5005",
		Handler:        router,
		ReadTimeout:    600 * time.Second, // Increase read timeout
		WriteTimeout:   600 * time.Second, // Increase write timeout
		IdleTimeout:    600 * time.Second, // Increase idle timeout
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
