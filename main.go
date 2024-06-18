package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-playground/validator/v10"
)

// Define your JSON structure
type Application struct {
	Name               string `json:"name" validate:"required"`
	Path               string `json:"path validate:"required"`
	PrimaryEndpoints   []string `json:"primary_endpoints validate:"required"`
	FailoverEndpoints  []string `json:"failover_endpoints"`
	PathRewrite        string `json:"path_rewrite"`
	OutlierDetection   struct {
		Consecutive5xx   *int `json:"consecutive_5xx"`
		BaseEjectionTime *int `json:"base_ejection_time"`
		MaxEjectionTime  *int `json:"max_ejection_time"`
	} `json:"outlier_detection"`
	SessionAffinity    struct {
		GenerateStickySession bool `json:"generate_sticky_session"`
		CookieName            string `json:"cookie_name"`
		CookiePath            string `json:"cookie_path"`
	} `json:"session_affinity"`
	ConnectionTimeout  *int `json:"connection_timeout"`
	MaxRequestsPerConnection *int `json:"max_requests_per_connection"`
	Keepalive          struct {
		Enabled         bool `json:"enabled"`
		TCPKeepalive    struct {
			KeepaliveProbes int `json:"keepalive_probes"`
			KeepaliveInterval int `json:"keepalive_interval"`
			KeepaliveTime int `json:"keepalive_time"`
		} `json:"tcp_keepalive"`
	} `json:"keepalive"`
	MaxConnections     *int `json:"max_connections"`
	ResponseTimeout    int `json:"response_timeout"`
	Healthcheck        struct {
		Timeout         int `json:"timeout"`
		Interval        int `json:"interval"`
		UnhealthyThreshold int `json:"unhealthy Threshold"`
		HealthyThreshold int `json:"healthy Threshold"`
		HTTP            struct {
			Host          string `json:"host"`
			Path          string `json:"path"`
			ExpectedResponse struct {
				Start     int `json:"start"`
				End       int `json:"end"`
			} `json:"expectedResponse"`
		} `json:"http"`
	} `json:"healthcheck"`
}

func main() {
	filename := "/Users/peeps/Applications/hello/configs.json" // Replace with the actual file path

	// Read the JSON data from the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Unmarshal the JSON data into the struct
	var app Application
	if err := json.Unmarshal(data, &app); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	// Validate the struct using the validator
	validate := validator.New()
	if err := validate.Struct(app); err != nil {
		// Print the validation errors
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation: %s\n", err.Field(), err.Tag())
		}
		return
	}

	fmt.Println("JSON is valid!")
}
