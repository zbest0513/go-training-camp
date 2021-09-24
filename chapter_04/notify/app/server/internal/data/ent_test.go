package data

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"notify-server/internal/data/ent"
	"notify-server/internal/data/ent/user"
	"testing"
	"time"
)

func TestEntSelect(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, port, database)

	uuid := uuid.New()
	db := Open(dsn)
	defer db.Close()
	u, err := db.User.Create().SetUUID(uuid).SetEmail("24566370@qq.com").SetMobile("15012341324").SetName("张三").Save(context.TODO())

	log.Println(fmt.Sprintf("保存用户: %v ,%+v", u, err))
	exec, err := db.User.Delete().Where(user.NameEQ("张三"), user.EmailEQ("24566370@qq.com")).Exec(context.TODO())
	log.Println(fmt.Sprintf("删除用户: %v ,%+v", exec, err))
}

const (
	username        = "root"
	password        = "zhangbin123"
	network         = "tcp"
	server          = "mysql.zbest.tech"
	port            = 13306
	database        = "notify"
	maxopenconns    = 100
	maxidleconns    = 16
	connmaxlifetime = 100 * time.Second
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

// Open new connection and start stats recorder.
func Open(dsn string) *ent.Client {
	db := sql.OpenDB(connector{dsn})
	db.SetMaxIdleConns(maxidleconns)
	db.SetMaxOpenConns(maxopenconns)
	db.SetConnMaxLifetime(connmaxlifetime)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv))
}
