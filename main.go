package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found. CONTAINER_API_PORT and HOST_API_PORT are both required.")
	}

	containerAPIPort, containerAPIPortExists := os.LookupEnv("CONTAINER_API_PORT")
	hostAPIPort, hostAPIPortExists := os.LookupEnv("HOST_API_PORT")

	if containerAPIPortExists && hostAPIPortExists {
		os.Setenv("containerAPIPort", containerAPIPort)
		os.Setenv("hostAPIPort", hostAPIPort)
		fmt.Println("Container expose port: " + os.Getenv("containerAPIPort"))
		fmt.Println("Localhost mount port (point postman here): " + os.Getenv("hostAPIPort"))
	} else {
		log.Fatal(".env vars not defined. CONTAINER_API_PORT and HOST_API_PORT are both required.")
	}
}

func main() {

	srv := &http.Server{
		Addr:         "0.0.0.0:" + os.Getenv("hostAPIPort"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "is ddddworking"}`))
	})

	log.Fatal(srv.ListenAndServe())
}
