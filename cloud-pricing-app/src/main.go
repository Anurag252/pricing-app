package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud-pricing-app/src/aws"
	"cloud-pricing-app/src/azure"
	"cloud-pricing-app/src/gcp"
	"cloud-pricing-app/src/server"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/aws", server.AwsHandler)
	mux.HandleFunc("/api/azure", server.AzureHandler)
	mux.HandleFunc("/api/gcp", server.GcpHandler)

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", server.CorsMiddleware(mux)))
	go func() {
		awsPricing, err := aws.FetchAndUpdateAWSPricesList()
		if err != nil {
			log.Fatalf("Error fetching AWS pricing: %v", err)
		}
		fmt.Println("AWS Pricing:", awsPricing)

	}()

	go func() {
		azurePricing, err := azure.FetchAndUpdateAzurePricesList()
		if err != nil {
			log.Fatalf("Error fetching Azure pricing: %v", err)
		}
		fmt.Println("Azure Pricing:", azurePricing)

	}()

	gcp.FetchAndUpdateGCPPricesList()

	go func() {
		gcpPricing, err := gcp.FetchAndUpdateGCPPricesList()
		if err != nil {
			log.Fatalf("Error fetching GCP pricing: %v", err)
		}
		fmt.Println("GCP Pricing:", gcpPricing)
	}()
}
