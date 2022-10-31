package handler

import (
	"context"
	"fmt"

	"github.com/teampui/pac"
	"github.com/uptrace/bun"
)

type RepoInterface interface {
	pac.Service
	getMysqlData() ([]MemberData, error)
	putMysqlData(books []*MemberData) error
	updateMysqlData(books []*MemberData) error
}
type PostgresRepo struct {
	db *bun.DB
}
type MemberData struct {
	bun.BaseModel `bun:"table:Golang_shane_data"`
	Id            int64 `bun:",pk,autoincrement"`
	Description   string
}

func (repo *PostgresRepo) Register(app *pac.App) {
	app.Repositories.Add("manager", repo)

	// try to get database connection from pac
	repo.db = pac.Must[*bun.DB](
		pac.Svc[*bun.DB](app, "db"),
		"repository/manager: cannot start due to no valid nsmh db found",
	)

	// 可在終端機看到執行的 sql query
	//repo.db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
}
func (repo *PostgresRepo) getMysqlData() ([]MemberData, error) {
	ctx := context.Background()

	var blocks []MemberData
	err := repo.db.NewSelect().Model(&blocks).Order("ID").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

// 插入資料進mysqlDb
func (repo *PostgresRepo) putMysqlData(books []*MemberData) error {
	ctx := context.Background()
	for i := 0; i < len(books); i++ {
		_, err := repo.db.NewInsert().Model(books[i]).Exec(ctx)
		if err != nil {
			return err
		}
		fmt.Println(books) // book id is scanned automatically
	}
	return nil
}

// 更新資料近mysql裡面
/*values := db.NewValues(&[]*Book{book1, book2})*/
func (repo *PostgresRepo) updateMysqlData(books []*MemberData) error {
	ctx := context.Background()
	for i := 0; i < len(books); i++ {
		_, err := repo.db.NewUpdate().Model(books[i]).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
