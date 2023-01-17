package models

import "gorm.io/gorm"

type Character struct {
  gorm.Model
  Name          string
  Money         float64
  Status        CharacterStatus
  Location      CharacterLocation
  Skills        []Skill `gorm:"many2many:character_skills;"`
  Occupations   []CharacterOccupation
}
