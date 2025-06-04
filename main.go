package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("failed to load env file: %w", err))
	}

	dbBtimApiHost := os.Getenv("DB_BTIM_API_HOST")
	dbBtimApiPort := os.Getenv("DB_BTIM_API_PORT")
	dbBtimApiUser := os.Getenv("DB_BTIM_API_USERNAME")
	dbBtimApiPassword := os.Getenv("DB_BTIM_API_PASSWORD")
	dbBtimApiDBName := os.Getenv("DB_BTIM_API_DATABASE")

	// Buka koneksi ke database
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", dbBtimApiUser, dbBtimApiPassword, dbBtimApiHost, dbBtimApiPort, dbBtimApiDBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke DB:", err)
	}

	// Model
	type Roles struct {
		RoleID      uint   `gorm:"primaryKey;column:role_id"`
		RoleName    string `gorm:"column:role_name"`
		Description string `gorm:"column:description"`
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	logTime := now.Format(time.RFC1123)
	uniqueCode := now.Format("02012006150405")

	// Insert data
	role := Roles{
		RoleName:    "Test Automation " + uniqueCode,
		Description: logTime,
	}

	result := db.Create(&role)
	if result.Error != nil {
		log.Fatal("Gagal insert:", result.Error)
	}

	fmt.Println("Berhasil insert role dengan ID:", role.RoleID)
}
