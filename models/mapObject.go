package models

import "gorm.io/gorm"

type MapObject struct {
  gorm.Model
  MapID   uint
  X       int
  Y       int
}
