package models

import "gorm.io/gorm"

type TrafficNode struct {
  gorm.Model
  MapID       uint
  X           int
  Y           int
}
