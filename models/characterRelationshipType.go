package models

import (
  "gorm.io/gorm"
  "seapvnk/mimi/enums"
)


type CharacterRelationshipType struct {
  gorm.Model
  Type enums.CharacterRelationshipType
}
