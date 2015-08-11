package steamweb

//Credentials represents a request parameters for logging in proccess
type Credentials struct {
	Password          string `url:"password"`
	Username          string `url:"username"`
	TwoFactorCode     string `url:"twofactorcode"`
	EmailAuth         string `url:"emailauth"`
	LoginFriendlyName string `url:"loginfriendlyname"`
	CaptchaGID        int    `url:"captchagid"`
	CaptchaText       string `url:"captcha_text"`
	Token             string `url:"-"`
	SteamID           string `url:"-"`
	EmailSteamID      string `url:"emailsteamid"`
	RSATimeStamp      string `url:"rsatimestamp"`
	RememberLogin     bool   `url:"remember_login"`
	DoNotCache        string `url:"donotcache"`
}

//NewCredentials creates minimum required parameters for logging in proccess
func NewCredentials(username, password, steamid string) Credentials {
	return Credentials{
		Username: username,
		Password: password,
		SteamID:  steamid,
	}
}
