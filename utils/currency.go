package utils

// Constants for all supported currencies
const (
	USD = "USD"
	BDT = "BDT"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, BDT:
		return true
	}
	return false
}
