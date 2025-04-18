package server

import (
	"cloud-pricing-app/src/db"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

// fetchPricingData fetches data from the specified table with pagination

// parsePageQueryParam parses the "page" query parameter from the request
func parsePageQueryParam(r *http.Request) int {
	page := 1 // Default to page 1 if not specified
	queryPage := r.URL.Query().Get("page")
	if queryPage != "" {
		if parsedPage, err := strconv.Atoi(queryPage); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}
	return page
}

// awsHandler handles requests for AWS pricing data
func AwsHandler(w http.ResponseWriter, r *http.Request) {
	page := parsePageQueryParam(r)
	data, err := db.FetchPricingData("AWS", page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch AWS pricing data: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// azureHandler handles requests for Azure pricing data
func AzureHandler(w http.ResponseWriter, r *http.Request) {
	page := parsePageQueryParam(r)
	data, err := db.FetchPricingData("Azure", page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch Azure pricing data: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// gcpHandler handles requests for GCP pricing data
func GcpHandler(w http.ResponseWriter, r *http.Request) {
	page := parsePageQueryParam(r)
	data, err := db.FetchPricingData("GCP", page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch GCP pricing data: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// corsMiddleware adds CORS headers to the response
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
