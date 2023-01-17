package models

import "gorm.io/gorm"

type CharacterStatus struct {
  gorm.Model
  CharacterID   uint
  Hp            int
  Mana          int
  Stamina       int
  Speed         int
  Defense       int
  Attack        int
}
