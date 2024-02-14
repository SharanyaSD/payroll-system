package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/api"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/app"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
)

func main() {

	//	fmt.Print(os.Getenv("PATH"))

	sqlDB, err := repository.InitializeDB()
	if err != nil {
		fmt.Println("error initializing database:", err)
		return
	}

	apiKey := os.Getenv("API_KEY")
	// Initialize service dependencies
	services := app.NewServices(sqlDB, apiKey)
	router := api.NewRouter(services)

	fmt.Println("HTTP Server started on port 8080...")
	err = http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("error starting server:", err)
	}

	// Goroutine for OTP server (port 8000)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("OTP Server started on port 8000...")
	// 	router := gin.Default()
	// 	app := api.Config{Router: router}
	// 	app.Routes()
	// 	if err := router.Run(":8000"); err != nil {
	// 		fmt.Println("error starting OTP server:", err)
	// 	}
	// }()

	// Wait for all goroutines to finish

}
