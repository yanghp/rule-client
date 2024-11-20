package client

import (
	kconf "github.com/yanghp/rule-client/config"
	"github.com/yanghp/rule-client/dto"
)

type Tenanter interface {
	Payload(pl *dto.Payload) (kconf.ConfigReader, error)
	Tenant(tenant *kconf.Tenant) (kconf.ConfigReader, error)
}

type Engine interface {
	Of(ruleName string) Tenanter
}
