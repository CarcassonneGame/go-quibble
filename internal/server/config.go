package server

import (
	"fmt"

	"github.com/quibbble/go-quibbble/internal/datastore"
	"github.com/quibbble/go-quibbble/pkg/http"
	"github.com/quibbble/go-quibbble/pkg/logger"
)

type Config struct {
	Environment string
	Log         logger.Config
	Router      http.RouterConfig
	Server      http.ServerConfig
	Datastore   datastore.DatastoreConfig
	Network     NetworkOptions
}

func (c Config) Str() string {
	c.Datastore.Cockroach.Username = "***"
	c.Datastore.Cockroach.Password = "***"
	return fmt.Sprintf("%+v", c)
}
