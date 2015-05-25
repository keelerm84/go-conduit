package conduit

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ConduitData struct {
	Conduit Connection `json:"__conduit__"`
}

type Connection struct {
	SessionKey   string `json:"sessionKey"`
	ConnectionID int    `json:"connectionID"`
	Host         string
}

type App struct {
	Client            string `json:"client"`
	ClientVersion     int    `json:"clientVersion"`
	ClientDescription string `json:"clientDescription"`
	User              string `json:"user"`
	Host              string `json:"host"`
	AuthToken         int64  `json:"authToken"`
	AuthSignature     string `json:"authSignature"`
}

func (a *App) Connect(cert string) Connection {
	token, signature := a.generateTokenAndSignature(cert)

	a.AuthToken = token
	a.AuthSignature = signature

	appParams, _ := json.Marshal(a)

	v := url.Values{}
	v.Set("params", string(appParams))
	v.Set("output", "json")
	v.Set("__conduit__", "true")

	resp, _ := http.PostForm(a.Host+"/api/conduit.connect", v)

	result := struct {
		Result Connection `json:"result"`
	}{}

	json.NewDecoder(resp.Body).Decode(&result)

	connection := result.Result
	connection.Host = a.Host

	return connection
}

func (a *App) generateTokenAndSignature(cert string) (token int64, signature string) {
	token = time.Now().Unix()
	signature = strconv.FormatInt(token, 10) + cert

	hasher := sha1.New()
	hasher.Write([]byte(signature))

	signature = hex.EncodeToString(hasher.Sum(nil))

	return token, signature
}
