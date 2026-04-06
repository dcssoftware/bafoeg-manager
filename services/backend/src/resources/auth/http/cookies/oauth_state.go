package cookies

import (
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/gofiber/fiber/v3"
)

const CookieNameOauthState = "auth_provder"

func GenerateOauthStateCookie(value string, delete bool) *fiber.Cookie {
	maxAge := 0
	maxDurationInMinutes := 15

	expires := time.Now().Add(time.Minute * time.Duration(maxDurationInMinutes))

	if delete {
		value = ""
		hoursToSubtract := 100
		expires = time.Now().Add(-(time.Hour * time.Duration(hoursToSubtract)))
		maxAge = -1
	}

	return &fiber.Cookie{
		Name:    CookieNameAuthProvider,
		Value:   value,
		Path:    configuration.CONST_OAUTHCOOKIES_PATH_API_V1,
		Expires: expires,
		MaxAge:  maxAge,
	}
}
