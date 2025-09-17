package main

import (
	"log"
	"net/http"
	"script_pilot/config"
	"script_pilot/config/logger"
	"script_pilot/internal/routes"
	"script_pilot/internal/service"
)

func main() {
	logger.Logger.Info("Server listening on http://localhost:8080")
	service.NewGeminiClient(config.Config.GeminiKey, "gemini-2.5-flash")
	log.Fatal(http.ListenAndServe(":8080", routes.NewRouter()))
}
