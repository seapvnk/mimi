package models

import (
  "gorm.io/gorm"
  "seapvnk/mimi/enums"
)

type Quest struct {
  gorm.Model
  Key         string
  Title       string
  Description string
  Active      bool
  Type        enums.QuestType
  Scheduled   bool
}
