package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

import "github.com/go-sql-driver/mysql"

var db *sql.DB

type Company struct {
	ID           int
	RrCd         int
	Name         string
	CompanyNameK string
	CompanyNameH string
	CompanyNameR string
	CompanyUrl   string
	CompanyType  int
	EStatus      int
	ESort        int
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "eki",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	company, err := CompanyByName("JR北海道")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Company found: %v\n", company)
}

func CompanyByName(name string) ([]Company, error) {
	var company []Company

	rows, err := db.Query("SELECT * FROM companies WHERE company_name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("companyByName %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var comp Company
		if err := rows.Scan(
			&comp.ID,
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
			return nil, fmt.Errorf("companyByName %q: %v", name, err)
		}
		company = append(company, comp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("companyByName %q: %v", name, err)
	}
	return company, nil
}
