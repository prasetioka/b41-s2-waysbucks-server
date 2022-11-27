package database

import (
  "waysbucks-api/models"
  "waysbucks-api/pkg/mysql"
  "fmt"
)

// Automatic Migration if Running App
func RunMigration() {
  err := mysql.DB.AutoMigrate(
    &models.User{},
    &models.Product{},
    &models.Topping{},
  )

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}