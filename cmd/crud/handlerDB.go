package crud

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type Handler struct {
	db *bun.DB
}

type HandlerConfig struct {
	User   string
	Pass   string
	DbName string
	Host   string
	Port   string
}

// func Newhandler(config HandlerConfig) Handler { // 將db user pass ... 將成參數可以帶入
// 	return Handler{db}
// }

func Newhandler(db *bun.DB) Handler {
	return Handler{db}
}

func (h *Handler) HandleSearchID(c *fiber.Ctx) error { // 撈取單筆資料

	db := h.db
	log.Printf("Successfully connected to database")

	// team_id := c.Params("team_id")

	// var member MemberData
	member, err := getMysqlData(db)
	if err != nil {
		return c.JSON(err)
	}
	fmt.Printf("teamMember:%+v\n", member)
	return c.JSON(member)
}
func (h *Handler) PutDataMysql(c *fiber.Ctx) error { // 撈取單筆資料

	db := h.db
	log.Printf("Successfully connected to database")

	// team_id := c.Params("team_id")

	// var member MemberData
	book, err := putMysqlData(db)
	if err != nil {
		return c.JSON(err)
	}
	fmt.Printf("成功插入:%+v\n", book)
	return c.JSON(book)
}

// 拿到mysql裡面的資料
func getMysqlData(mysqlDb *bun.DB) ([]MemberData, error) {
	ctx := context.Background()

	var blocks []MemberData
	err := mysqlDb.NewSelect().Model(&blocks).Order("ID").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

// 插入資料進mysqlDb
func putMysqlData(mysqlDb *bun.DB) (MemberData, error) {
	book := MemberData{Id: 5, Description: "我好想偷懶"}
	ctx := context.Background()
	_, err := mysqlDb.NewInsert().Model(&book).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(book) // book id is scanned automatically
	return book, nil
}
