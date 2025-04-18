package gcp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"cloud-pricing-app/src/db"
)

type ServicePricing struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type PricingInfo struct {
	Compute []ServicePricing `json:"compute"`
	Storage []ServicePricing `json:"storage"`
}

func FetchAndUpdateGCPPricesList() (*PricingInfo, error) {
	apiKey := os.Getenv("APIKEYGOOGLE")
	pageSize := 1000
	var nextPageToken string
	var pricingInfo PricingInfo

	for {
		url := fmt.Sprintf("https://cloudbilling.googleapis.com/v1/services/6F81-5844-456A/skus?key=%s&pageSize=%v&pageToken=%s", apiKey, pageSize, nextPageToken)
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get products: %v", resp.Status)
		}

		var rawData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&rawData); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}

		// Parse the raw data to extract compute and storage pricing
		for _, product := range rawData["skus"].([]interface{}) {
			productMap := product.(map[string]interface{})
			sku := productMap["skuId"].(string)
			priceItemBytes, err := json.Marshal(product)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal price item: %v", err)
			}
			priceItemStr := string(priceItemBytes)

			err = db.InsertGCPPricingData(sku, priceItemStr)
			if err != nil {
				return nil, fmt.Errorf("failed to insert price item into database: %v", err)
			}
		}

		nextPageToken, _ = rawData["nextPageToken"].(string)
		if nextPageToken == "" {
			break
		}
	}

	//pageSize := "1000"
	//pagetoken := "
	url := fmt.Sprintf("https://cloudbilling.googleapis.com/v1/services/6F81-5844-456A/skus?key=%s&pageSize=%v", apiKey, pageSize)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rawData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	//var pricingInfo PricingInfo

	// Parse the raw data to extract compute and storage pricing
	for _, product := range rawData["skus"].([]interface{}) {
		productMap := product.(map[string]interface{})
		sku := productMap["skuId"].(string)
		priceItemBytes, err := json.Marshal(product)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal price item: %v", err)
		}
		priceItemStr := string(priceItemBytes)

		err = db.InsertGCPPricingData(sku, priceItemStr)
		if err != nil {
			return nil, fmt.Errorf("failed to insert price item into database: %v", err)
		}
	}

	return &pricingInfo, nil
}
