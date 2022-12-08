package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// db, err := gorm.Open("mysql", "root:qwer1234@tcp(127.0.0.1:3306)/Guess_The_Logo?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open("root:qwer1234@tcp(127.0.0.1:3306)/Guess_The_Logo?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&Logo{})
	db.AutoMigrate(&Logo{})
	fmt.Println(db)
	records := readCsvFile("./LogoDatabase.csv")
	fmt.Println(records[1])
	var logoList []Logo
	for i, item := range records {
		fmt.Println(i, item)
		var logo Logo
		logo.Name = item[0]
		logo.ImagePath = item[1]
		logoList = append(logoList, logo)

	}
	if errList := db.Create(&logoList).Error; errList != nil {
		fmt.Println(errList)
	}

}

type Logo struct {
	gorm.Model
	ImagePath string `json:"Imagepath" gorm:"type:varchar(300)"`
	Name      string `json:"logo_name" gorm:"type:varchar(300)"`
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
