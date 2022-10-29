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

func (h *Handler) HandleSearchTeamID(c *fiber.Ctx) error { // 撈取單筆資料

	db := h.db
	log.Printf("Successfully connected to database")

	// team_id := c.Params("team_id")

	// var member MemberData
	member, err := getMysqlData(db)
	if err != nil {
		return c.JSON(err)
	}
	// row := db.QueryRow("select name, ig_url, status from cpbl_member where team_id = ?", team_id)
	// //Scan對應的欄位與select語法的欄位順序一致
	// if err := row.Scan(&member.Name, &member.Ig_url, &member.Status); err != nil {
	// 	fmt.Printf("scan failed, err:%v\n", err)

	// }
	fmt.Printf("teamMember:%+v\n", member)
	return c.JSON(member)
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

// func (h *Handler) HandleSearchPerformer(c *fiber.Ctx) error { // 透過番號撈出b資料

// 	db := h.db
// 	log.Printf("Successfully connected to database")

// 	performer := c.Params("performer")

// 	var video Data

// 	row := db.Query("select id, title, performer, vd_number, detail, maker, dmm_url, tag, series, created_at, updated_at from dmm where performer = ?", performer)
// 	//Scan對應的欄位與select語法的欄位順序一致
// 	if err := row.Scan(&video.ID, &video.Title, &video.Performer, &video.Vd_number, &video.Maker,
// 		&video.Detail, &video.Dmm_url, &video.Tag, &video.Series, &video.Created_at, &video.Updated_at); err != nil {
// 		fmt.Printf("scan failed, err:%v\n", err)

// 	}
// 	fmt.Printf("video:%+v\n", video)
// 	return c.JSON(video)
// }

// func SelectDatasByID(db *sql.DB, idStart int, idEnd int) ([]Data, error) { // 選擇多行 + 多條件
// log.Printf("Getting id, title, performer by ID range")
// query := `select id, title, performer from dmm where id >= ? && id <= ?;` // 選取id範圍的title, performer...
// stmt, err := db.Prepare(query)
// if err != nil {
// 	log.Printf("Error %s when preparing SQL statement", err)
// 	return []Data{}, err
// }
// defer stmt.Close()
// rows, err := stmt.Query(idStart, idEnd)
// if err != nil {
// 	return []Data{}, err
// 	}
// 	defer rows.Close()
// var datas = []Data{}
// 	for rows.Next() { // 逐筆印出
// 		var vid Data
// 		if err := rows.Scan(&vid.ID, &vid.Title, &vid.Performer); err != nil {
// 			return []Data{}, err
// 		}
// 		datas = append(datas, vid)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return []Data{}, err
// 	}
// 	fmt.Println(datas)
// 	return datas, nil
// }

// func (h *Handler) HandlePostApi(c *fiber.Ctx) error {
// 	db := h.db

// 	fmt.Print("called post api person\n")

// 	video := new(Data) // &Data{}

// 	if err := c.BodyParser(video); err != nil {
// 		fmt.Printf("bodyparser error: %v\n", err)
// 		return c.SendString("error")
// 	}

// 	err := Insert(db, *video)
// 	if err != nil {
// 		log.Printf("Insert product failed with error %s", err)
// 		return err
// 	}

// 	fmt.Printf("video:%+v\n", video)
// 	return c.JSON(video)

// }

// func handleSearchVdNumber(c *fiber.Ctx) error { // 透過番號撈出b資料

// 	db, err := crud.DbConnection() // 連線db
// 	if err != nil {
// 		log.Printf("Error %s when getting db connection", err)
// 		return err
// 	}
// 	defer db.Close()

// 	log.Printf("Successfully connected to database")

// 	vd_number := c.Params("vd_number")

// 	var video crud.Data

// 	row := db.QueryRow("select id, title, performer, vd_number, detail, maker, dmm_url, tag, series, created_at, updated_at from dmm where vd_number = ?", vd_number)
// 	//Scan對應的欄位與select語法的欄位順序一致
// 	if err := row.Scan(&video.ID, &video.Title, &video.Performer, &video.Vd_number, &video.Maker,
// 		&video.Detail, &video.Dmm_url, &video.Tag, &video.Series, &video.Created_at, &video.Updated_at); err != nil {
// 		fmt.Printf("scan failed, err:%v\n", err)

// 	}
// 	fmt.Printf("video:%+v\n", video)
// 	return c.JSON(video)
// }
// 撈取單筆資料
// func (h *Handler) HandleSearchTeamID(c *fiber.Ctx) error { // 撈取單筆資料

// 	db := h.db
// 	log.Printf("Successfully connected to database")

// 	team_id := c.Params("team_id")

// 	var member Member

// 	row := db.QueryRow("select name, ig_url, status from cpbl_member where team_id = ?", team_id)
// 	//Scan對應的欄位與select語法的欄位順序一致
// 	if err := row.Scan(&member.Name, &member.Ig_url, &member.Status); err != nil {
// 		fmt.Printf("scan failed, err:%v\n", err)

// 	}
// 	fmt.Printf("teamMember:%+v\n", member)
// 	return c.JSON(member)
// }
