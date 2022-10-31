package main

import (
	"database/sql"
	"fmt"
	"shane_play/cmd/crud"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

const (
	mysqlDsn = "shane:GKbCoMubLMQ6o@tcp(sgpdb.itlab.tw:8889)/shane?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {

	sqldb, err := sql.Open("mysql", mysqlDsn)
	if err != nil {
		panic(err)
	}

	mysqlDb := bun.NewDB(sqldb, mysqldialect.New())
	// 下面這行是要顯示出sql的語法
	mysqlDb.AddQueryHook(bundebug.NewQueryHook())

	h := crud.Newhandler(mysqlDb) // h = {0xc0001835f0}

	// SelectID(db *sql.DB, DataID int)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error { //Middleware 設計用於更改請求或響應的函數稱爲中間件函數。
		fmt.Printf("hello world")
		return c.SendString("welcome cpbl")
	})

	// 在Golang的世界裡,要讓這個函式可以被引用,一定要用大寫
	app.Get("/search", h.HandleSearchID)
	app.Put("/put", h.PutDataMysql)
	app.Put("/update", h.UpdateMysqlData)
	app.Put("/insert", h.InsertData)
	app.Listen(":3000")
}
