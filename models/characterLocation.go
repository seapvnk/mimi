package models

import "gorm.io/gorm"

type CharacterLocation struct {
  gorm.Model
  CharacterID   uint
  MapID         uint
  X             int
  Y             int
}
