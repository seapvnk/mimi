package repositories

import (
  "seapvnk/mimi/database"
  "seapvnk/mimi/models"
)

func GetCharactersAll() []models.Character {
  db := database.Connection()

  var characters []models.Character
  db.Find(&characters)

  return characters
}

func GetCharacterById(id int) models.Character {
  db := database.Connection()

  var character models.Character
  db.First(&character, id)

  return character
}
