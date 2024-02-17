package data

type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"recquired"`
}

type VerifyData struct {
	user *OTPData `json:"user,omitempty" validate:"recquired'`
	Code string   `json:"code,omitempty" validate:"recquired:`
}
