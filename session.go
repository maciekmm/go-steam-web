package steamweb

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/maciekmm/go-steam-web/utils"
)

var (
	IncorrectPassword  = errors.New("Incorrect password")
	InvalidToken       = errors.New("Invalid token")
	InvalidName        = errors.New("Invalid name")
	errorGettingKey    = "Error while getting key, Cause: %s"
	ServersUnavailable = errors.New("Services unavailable")
)

type Session struct {
	httpClient *http.Client
	key        SteamPublicKey
}

func (sess *Session) GetSessionID() string {
	ur, _ := url.Parse("https://steamcommunity.com")
	cookies := sess.httpClient.Jar.Cookies(ur)
	for _, v := range cookies {
		if v.Name == "sessionid" {
			return v.Value
		}
	}
	return ""
}

type SteamPublicKey struct {
	PublicKeyExp string `json:"publickey_exp,omitempty"`
	PublicKeyMod string `json:"publickey_mod,omitempty"`
	SteamID      uint64 `json:"steamid,string,omitempty"`
	Success      bool   `json:"success"`
	Timestamp    uint64 `json:"timestamp,string,omitempty"`
	TokenGID     string `json:"token_gid,omitempty"`
}

func (spk SteamPublicKey) modulus() (*big.Int, error) {
	by, er := hex.DecodeString(spk.PublicKeyMod)
	if er != nil {
		return nil, er
	}
	bi := big.NewInt(0)
	return bi.SetBytes(by), nil
}

func (spk SteamPublicKey) exponent() (int64, error) {
	return strconv.ParseInt(spk.PublicKeyExp, 16, 0)
}

func NewSession() (*Session, error) {
	jar, _ := cookiejar.New(nil)
	sess := &Session{
		httpClient: &http.Client{
			Jar: jar,
		},
	}
	resp, err := sess.httpClient.Do(sess.newRequest("GET", "https://steamcommunity.com", nil))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return sess, nil
}

func (sess *Session) Login(credentials Credentials) (*LoginResponse, error) {
	key, err := sess.getRSA(credentials.Username)
	if err != nil {
		return nil, err
	}
	encryptedPassword, err := sess.encryptPassword(credentials.Password, key)
	if err != nil {
		return nil, err
	}
	credentials.Password = encryptedPassword
	credentials.RSATimeStamp = strconv.FormatUint(key.Timestamp, 10)
	credentials.DoNotCache = strconv.FormatInt(time.Now().Unix(), 10)
	req := sess.newRequest("POST", "https://steamcommunity.com/login/dologin/", strings.NewReader(utils.ToURLValues(&credentials).Encode()))
	fmt.Println(utils.ToURLValues(&credentials).Encode())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if credentials.Token != "" {
		//Workaround for steamguard
		req.AddCookie(&http.Cookie{
			Name:   fmt.Sprintf("steamMachineAuth%s", credentials.SteamID),
			Value:  credentials.Token,
			Path:   "/",
			Domain: ".steamcommunity.com",
			Secure: true,
		})
	}
	resp, err := sess.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	loginr := new(LoginResponse)
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf.Bytes(), loginr)
	return loginr, err
}

func (sess *Session) newRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("User-Agent", "Mozilla/5.0 ;Windows NT 6.1; WOW64; Trident/7.0; rv:11.0; like Gecko")
	if err != nil {
		panic(err.Error())
	}
	return req
}

func (sess *Session) encryptPassword(password string, spk *SteamPublicKey) (string, error) {
	pk := new(rsa.PublicKey)
	exp, err := spk.exponent()
	if err != nil {
		return "", fmt.Errorf(errorGettingKey, err.Error())
	}
	pk.E = int(exp)
	if pk.N, err = spk.modulus(); err != nil {
		return "", fmt.Errorf(errorGettingKey, err.Error())
	}
	out, err := rsa.EncryptPKCS1v15(rand.Reader, pk, []byte(password))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(out), nil
}

func (sess *Session) getRSA(username string) (*SteamPublicKey, error) {
	reqParams := make(url.Values)
	reqParams.Add("username", username)
	reqParams.Add("donotcache", strconv.FormatUint(uint64(time.Now().Unix()), 10))
	req := sess.newRequest("POST", "https://steamcommunity.com/login/getrsakey/", strings.NewReader(reqParams.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := sess.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(errorGettingKey, err.Error())
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, fmt.Errorf(errorGettingKey, err.Error())
	}
	spk := new(SteamPublicKey)
	if json.Unmarshal(buf.Bytes(), spk) != nil {
		return nil, fmt.Errorf(errorGettingKey, err.Error())
	}
	if !spk.Success {
		return nil, errors.New(errorGettingKey)
	}
	return spk, nil
}