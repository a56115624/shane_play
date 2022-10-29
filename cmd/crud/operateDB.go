package crud

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// db_host=sgpdb.itlab.tw
// db_user=shane
// db_password=GKbCoMubLMQ6o
// db_name=shane
// db_charset=utf8mb4
// db_port=8889

// const (
// 	username = "root"
// 	password = "123456789"
// 	hostname = "127.0.0.1:3306"
// 	dbname   = "av_video"
// )

const (
	db_host     = "sgpdb.itlab.tw"
	db_user     = "shane"
	db_password = "GKbCoMubLMQ6o"
	db_name     = "shane"
	db_charset  = "utf8mb4"
	db_port     = "8889"
)

func dsn(db_name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", db_user, db_password, db_host+":"+db_port, db_name)
}

func DbConnection() (*sql.DB, error) { //建立db連線
	// db, err := sql.Open("mysql", dsn(""))
	db, err := sql.Open("mysql", dsn(db_name))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	//defer db.Close()

	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancelfunc()
	// res, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname) // 假如db不存在則創建
	// if err != nil {
	// 	log.Printf("Error %s when creating DB\n", err)
	// 	return nil, err
	// }
	// no, err := res.RowsAffected()
	// if err != nil {
	// 	log.Printf("Error %s when fetching rows", err)
	// 	return nil, err
	// }
	// log.Printf("rows affected %d\n", no)

	// db, err = sql.Open("mysql", dsn(dbname))
	// if err != nil {
	// 	log.Printf("Error %s when opening DB", err)
	// 	return nil, err
	// }
	//defer db.Close()

	// db.SetMaxOpenConns(20)                 // 設定最大連線數
	// db.SetMaxIdleConns(20)                 // 設定閒置連線量
	// db.SetConnMaxLifetime(time.Minute * 5) // 設置了連接可重用的最大時間

	// ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancelfunc()
	err = db.Ping()
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", db_name)
	return db, nil
}

// func CreateDataTable(db *sql.DB) error { // 創建table
// 	query := `CREATE TABLE IF NOT EXISTS gotable(id int primary key, site_id int, title text,
//          performer text, maker VARCHAR(255), detail text, vd_number text, tag text,
// 		 series text, dmm_url text, created_at VARCHAR(45), updated_at VARCHAR(45))`
// 	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancelfunc()
// 	res, err := db.Exec(query)
// 	if err != nil {
// 		log.Printf("Error %s when creating data table", err)
// 		return err
// 	}
// 	rows, err := res.RowsAffected()
// 	if err != nil {
// 		log.Printf("Error %s when getting rows affected", err)
// 		return err
// 	}
// 	log.Printf("Rows affected when creating table: %d", rows)
// 	return nil
// }

// func Insert(db *sql.DB, d Data) error { // 插入單筆資料
// 	query := `insert into gotable(id, site_id, title, performer, maker, detail, vd_number, tag, series,
// 		dmm_url, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
// 	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancelfunc()
// 	// stmt, err := db.PrepareContext(ctx, query)
// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return err
// 	}
// 	defer stmt.Close()
// 	res, err := stmt.Exec(d.ID, d.Site_id, d.Title, d.Performer, d.Maker, d.Detail, d.Vd_number,
// 		d.Series, d.Tag, d.Dmm_url, d.Created_at, d.Updated_at)
// 	if err != nil {
// 		log.Printf("Error %s when inserting row into datas table", err)
// 		return err
// 	}
// 	rows, err := res.RowsAffected() // 受影響的資料筆數
// 	if err != nil {
// 		log.Printf("Error %s when finding rows affected", err)
// 		return err
// 	}
// 	log.Printf("%d datas created ", rows) // 創建了幾筆資料
// 	return nil
// }

