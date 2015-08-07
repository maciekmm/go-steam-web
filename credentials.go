package steamweb

//Credentials represents a request parameters for logging in proccess
type Credentials struct {
	Password          string `uval:"password"`
	Username          string `uval:"username"`
	TwoFactorCode     string `uval:"twofactorcode"`
	EmailAuth         string `uval:"emailauth"`
	LoginFriendlyName string `uval:"loginfriendlyname"`
	CaptchaGID        int    `uval:"captchagid"`
	CaptchaText       string `uval:"captcha_text"`
	Token             string `uval:"-"`
	SteamID           string `uval:"-"`
	EmailSteamID      string `uval:"emailsteamid"`
	RSATimeStamp      string `uval:"rsatimestamp"`
	RememberLogin     bool   `uval:"remember_login"`
	DoNotCache        string `uval:"donotcache"`
}

//NewCredentials creates minimum required parameters for logging in proccess
func NewCredentials(username, password, steamid string) Credentials {
	return Credentials{
		Username: username,
		Password: password,
		SteamID:  steamid,
	}
}
