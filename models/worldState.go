package models

import "gorm.io/gorm"

type WorldState struct {
  gorm.Model
  Day     uint8
  Month   uint8
  Hour    uint8
  Minute  uint8
}
