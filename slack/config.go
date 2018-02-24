package slack

type Config struct {
	ClientID          string `json:"clientId"`
	ClientSecret      string `json:"clientSecret"`
	VerificationToken string `json:"verificationToken"`
}
