package main

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type JWKey struct {
	Kty string `json:"kty"`
	Use string `json:"use,omitempty"`
	Kid string `json:"kid,omitempty"`
	Alg string `json:"alg,omitempty"`

	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Y   string `json:"y,omitempty"`
	D   string `json:"d,omitempty"`
	N   string `json:"n,omitempty"`
	E   string `json:"e,omitempty"`
	K   string `json:"k,omitempty"`
}

type JWTHeader struct {
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

type JWKeys struct {
	Keys []JWKey `json:"keys"`
}

func main() {
	// sample token string taken from the New example
	tokenString := "eyJraWQiOiI4NkQ4OEtmIiwiYWxnIjoiUlMyNTYifQ.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoic2VydmVyLmNhcnNjaG9vbC5mYW5jeS5jb20iLCJleHAiOjE2MDg4MTc5MjEsImlhdCI6MTYwODczMTUyMSwic3ViIjoiMDAwNzIwLmEwZDg4NDUyOTdkYTQ1OWJiY2RlMjYzMDZkMDJlZGIwLjA5NTkiLCJhdF9oYXNoIjoiRHo0a3JjQ2cxaUdrSVVjamtuUnhhZyIsImF1dGhfdGltZSI6MTYwODczMTQ5Miwibm9uY2Vfc3VwcG9ydGVkIjp0cnVlfQ.bzqWKxbowgYEkXQBItgzaMEYin2hygHRyhMzrVa1jR6e0Hb5WxHZfvtZ1iYEegsn7pv0qcTdwDdh5w0VZKgsxIjFtpmNCj8TWCOTO1Y5nM7C6Fm_WeEG4mgvSuJ1jWN9pbCs88VmGy0kgK0VktoyAdrMVebLa7yxbN6ffeqcDWy7QJiOgXOmavDdYnrPy2FJlahrJE1ID-HUK1VVJ0XAYq_4J315T0qJhbeP7Sc6eDcgzpHzJz8g2OwFQZeOCG1lHkgiPIEV41bDnetyTKCG1KzxGhLLC3UwlK3k-nMawk_-GWcCY6vwXKQvPMvLmOdroXDu7GYO_2Eoe5OQoQ0zSw"
	chunk := strings.Split(tokenString, ".")
	data, err := base64.RawURLEncoding.DecodeString(chunk[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	header := new(JWTHeader)
	if err := json.Unmarshal(data, header); err != nil {
		log.Fatal(err)
		return
	}
	jwkeys, err := GetAuthKeys()
	if err != nil {
		log.Fatal(err)
		return
	}
	pk, err := extractPublicKeyFromJWK(*jwkeys, "RSA", header.Kid)
	if err != nil {
		log.Fatal(err)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return pk, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}

func GetAuthKeys() (*JWKeys, error) {
	url := "https://appleid.apple.com/auth/keys"
	cli := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("GET %s Response: %d", resp.StatusCode)
	}
	row := new(JWKeys)
	if err := json.NewDecoder(resp.Body).Decode(row); err != nil {
		return nil, err
	}
	return row, nil
}

func extractPublicKeyFromJWK(jwks JWKeys, outType, kid string) (*rsa.PublicKey, error) {
	for _, jwk := range jwks.Keys {
		if jwk.Kty != "RSA" {
			return nil, errors.Errorf("invalid key type: %s", jwk.Kty)
		}
		// decode the base64 bytes for n
		nb, err := base64.RawURLEncoding.DecodeString(jwk.N)
		if err != nil {
			return nil, err
		}
		e := 0
		// The default exponent is usually 65537, so just compare the
		// base64 for [1,0,1] or [0,1,0,1]
		if jwk.E == "AQAB" || jwk.E == "AAEAAQ" {
			e = 65537
		} else {
			// need to decode "e" as a big-endian int
			return nil, errors.Errorf("need to deocde e: %s", jwk.E)
		}

		pk := &rsa.PublicKey{
			N: new(big.Int).SetBytes(nb),
			E: e,
		}
		return pk, nil
	}
	return nil, errors.Errorf("invalid jwt")
}
