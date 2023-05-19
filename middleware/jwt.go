package middleware

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type jwks struct {
	Keys []jsonWebKeys `json:"keys"`
}

type jsonWebKeys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func getPemCert(token *jwt.Token) (*rsa.PublicKey, error) {
	var jwk jsonWebKeys
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	resp, err := http.Get(fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks = &jwks{}
	err = json.NewDecoder(resp.Body).Decode(jwks)

	if err != nil {
		return nil, err
	}

	for _, k := range jwks.Keys {
		if token.Header["kid"] == k.Kid {
			jwk = k
			break
		}
	}

	if jwk.Kid == "" {
		return nil, errors.New("unable to find appropriate key")
	}

	decodedN, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, err
	}

	decodedE, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, err
	}

	n := big.NewInt(0).SetBytes(decodedN)
	e := big.NewInt(0).SetBytes(decodedE).Int64()

	return &rsa.PublicKey{
		N: n,
		E: int(e),
	}, nil
}

// jwtMiddleware validates JWT from Authorization header in the format `Bearer {token}`.
func JwtMiddleware(c *fiber.Ctx) error {
	godotenv.Load(".env")
	auth0Audience := os.Getenv("AUTH0_AUDIENCE")
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing Authorization Header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Authorization Header - no parts")
	}

	if headerParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Authorization Header - no Bearer")
	}

	tokenString := headerParts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		rsaPublicKey, err := getPemCert(token)
		if err != nil {
			return nil, err
		}

		return rsaPublicKey, nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Token - error")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		aud := claims["aud"]
		if aud == nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid audience")
		}

		foundAud := false

		switch aud := aud.(type) {
		case []interface{}:
			for _, a := range aud {
				if a == auth0Audience {
					foundAud = true
					break
				}
			}
		case string:
			if aud == auth0Audience {
				foundAud = true
			}
		default:
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid audience format")
		}
		if !foundAud {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid audience")
		}

		c.Locals("user", claims)
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Token Claims")
	}

	return c.Next()
}
