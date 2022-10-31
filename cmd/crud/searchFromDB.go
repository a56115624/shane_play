package crud

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

// 記得資料庫的欄位都要大寫
type MemberData struct {
	// bun.BaseModel `bun:"table:cpbl_member"`
	// ID            int    `json:"id"`
	// Team_id       int    `json:"team_id"`
	// Name          string `json:"name"`
	// Ig_url        string `json:"ig_url"`
	// Status        int    `json:"status"`

	bun.BaseModel `bun:"table:Golang_shane_data"`
	Id            int64 `bun:",pk,autoincrement"`
	Description   string
}

func SearchFromDB() {
	// fmt.Println("go mysql")
	const (
		MaxLifetime  int = 10
		MaxOpenConns int = 10
		MaxIdleConns int = 10
	)
	db, err := sql.Open("mysql", dsn("shane")) // 建立連線

	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second) // 設置了連接可重用的最大時間
	db.SetMaxOpenConns(MaxOpenConns)                                // 設定最大連線數
	db.SetMaxIdleConns(MaxIdleConns)                                // 設定閒置連線量

	if err != nil {
		panic(err.Error()) // 在 Go 語言沒有拋出例外的機制，如果在執行過程需要中斷目前的流程，則可以透過 panic 機制
	}

	defer db.Close() // defer 是用來宣告 function 結束前的動作。

	fmt.Println("successfully connected to mysql db")

	// results, err := db.Query("select id, title, performer, vd_number from dmm where id = ?", 1)

	// if err != nil {
	// 	panic(err.Error())
	// }

}

func serchMember(c *fiber.Ctx) error {
	fmt.Print("called post api person\n")
	data := new(MemberData) // 同等於 &Video{} 寫法
	if err := c.BodyParser(data); err != nil {
		fmt.Printf("bodyparser error: %v\n", err)
		return c.SendString("error")
	}
	return c.JSON(data)

	// return c.SendString("ok")
}
