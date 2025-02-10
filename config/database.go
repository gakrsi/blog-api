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
        dsn = "host=dpg-cukru3d2ng1s7380epjg-a.oregon-postgres.render.com user=blogapiuser password=SeW0D4uRbVRJDliDdzIsp3L3cZgwDqig dbname=blogapi_ntkd port=5432 sslmode=disable"
    }
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("Database connected successfully!")
    DB.AutoMigrate(&models.BlogPost{})

}
