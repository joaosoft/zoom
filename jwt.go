package zoom

import (
	"fmt"
	"github.com/joaosoft/auth-types/jwt"
	"net/http"
	"time"
)

func jwtToken(key string, secret string) (string, error) {
	claims := jwt.Claims{
		jwt.ClaimsIssuedAtKey: key,
		jwt.ClaimsExpireAtKey: time.Now().Local().Add(time.Second * time.Duration(5000)).Unix(),
	}

	return jwt.New(jwt.SignatureHS256).Generate(claims, secret)
}

func (c *Client) addRequestAuth(req *http.Request, err error) (*http.Request, error) {
	if err != nil {
		return nil, err
	}

	// establish JWT token
	ss, err := jwtToken(c.config.ApiKey, c.config.ApiSecret)
	if err != nil {
		return nil, err
	}

	// set JWT Authorization header
	req.Header.Add(headerAuthorization, fmt.Sprintf(headerValueBearer, ss))

	return req, nil
}
