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

	app.Get("/:team_id", h.HandleSearchTeamID)

	app.Listen(":3000")
}

// package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"tseng-api/cmd/crud"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gofiber/fiber"
// 	"github.com/uptrace/bun"
// 	"github.com/uptrace/bun/dialect/mysqldialect"
// 	"github.com/uptrace/bun/extra/bundebug"
// )

// const (
// 	mysqlDsn = "shane:GKbCoMubLMQ6o@tcp(sgpdb.itlab.tw:8889)/shane?charset=utf8mb4&parseTime=True&loc=Local"
// )

// var (
// 	mysqlDb *bun.DB
// )

// type Handler struct {
// 	db *bun.DB
// }

// type MemberData struct {
// 	bun.BaseModel `bun:"table:cpbl_member"`
// 	ID            int    `json:"id"`
// 	Team_id       int    `json:"team_id"`
// 	Name          string `json:"name"`
// 	Ig_url        string `json:"ig_url"`
// 	Status        int    `json:"status"`

// 	// bun.BaseModel `bun:"table:Golang_shane_data"`
// 	// Id            int64 `bun:",pk"`
// 	// Description   string
// }

// func Newhandler(db *bun.DB) Handler {
// 	return Handler{db}
// }

// func (h *Handler) HandleGetMysqlData(c *fiber.Ctx) error { // 撈取單筆資料

// 	// db := h.db
// 	// log.Printf("Successfully connected to database")

// 	team_id := c.Params("team_id")

// 	var member MemberData

// 	ctx := context.Background()

// 	var blocks []crud.MemberData
// 	err := mysqlDb.NewSelect().Model(&blocks).Where(team_id).Scan(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	// return blocks, nil

// 	// row := db.QueryRow("select name, ig_url, status from cpbl_member where team_id = ?", team_id)
// 	// //Scan對應的欄位與select語法的欄位順序一致
// 	// if err := row.Scan(&member.Name, &member.Ig_url, &member.Status); err != nil {
// 	// 	fmt.Printf("scan failed, err:%v\n", err)

// 	// }
// 	fmt.Printf("teamMember:%+v\n", member)
// 	return c.JSON(blocks)
// }

// // 連線進資料庫mysqlDb
// func init() {
// 	sqldb, err := sql.Open("mysql", mysqlDsn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	mysqlDb = bun.NewDB(sqldb, mysqldialect.New())
// 	// 下面這行是要顯示出sql的語法
// 	mysqlDb.AddQueryHook(bundebug.NewQueryHook())

// }

// // func DbConnection() (*sql.DB, error) { //建立db連線
// // 	// db, err := sql.Open("mysql", dsn(""))
// // 	db, err := sql.Open("mysql", dsn(db_name))
// // 	if err != nil {
// // 		log.Printf("Error %s when opening DB\n", err)
// // 		return nil, err
// // 	}

// // 拿到mysql裡面的資料
// func getMysqlData() ([]crud.MemberData, error) {
// 	ctx := context.Background()

// 	var blocks []crud.MemberData
// 	err := mysqlDb.NewSelect().Model(&blocks).Order("ID").Scan(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return blocks, nil
// }

// // 執行
// func main() {

// 	blocks, err := getMysqlData()

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(blocks)
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error { //Middleware 設計用於更改請求或響應的函數稱爲中間件函數。
// 		fmt.Printf("hello world")
// 		return c.SendString("welcome cpbl")
// 	})

// 	app.Get("/:team_id", getMysqlData())
// 	// app.Get("/:performer", h.HandleSearchPerformer)
// 	// app.Post("/video", h.HandlePostApi)
// 	app.Listen(":3000")
// }
