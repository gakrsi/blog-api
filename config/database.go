// config/database.go
package config

import (
    "blog-api/models"
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "postgresql://blogapiuser:SeW0D4uRbVRJDliDdzIsp3L3cZgwDqig@dpg-cukru3d2ng1s7380epjg-a.oregon-postgres.render.com/blogapi_ntkd"
    // if dsn == "" {
    //     dsn = "host=dpg-cukru3d2ng1s7380epjg-a.oregon-postgres.render.com user=blogapiuser password=SeW0D4uRbVRJDliDdzIsp3L3cZgwDqig dbname=blogapi_ntkd port=5432 sslmode=disable"
    //     postgresql://blogapiuser:SeW0D4uRbVRJDliDdzIsp3L3cZgwDqig@dpg-cukru3d2ng1s7380epjg-a.oregon-postgres.render.com/blogapi_ntkd
    // }
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("Database connected successfully!")
    DB.AutoMigrate(&models.BlogPost{})

}
