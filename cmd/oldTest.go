package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

/*這個code只能更改資料庫的值跟索引,無法開啟api*/

// 記得資料庫的欄位都要大寫
type GolangShaneData struct {
	// 實際在使用的資料庫
	// bun.BaseModel `bun:"table:jpmnb_img"`
	// ImgId         int64 `bun:",pk"`
	// AlbumId       int
	// ImgUrl        string
	// ImgPage       int
	// Status        int

	// testData
	bun.BaseModel `bun:"table:Golang_shane_data"`
	Id            int64 `bun:",pk"`
	Description   string
}

const (
	mysqlDsn = "shane:GKbCoMubLMQ6o@tcp(sgpdb.itlab.tw:8889)/shane?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	mysqlDb *bun.DB
)

// 連線進資料庫mysqlDb
func init() {
	sqldb, err := sql.Open("mysql", mysqlDsn)
	if err != nil {
		panic(err)
	}

	mysqlDb = bun.NewDB(sqldb, mysqldialect.New())
	// 下面這行是要顯示出sql的語法
	mysqlDb.AddQueryHook(bundebug.NewQueryHook())

}

// 拿到mysql裡面的資料
func getMysqlData() ([]GolangShaneData, error) {
	ctx := context.Background()

	var blocks []GolangShaneData
	err := mysqlDb.NewSelect().Model(&blocks).Order("ID").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

// 插入資料進mysql裡面
func putMysqlData() (GolangShaneData, error) {
	book := GolangShaneData{Id: 3, Description: "我好想睡覺"}
	ctx := context.Background()
	_, err := mysqlDb.NewInsert().Model(&book).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(book) // book id is scanned automatically
	return book, nil
}

// 更新資料近mysql裡面
/*values := db.NewValues(&[]*Book{book1, book2})*/
func updateMysqlData(books []*GolangShaneData) ([]*GolangShaneData, error) {
	ctx := context.Background()

	for i := 0; i < len(books); i++ {
		_, err := mysqlDb.NewUpdate().Model(books[i]).WherePK().Exec(ctx)
		if err != nil {
			panic(err)
		}
	}

	return books, nil
}

// 執行
func main() {

	putString := []string{"我假日還要工作", "我假日想放假放假", "我假日不想開會"}
	book1 := GolangShaneData{Id: 1, Description: putString[0]}
	book2 := GolangShaneData{Id: 2, Description: putString[1]}
	book3 := GolangShaneData{Id: 3, Description: putString[2]}
	books := []*GolangShaneData{&book1, &book2, &book3}
	_, err := updateMysqlData(books)
	if err != nil {
		panic(err)
	}
	blocks, err := getMysqlData()

	if err != nil {
		panic(err)
	}

	fmt.Println(blocks)
}
