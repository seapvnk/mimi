package repositories

import (
  "seapvnk/mimi/database"
  "seapvnk/mimi/models"
)

func GetMapsAll() []models.Map {
  db := database.Connection()

  var dmaps []models.Map
  db.Find(&dmaps)

  return dmaps
}

func GetMapById(id int) models.Map {
  db := database.Connection()

  var dmap models.Map
  db.First(&dmap, id)

  return dmap
}
