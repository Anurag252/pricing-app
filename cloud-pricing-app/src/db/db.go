package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbURL    = "postgresql://postgres.%s:fO24aoA1V5b9wkMp@aws-0-us-west-1.pooler.supabase.com:6543/postgres"
	pageSize = 20 // Number of records per page
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf(dbURL, os.Getenv("DBPASS")))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

func InsertAWSPricingData(sku string, priceItem string) error {
	//timestamp := time.Now()

	_, err := db.Exec(`INSERT INTO "AWS" (sku, priceitem) VALUES ($1, $2) ON CONFLICT (sku) DO UPDATE SET priceitem = EXCLUDED.priceitem`,
		sku, priceItem)
	if err != nil {
		return err
	}

	return nil
}

// PricingData represents the structure of a pricing row
type PricingData struct {
	SKU       string `json:"sku"`
	PriceItem string `json:"priceitem"`
	Timestamp string `json:"created_at"`
}

func FetchPricingData(tableName string, page int) ([]PricingData, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf(`SELECT sku, priceitem, created_at FROM "%s" LIMIT %d OFFSET %d`, tableName, pageSize, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pricingData []PricingData
	for rows.Next() {
		var data PricingData
		if err := rows.Scan(&data.SKU, &data.PriceItem, &data.Timestamp); err != nil {
			return nil, err
		}
		pricingData = append(pricingData, data)
	}

	return pricingData, nil
}

func InsertAzurePricingData(sku string, priceItem string) error {

	_, err := db.Exec(`INSERT INTO "Azure" (sku, priceitem) VALUES ($1, $2) ON CONFLICT (sku) DO UPDATE SET priceitem = EXCLUDED.priceitem`,
		sku, priceItem)
	if err != nil {
		return err
	}

	return nil
}

func InsertGCPPricingData(sku string, priceItem string) error {

	_, err := db.Exec(`INSERT INTO "GCP" (sku, priceitem) VALUES ($1, $2) ON CONFLICT (sku) DO UPDATE SET priceitem = EXCLUDED.priceitem`,
		sku, priceItem)
	if err != nil {
		return err
	}

	return nil
}
