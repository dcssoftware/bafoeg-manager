package cookies

import (
	"net/http"
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/gofiber/fiber/v3"
)

const CookieNameJwtToken = "jwt"

func GenerateJwtToken(value string, delete bool) *fiber.Cookie {
	maxAge := 0
	expires := time.Now().Add(time.Minute * time.Duration(configuration.Security.JwtTokenValidInMinutes))

	if delete {
		value = ""
		hoursToSubtract := 100
		expires = time.Now().Add(-(time.Hour * time.Duration(hoursToSubtract)))
		maxAge = -1
	}

	return &fiber.Cookie{
		Name:    CookieNameJwtToken,
		Path:    "/",
		Value:   value,
		Expires: expires,
		MaxAge:  maxAge,
	}
}

func GenerateHttpJwtToken(value string, delete bool) *http.Cookie {
	maxAge := 0
	expires := time.Now().Add(time.Minute * time.Duration(configuration.Security.JwtTokenValidInMinutes))

	if delete {
		value = ""
		hoursToSubtract := 100
		expires = time.Now().Add(-(time.Hour * time.Duration(hoursToSubtract)))
		maxAge = -1
	}

	return &http.Cookie{
		Name:    CookieNameJwtToken,
		Path:    "/",
		Value:   value,
		Expires: expires,
		MaxAge:  maxAge,
	}
}
