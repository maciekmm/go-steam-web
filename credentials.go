package steamweb

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

func NewCredentials(username, password, steamid string) Credentials {
	return Credentials{
		Username: username,
		Password: password,
		SteamID:  steamid,
	}
}
