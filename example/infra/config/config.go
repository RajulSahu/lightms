package config

import (
	lightms "github.com/RajulSahu/lightms/v3"
	"github.com/RajulSahu/lightms/v3/example/core/impl"
	"github.com/RajulSahu/lightms/v3/example/core/port"
	"github.com/RajulSahu/lightms/v3/example/infra/config/prop"
	"github.com/RajulSahu/lightms/v3/example/infra/primary"
	"github.com/RajulSahu/lightms/v3/example/infra/secondary"
)

type PropsConfig struct {
	DatabaseProp *prop.DatabaseProp
	PubSubProp   *prop.PubSubProp
}

func (c *PropsConfig) Subs() {
	lightms.AddInstance(primary.NewJohnDoeSubscription)
}

var _ = lightms.AddInstance(primary.NewJaneDoeSubscription)

var _ = lightms.AddInstance(impl.NewJohnDoeImpl).
	AndInjections().
	WithInjection("postgres.PersistencePort", lightms.TOF[port.PersistencePort]())

var _ = lightms.AddInstance(secondary.NewPersistence1Port).
	WithAlias("postgres.PersistencePort", lightms.TOF[port.PersistencePort]())
