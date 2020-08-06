package oauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Config struct {
	ConsumerKey       string `json:"APP_TOKEN"`
	ConsumerSecret    string `json:"APP_SECRET"`
	AccessToken       string `json:"ACCESS_TOKEN"`
	AccessTokenSecret string `json:"ACCESS_TOKEN_SECRET"`
}

func srand(size int) string {
	var alpha = "abcdefghijkmnpqrstuvwxyz23456789"
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

func Parameters(config Config) map[string]string {
	return map[string]string{
		"oauth_consumer_key":     config.ConsumerKey,
		"oauth_nonce":            srand(13),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        strconv.FormatInt(time.Now().Unix(), 10),
		"oauth_token":            config.AccessToken,
		"oauth_version":          "1.0",
	}
}

func OauthHeader(uri string, config Config) string {
	oauthHeader := fmt.Sprintf(`OAuth realm="%s",`, url.QueryEscape(uri))

	params := Parameters(config)
	something := ""

	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		param := fmt.Sprintf(`%s="%s"`, k, params[k])
		some := fmt.Sprintf(`%s=%s`, k, params[k])
		if len(keys) != i+1 {
			param += ","
			some += "&"
		}

		oauthHeader += param
		something += some
	}

	message := BaseString("GET", uri) + url.QueryEscape(something)

	signature := AuthSignature(message, SigningKey(config.ConsumerSecret, config.AccessTokenSecret))
	oauthHeader += fmt.Sprintf(`,oauth_signature="%s"`, signature)

	return oauthHeader
}

func BaseString(method, uri string) string {
	return fmt.Sprintf("%s&%s&", method, url.QueryEscape(uri))
}

func SigningKey(appSecret, accessTokenSecret string) string {
	return fmt.Sprintf("%s&%s", url.QueryEscape(appSecret), url.QueryEscape(accessTokenSecret))
}

func AuthSignature(baseString, signingKey string) string {
	h := hmac.New(sha1.New, []byte(signingKey))
	h.Write([]byte(baseString)) //nolint

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}
