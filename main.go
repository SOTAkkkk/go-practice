package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Company struct {
	CompanyCd    int    `json:"company_cd"`
	RrCd         int    `json:"rr_cd"`
	Name         string `json:"name"`
	CompanyNameK string `json:"company_name_k"`
	CompanyNameH string `json:"company_name_h"`
	CompanyNameR string `json:"company_name_r"`
	CompanyUrl   string `json:"company_url"`
	CompanyType  int    `json:"company_type"`
	EStatus      int    `json:"e_status"`
	ESort        int    `json:"e_sort"`
}

func main() {
	// Capture connection properties.
	//cfg := mysql.Config{
	//	//User:                 os.Getenv("DBUSER"),
	//	//Passwd:               os.Getenv("DBPASS"),
	//	//Net:                  "tcp",
	//	//Addr:                 "127.0.0.1:3306",
	//	DBName:               "eki",
	//	AllowNativePasswords: true,
	//}
	// Get a database handle.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}
	//db, err = sql.Open("mysql", cfg.FormatDSN())
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// Ping the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database!")

	// DBからデータ取得
	company, err := CompanyByName("JR北海道")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Company found: %v\n", company)

	// JSON化
	companyJson, err := json.Marshal(company)
	if err != nil {
		log.Fatal(err)
	}

	// リクエストハンドラーの設定
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", companyJson)
		log.Println("access '/'")
	})

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func CompanyByName(name string) ([]Company, error) {
	var companies []Company

	// SQLクエリの実行
	rows, err := db.Query("SELECT * FROM companies WHERE company_name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("failed to query company by name %q: %v", name, err)
	}
	defer rows.Close()

	// レコードの処理
	for rows.Next() {
		var comp Company
		if err := rows.Scan(
			&comp.CompanyCd,
			&comp.RrCd,
			&comp.Name,
			&comp.CompanyNameK,
			&comp.CompanyNameH,
			&comp.CompanyNameR,
			&comp.CompanyUrl,
			&comp.CompanyType,
			&comp.EStatus,
			&comp.ESort,
		); err != nil {
			return nil, fmt.Errorf("failed to scan company by name %q: %v", name, err)
		}
		companies = append(companies, comp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over company rows by name %q: %v", name, err)
	}
	return companies, nil
}
