package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/status", getStatus)
	http.HandleFunc("/healthz", getHealth)
	http.ListenAndServe(":8080", nil)
}

func getStatus(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"pod_name": os.Getenv("POD_NAME"),
		"pod_ip":   os.Getenv("POD_IP"),
		"version":  os.Getenv("VERSION"),
	}
	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(j)
}

func getHealth(w http.ResponseWriter, req *http.Request) {
	isHealthy, err := strconv.ParseBool(os.Getenv("IS_HEALTHY"))
	if err != nil {
		isHealthy = true
	}
	w.Header().Set("Content-Type", "application/json")
	status := map[string]string{}
	if !isHealthy {
		w.WriteHeader(http.StatusServiceUnavailable)
		status["status"] = "unhealthy"
	} else {
		w.WriteHeader(http.StatusOK)
		status["status"] = "healthy"
	}
	j, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	w.Write(j)
}
