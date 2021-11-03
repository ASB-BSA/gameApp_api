package main

import (
	"boomin_game_api/src/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	err = godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		panic("Error getting .env data!")
	}

	DBName := os.Getenv("DB")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	DB, err = gorm.Open(mysql.Open(DBUser+":"+DBPass+"@tcp(db:3306)/"+DBName+"?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func main() {
	Connect()
	raw, err := ioutil.ReadFile("../../assets/character.json")
	if err != nil {
		panic(err)
	}

	var character []models.Characters
	json.Unmarshal(raw, &character)

	for _, ft := range character {
		fmt.Println(ft.Name)
	}

	DB.Create(&character)
}
