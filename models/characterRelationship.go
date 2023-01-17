package models

import "gorm.io/gorm"

type CharacterRelationship struct {
  gorm.Model
  CharacterFromID uint
  CharacterTo uint
  RelationshipTypes []CharacterRelationshipType
}
