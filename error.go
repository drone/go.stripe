package stripe

// Stripe-provided error codes and types
// See https://stripe.com/docs/api#errors
const (
	ErrTypeInvalidRequest = "invalid_request_error"
	ErrTypeAPI            = "api_error"
	ErrTypeCard           = "card_error"

	ErrCodeIncorrectNumber    = "incorrect_number"
	ErrCodeInvalidNumber      = "invalid_number"
	ErrCodeInvalidExpiryMonth = "invalid_expiry_month"
	ErrCodeInvalidExpiryYear  = "invalid_expiry_year"
	ErrCodeInvalidCVC         = "invalid_cvc"
	ErrCodeExpiredCard        = "expired_card"
	ErrCodeIncorrectCVC       = "incorrect_cvc"
	ErrCodeIncorrectZIP       = "incorrect_zip"
	ErrCodeCardDeclined       = "card_declined"
	ErrCodeMissing            = "missing"
	ErrCodeProcessingError    = "processing_error"
)

// Error encapsulates an error returned by the Stripe REST API.
// Detail.Code and Detail.Param may be empty.
type Error struct {
	Code   int
	Detail struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Code    string `json:"code,omitempty"`
		Param   string `json:"param,omitempty"`
	} `json:"error"`
}

func (e *Error) Error() string {
	return e.Detail.Message
}
