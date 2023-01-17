package database

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

func Connection() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{})
  if err != nil {
    panic(err.Error())
  }

  return db
}
