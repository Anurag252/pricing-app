package aws

import (
	"encoding/json"
	"fmt"

	"cloud-pricing-app/src/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pricing"
)

type ServicePricing struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type PricingInfo struct {
	Compute []ServicePricing `json:"compute"`
}

func FetchAndUpdateAWSPricesList() (*PricingInfo, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	svc := pricing.New(sess)

	input := &pricing.GetProductsInput{
		ServiceCode: aws.String("AmazonEC2"),
		Filters: []*pricing.Filter{
			{
				Type:  aws.String("TERM_MATCH"),
				Field: aws.String("productFamily"),
				Value: aws.String("Compute Instance"),
			},
		},
	}

	var pricingInfo PricingInfo
	var nextToken *string

	for {
		result, err := svc.GetProducts(input)
		if err != nil {
			return nil, fmt.Errorf("failed to get products: %v", err)
		}

		for _, priceItem := range result.PriceList {
			sku := priceItem["product"].(map[string]interface{})["sku"].(string)
			priceItemBytes, err := json.Marshal(priceItem)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal price item: %v", err)
			}
			priceItemStr := string(priceItemBytes)

			err = db.InsertAWSPricingData(sku, priceItemStr)
			if err != nil {
				return nil, fmt.Errorf("failed to insert price item into database: %v", err)
			}
		}

		nextToken = result.NextToken
		if nextToken == nil {
			break
		}
		input.NextToken = nextToken
	}

	return &pricingInfo, nil
}
