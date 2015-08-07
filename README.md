# go-steam-web
Authentication with steam via https://steamcommunity.com website.

## Usage
Basic authentication
```
session := steamweb.NewSession()
response, err := session.Login(&steamweb.NewCredentials("login","password","steamid64"))
//Error checking
if response.LoginComplete {
    //We are logged in :>
}
```
With steamguard enabled you have to extract steamMachineAuth cookie.
Chrome:
1. Go to https://steamcommunity.com
2. Click F12 to open developer console
3. Click Resources and go to cookies/steamcommunity.com
4. Look for steamMachineAuth<steamid> and copy value
5. Supply the value to Credentials.Token ``` credentials.Token="tokengoeshere"```
