package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/teampui/pac"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"os"
)

const (
	mysqlDsn = "shane:GKbCoMubLMQ6o@tcp(sgpdb.itlab.tw:8889)/shane?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	app := pac.NewApp(
		pac.ListenPortFromEnv(":7777"),   // 如果環境變數裡沒設定的話，預設 :7777
		pac.UseLogger(),                  // 使用請求記錄器
		ProvideMysqlDB(os.Getenv("DSN")), // 使用 BunDB 作為資料庫層, // 使用 BunDB 作為資料庫層
		// redis.ProvideSession(redis.SessionConfig{
		// 	ClientKeystore: "cookie:942",
		// 	RedisURL:       os.Getenv("REDIS_DSN"),
		// 	Expiration:     24 * time.Hour,
		// }),
	)

	app.Add(&PostgresRepo{})
	app.Add(&Handler{})

	app.Start()
}
func ProvideMysqlDB(dsn string) pac.AppOption {
	if dsn == "" {
		panic("pac/bundb: cannot start, missed DSN settings")
	}

	return func(a *pac.App) {
		sqldb, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		a.Services.Add("db", bun.NewDB(sqldb, mysqldialect.New()))
	}
}
