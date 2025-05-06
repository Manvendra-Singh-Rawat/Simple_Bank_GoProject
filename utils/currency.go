package utils

const (
	USD = "USD"
	INR = "INR"
)

// Return true if input currency is supported
func isSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR:
		return true
	}
	return false
}