// func MultipleInsert(db *sql.DB, datas []Data) error { // 插入多筆資料
// 	query := `insert into gotable(id, site_id, title, performer, maker, detail, vd_number, tag, series,
// 		dmm_url, created_at, updated_at) values `
// 	var inserts []string
// 	var params []interface{}
// 	for _, v := range datas {
// 		inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)") // 因為不確定插入的資料筆數, 所以將每筆append inserts
// 		params = append(params, v.ID, v.Site_id, v.Title, v.Performer, v.Maker, v.Detail,
// 			v.Series, v.Tag, v.Vd_number, v.Dmm_url, v.Created_at, v.Updated_at)
// 	}
// 	// fmt.Println(inserts)  = [(?, ?, ?, ?, ?) (?, ?, ?, ?, ?)], 為了符合sql語法, 所以中間要用","串起來
// 	queryVals := strings.Join(inserts, ",") // 將複數筆資料串起來 (?, ?, ?, ?, ?), (?, ?, ?, ?, ?), (?, ?, ?, ?, ?)...
// 	query = query + queryVals               // 與 insert into 語句串起來
// 	log.Println("query is", query)
// 	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancelfunc()
// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return err
// 	}
// 	defer stmt.Close()
// 	res, err := stmt.Exec(params...)
// 	// log.Println(params...) params... = 2 Top of AVstar三上悠亜デビュー5周年プレミアムBEST 最新12タイトル86コーナー8時間スペシャル 三上悠亜 ofje00352 https://www.dmm.co.jp/litevideo/-/detail/=/cid=ofje00352/ 5 ドS愛人の超絶上から見下しマウント騎乗位！ダメな男を罵ることで快感を得る激ヤバ痴女の密室射精管理 本庄鈴 本庄鈴 1stars00450 https://www.dmm.co.jp/litevideo/-/detail/=/cid=1stars00450/
// 	// log.Println(params) params = [2 Top of AVstar三上悠亜デビュー5周年プレミアムBEST 最新12タイトル86コーナー8時間スペシャル 三上悠亜 ofje00352 https://www.dmm.co.jp/litevideo/-/detail/=/cid=ofje00352/ 5 ドS愛人の超絶上から見下しマウント騎乗位！ダメな男を罵ることで快感を得る激ヤバ痴女の密室射精管理 本庄鈴 本庄鈴 1stars00450 https://www.dmm.co.jp/litevideo/-/detail/=/cid=1stars00450/]
// 	if err != nil {
// 		log.Printf("Error %s when inserting row into products table", err)
// 		return err
// 	}
// 	rows, err := res.RowsAffected()
// 	if err != nil {
// 		log.Printf("Error %s when finding rows affected", err)
// 		return err
// 	}
// 	log.Printf("%d products created simulatneously", rows)
// 	return nil
// }

// func SelectID(db *sql.DB, DataID int) (Data, error) { // 選擇單行, 給予id, 搜索名稱

// 	log.Printf("Getting video Datas")

// 	query := `select id, title, performer, vd_number, dmm_url from dmm where id = ?`
// 	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancelfunc()
// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return Data{}, err
// 	}
// 	defer stmt.Close()

// 	var datas Data
// 	row := stmt.QueryRow(DataID)

// 	if err := row.Scan(&datas.ID, &datas.Title, &datas.Performer, &datas.Vd_number, &datas.Dmm_url); err != nil {
// 		fmt.Println(err)
// 		return Data{}, err
// 	}
// 	fmt.Println(datas)
// 	return datas, nil
// }

// func SelectDatasByID(db *sql.DB, idStart int, idEnd int) ([]Data, error) { // 選擇多行 + 多條件
// 	log.Printf("Getting id, title, performer by ID range")
// 	query := `select id, title, performer from dmm where id >= ? && id <= ?;` // 選取id範圍的title, performer...
// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return []Data{}, err
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query(idStart, idEnd)
// 	if err != nil {
// 		return []Data{}, err
// 	}
// 	defer rows.Close()
// 	var datas = []Data{}
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

// func DeleteData(db *sql.DB, DataID int) error { // 根據條件id 刪除資料
// 	query := "delete from gotable where id = ?"
// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return err
// 	}

// 	defer stmt.Close()

// 	res, err := stmt.Exec(DataID)
// 	if err != nil {
// 		log.Printf("Error %s when inserting row into datas table", err)
// 		return err
// 	}

// 	rows, err := res.RowsAffected() // 受影響的資料筆數
// 	if err != nil {
// 		log.Printf("Error %s when finding rows affected", err)
// 		return err
// 	}

// 	log.Printf("%d datas deleted ", rows) // 創建了幾筆資料
// 	return nil

// 	// prdID, err := res.LastInsertId() // 獲取最後插入的id
// 	// if err != nil {
// 	// 	log.Printf("Error %s when getting last inserted product", err)
// 	// 	return err
// 	// }
// 	// log.Printf("Data with ID %d created", prdID)
// 	// return nil
// }
