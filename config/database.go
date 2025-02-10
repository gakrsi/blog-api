// config/database.go
package config

import (
    "blog-api/models"
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        dsn = "host=localhost user=travlouser password=tz1rx2ac3vv4lb5on6 dbname=travlodb1 port=5432 sslmode=disable"
    }
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("Database connected successfully!")
    DB.AutoMigrate(&models.BlogPost{})

}