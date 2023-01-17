package services

import (
  "fmt"

  "seapvnk/mimi/models"
  "seapvnk/mimi/repositories"
)

func LoadCharactersFromMap(mapId int) []models.Character {
  characters := repositories.GetCharactersAll()
  LoadCharacterAssets(characters)

  return characters
}

func LoadCharacterAssets(characters []models.Character) {
  for _, character := range characters {
      LoadCharacterAsset(character)
    }
}

func LoadCharacterAsset(character models.Character) {
  fmt.Println("character asset loaded for:", character.Name)
}
