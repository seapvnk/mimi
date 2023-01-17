package models

import "gorm.io/gorm"

type QuestReward struct {
  gorm.Model
  QuestID   uint
  Title     string
  Message   string
  Money     float64
  Items     []Item  `gorm:"many2many:quest_reward_items;"`
}
