package models

import (
  "gorm.io/gorm"
  "seapvnk/mimi/enums"
)

type Tile struct {
  gorm.Model
  MapID uint
  X     int
  Y     int
  Type  enums.TileType
}
