package cookies

import (
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/gofiber/fiber/v3"
)

const CookieNameRefreshToken = "refresh_token"

func GenerateRefreshToken(value string, delete bool) *fiber.Cookie {

	maxAge := 0
	expires := time.Now().Add(time.Minute * 60 * 24 * time.Duration(configuration.Security.RefreshTokenValidityInDays))

	if delete {
		value = ""
		hoursToSubtract := 100
		expires = time.Now().Add(-(time.Hour * time.Duration(hoursToSubtract)))
		maxAge = -1
	}

	// set refreshToken
	return &fiber.Cookie{
		Name:     CookieNameRefreshToken,
		Path:     configuration.CONST_AUTHCOOKIES_PATH_API_V1,
		Value:    value,
		HTTPOnly: true,
		Expires:  expires,
		MaxAge:   maxAge,
		// Secure:   true,
		SameSite: "Strict",
	}
}
