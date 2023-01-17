package models

import (
  "gorm.io/gorm"
  "seapvnk/mimi/enums"
)

type Item struct {
  gorm.Model
  Key   string
  Name  string
  Type  enums.ItemType
}
