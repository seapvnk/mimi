package game

import (
  "github.com/gen2brain/raylib-go/raylib"
  "seapvnk/mimi/models"
)

type Game struct {
  WorldState      models.WorldState
  CurrentMap      models.Map

  ScreenWidth     uint8
  ScreenHeight    uint8
  FramesCounter   uint8
  Offset          rl.Vector2
}
