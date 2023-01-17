package game

import "github.com/gen2brain/raylib-go/raylib"

type SkillEffect struct {
  Attack    int
  Defense   int
  Duration  int
  Speed     float32
  Direction rl.Vector2
}
