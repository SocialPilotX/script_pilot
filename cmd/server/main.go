package main

import (
	"log"
	"net/http"
	"script_pilot/config/logger"
)

func main() {
	logger.Logger.Info("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
