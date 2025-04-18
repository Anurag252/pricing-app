package types

type PricingInfo struct {
	Service   string  `json:"service"`
	Region    string  `json:"region"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
}

type ServicePricing struct {
	Compute PricingInfo `json:"compute"`
	Storage PricingInfo `json:"storage"`
}