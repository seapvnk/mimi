package models

import "gorm.io/gorm"

type District struct {
  gorm.Model
  Name        string
  Maps        []Map
}
