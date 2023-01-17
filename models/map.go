package models

import "gorm.io/gorm"

type Map struct {
  gorm.Model
  Name          string
  DistrictID    uint
  Tiles         []Tile
  Objects       []MapObject
  TrafficNodes  []TrafficNode
}
