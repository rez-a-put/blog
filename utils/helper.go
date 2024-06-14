package utils

import (
	m "blog/model"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	projectDirName := "blog"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}
}

func GetEnvByKey(key string) string {
	return os.Getenv(key)
}

// setup response data
func ReturnResponse(w http.ResponseWriter, statusCode int, respMsg string, retData interface{}) {
	respData := &m.Response{
		Status:  strconv.Itoa(statusCode),
		Message: respMsg,
		Data:    retData,
	}

	// convert data into json and send as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(respData)
	if err != nil {
		panic(err.Error())
	}
}
