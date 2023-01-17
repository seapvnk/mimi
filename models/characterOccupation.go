package models

import "gorm.io/gorm"

type CharacterOccupation struct {
  gorm.Model
  Name      string
  PlaceKey  string
}
