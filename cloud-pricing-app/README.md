# Cloud Pricing Application

This project is a cloud pricing application that queries AWS, Azure, and GCP to fetch the prices of compute and storage services. It provides a unified interface to access pricing information from these major cloud providers.

## Project Structure

```
cloud-pricing-app
├── src
│   ├── main.go          # Entry point of the application
│   ├── aws
│   │   └── pricing.go   # AWS pricing query functions
│   ├── azure
│   │   └── pricing.go   # Azure pricing query functions
│   ├── gcp
│   │   └── pricing.go   # GCP pricing query functions
│   └── types
│       └── pricing.go   # Types and structures for pricing data
├── go.mod               # Module definition
├── go.sum               # Module dependency checksums
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd cloud-pricing-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run src/main.go
   ```

## Usage

The application will query the pricing APIs of AWS, Azure, and GCP and display the pricing information for compute and storage services. You can modify the source code in `src/main.go` to customize the queries or the output format.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.