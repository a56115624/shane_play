package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/teampui/pac"
	"github.com/uptrace/bun"
	"log"
)

type Handler struct {
	managerRepo RepoInterface
}

func (h *Handler) Register(app *pac.App) {
	h.managerRepo = pac.Must[RepoInterface](
		pac.Repo[RepoInterface](app, "manager"),
		"service/manager: cannot start due to no valid manager repo found")

	r := app.Router()
	r.Get("/search", h.HandleSearchID)
	r.Put("/put", h.PutDataMysql)
	r.Put("/update", h.UpdateMysqlData)
	r.Put("/insert", h.InsertData)

}
func (h *Handler) HandleSearchID(c *fiber.Ctx) error { // 撈取單筆資料

	log.Printf("Successfully connected to database")

	// team_id := c.Params("team_id")

	h.managerRepo.getMysqlData()
	// var member MemberData
	// member, err := getMysqlData(h.db)
	member, err := h.managerRepo.getMysqlData()
	if err != nil {
		return c.JSON(err)
	}
	fmt.Printf("teamMember:%+v\n", member)
	return c.JSON(member)
}
func (h *Handler) PutDataMysql(c *fiber.Ctx) error { // 插入單筆資料

	log.Printf("Successfully connected to database")

	book := new([]*MemberData)
	err := c.BodyParser(book)
	if err != nil {
		return c.JSON(err)
	}

	// err = putMysqlData(db, *book)
	err = h.managerRepo.putMysqlData(*book)
	if err != nil {
		return c.JSON(err)
	}
	fmt.Printf("成功插入:%+v\n", book)
	return c.JSON(book)

}

func (h *Handler) UpdateMysqlData(c *fiber.Ctx) error {

	log.Printf("Successfully connected to database")
	// var books  []*MemberData
	books := new([]*MemberData)
	err := c.BodyParser(books)
	if err != nil {
		return c.JSON(err)
	}
	// books := []*GolangShaneData{&book1, &book2, &book3}
	// err = updateMysqlData(h.db, *books)
	h.managerRepo.updateMysqlData(*books)
	if err != nil {
		return c.JSON(err)
	}
	fmt.Printf("成功插入:%+v\n", books)
	return c.JSON(books)
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
func putMysqlData(mysqlDb *bun.DB, books []*MemberData) error {
	ctx := context.Background()
	for i := 0; i < len(books); i++ {
		_, err := mysqlDb.NewInsert().Model(books[i]).Exec(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(books) // book id is scanned automatically
	}
	return nil
}

// 更新資料近mysql裡面
/*values := db.NewValues(&[]*Book{book1, book2})*/
func updateMysqlData(mysqlDb *bun.DB, books []*MemberData) error {
	ctx := context.Background()
	for i := 0; i < len(books); i++ {
		_, err := mysqlDb.NewUpdate().Model(books[i]).WherePK().Exec(ctx)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

// 接收到我postman 的值
func (h *Handler) InsertData(c *fiber.Ctx) error {

	log.Printf("Successfully connected to database")
	data := new(MemberData)
	if err := c.BodyParser(data); err != nil {
		return c.JSON(err)
	}
	fmt.Printf("成功插入:%+v\n", data)
	return c.JSON(data)
}
