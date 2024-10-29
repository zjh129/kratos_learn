package data

import (
	"context"
	"kratos_learn/internal/conf"
	"kratos_learn/internal/data/ent"
	"kratos_learn/internal/data/ent/migrate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewGreeterRepo)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
}

// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data"))
	d := &Data{
		db:  entClient,
		log: log,
	}
	cleanup := func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}

func NewEntClient(c *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data/ent"))

	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
