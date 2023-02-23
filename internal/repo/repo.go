package repo

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/realHoangHai/authenticator/config"
	"github.com/realHoangHai/authenticator/internal/common"
	"github.com/realHoangHai/authenticator/internal/ent"
	"github.com/realHoangHai/authenticator/internal/ent/migrate"
)

var ProviderRepoSet = wire.NewSet(NewRepo)
var _ IRepo = (*Repo)(nil)

type Repo struct {
	*ent.Client
	ctx context.Context
}

func NewRepo(ctx context.Context) (IRepo, error) {
	driver := config.C.DBConfig.Driver
	source := config.C.DBConfig.DSN()

	conn, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}

	client := ent.NewClient(ent.Driver(conn))
	opts := []schema.MigrateOption{
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	}

	if config.C.DBConfig.Migration {
		if err := client.Schema.Create(ctx, opts...); err != nil {
			defer func() {
				_ = client.Close()
			}()
			return nil, err
		}
	}

	if config.C.AppConfig.Mode == common.DebugMode {
		client = client.Debug()
	}

	return &Repo{
		Client: client,
		ctx:    ctx,
	}, nil
}
