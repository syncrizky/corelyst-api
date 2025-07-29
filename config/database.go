package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Gagal Memuat .env File!")
	}

	// Ambil variabel dari .env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Format DSN dengan timeout dan lokasi
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&timeout=5s",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	// Buka koneksi DB
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Gagal Membuka Koneksi DB: %v", err)
	}

	// Set pengaturan connection pool
	DB.SetMaxOpenConns(10)                 // maksimum koneksi aktif
	DB.SetMaxIdleConns(5)                  // maksimum koneksi idle
	DB.SetConnMaxLifetime(5 * time.Minute) // reset koneksi setiap 5 menit

	// Tes koneksi ke database
	err = DB.Ping()
	if err != nil {
		log.Fatalf("❌ Ping DB gagal: %v", err)
	}

	log.Println("✅ Koneksi ke database berhasil!")
}
