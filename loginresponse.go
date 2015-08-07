package steamweb

type LoginResponse struct {
	Success           bool   `json:"success"`
	Message           string `json:"message,omitempty"`
	RequiresTwoFactor bool   `json:"requires_twofactor"`

	//If requires steamguard
	EmailAuthNeeded bool   `json:"emailauth_needed,omitempty"`
	EmailDomain     string `json:"emaildomain"`
	EmailSteamID    string `json:"emailsteamid"`

	//When logged in
	LoginComplete      bool               `json:"login_complete,omitempty"`
	TransferURL        string             `json:"transfer_url,omitempty"`
	TransferParameters TransferParameters `json:"transfer_parameters,omitempty"`
}

type TransferParameters struct {
	SteamID       string `json:"steamid"`
	Token         string `json:"token"`
	Auth          string `json:"auth"`
	RememberLogin bool   `json:"remember_login"`
	TokenSecure   string `json:"token_secure"`
}
