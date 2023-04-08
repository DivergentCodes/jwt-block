// Package blocklist manages inspecting tokens and interacting with the cache
//
// Includes JWT parsing, validation, and verification. Also includes cache management and lookup functions.
package blocklist

import (
	"context"
	"errors"

	"github.com/divergentcodes/jwt-block/internal/core"
)

var (
	redisContext = context.TODO()

	SuccessTokenBlocked   = "Token blocked"
	SuccessTokenUnblocked = "Token unblocked"
	SuccessTokenExists    = "Token exists in the blocklist"
	SuccessTokenNotExists = "Token is not in the blocklist"
	SuccessTokenIsAllowed = "JWT is allowed"
	SuccessTokenIsBlocked = "JWT is blocked"

	ErrMisconfiguredCache = errors.New("server cache configuration error")
	ErrNoExpForTTL        = errors.New("token has no set expiration")
)

func init() {
	core.InitConfigDefaults()
}
