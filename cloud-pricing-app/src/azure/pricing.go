package azure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud-pricing-app/src/db"
)

type ServicePricing struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type PricingInfo struct {
	Compute []ServicePricing `json:"compute"`
}

func FetchAndUpdateAzurePricesList() (*PricingInfo, error) {
	url := "https://prices.azure.com/api/retail/prices"
	var pricingInfo PricingInfo
	var nextPage string

	for {
		if nextPage != "" {
			url = nextPage
		}

		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to get products: %v", err)
		}
		defer resp.Body.Close()

		var result struct {
			Items    []map[string]interface{} `json:"Items"`
			NextPage string                   `json:"NextPage"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}

		for _, priceItem := range result.Items {
			sku := priceItem["skuId"].(string)
			priceItemBytes, err := json.Marshal(priceItem)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal price item: %v", err)
			}
			priceItemStr := string(priceItemBytes)

			err = db.InsertAzurePricingData(sku, priceItemStr)
			if err != nil {
				return nil, fmt.Errorf("failed to insert price item into database: %v", err)
			}
		}

		nextPage = result.NextPage
		if nextPage == "" {
			break
		}
	}

	return &pricingInfo, nil
}
