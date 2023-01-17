package models

import (
  "gorm.io/gorm"
  "seapvnk/mimi/enums"
)

type Skill struct {
  gorm.Model
  Method    string
  Modifier  float64
  Cooldown  int
  Type      enums.SkillType
}
