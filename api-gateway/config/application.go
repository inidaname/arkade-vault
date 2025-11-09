package config

import (
	"sync"

	"github.com/inidaname/arkade-vault/api-gateway/pkg/types"
)

var (
	instance *types.Application
	once     sync.Once
)

func CreateApplication() *types.Application {
	once.Do(func() {})
}