package cookies

import (
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/gofiber/fiber/v3"
)

const CookieNameAuthProvider = "auth_provder"

func GenerateAuthProviderCookie(value string, delete bool) *fiber.Cookie {
	maxAge := 0
	maxDurationToLogin := 10
	expires := time.Now().Add(time.Minute * time.Duration(maxDurationToLogin))

	if delete {
		value = ""
		hoursToSubtract := 100
		expires = time.Now().Add(-(time.Hour * time.Duration(hoursToSubtract)))
		maxAge = -1
	}

	return &fiber.Cookie{
		Name:    CookieNameAuthProvider,
		Path:    configuration.CONST_OAUTHCOOKIES_PATH_API_V1,
		Value:   value,
		Expires: expires,
		MaxAge:  maxAge,
	}
}
