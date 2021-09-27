package pkg

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"notify-server/internal/data/ent"
	"notify/pkg/config"
	"time"
)

type connector struct {
	dsn string
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (connector) Driver() driver.Driver {
	return ocsql.Wrap(
		mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}

// NewClient new connection and start stats recorder.
func NewClient(vip *config.Config) *ent.Client {
	var viper = vip.Vip
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", viper.GetString("mysql.username"), viper.
		GetString("mysql.password"), viper.GetString("mysql.network"), viper.
		GetString("mysql.server"), viper.GetInt("mysql.port"), viper.GetString("mysql.database"))
	db := sql.OpenDB(connector{dsn})
	db.SetMaxIdleConns(viper.GetInt("mysql.maxidleconns"))
	db.SetMaxOpenConns(viper.GetInt("mysql.maxopenconns"))
	db.SetConnMaxLifetime(time.Duration(viper.GetInt("mysql.connmaxlifetime")) * time.Second)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv))
}
