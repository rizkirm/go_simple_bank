package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

//IsSupportCurrency returns true if the currency is supported
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
