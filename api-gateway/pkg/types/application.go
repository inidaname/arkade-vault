package types

import (
	"log/slog"

	// "github.com/inidaname/mosque/api_gateway/pkg/messages/mailer"
	cache "github.com/inidaname/arkade-vault/api-gateway/pkg/store/cache"
)

type Application struct {
	Config        Config
	Logger        *slog.Logger
	Authenticator Authenticator
	Cache cache.CacheService
}