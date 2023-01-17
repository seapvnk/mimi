package game

import (
  "github.com/gen2brain/raylib-go/raylib"
  "seapvnk/mimi/models"
)

type Character struct {
  Data          models.Character

  Mana                      int
  Speed                     float32
  Position, Direction, Aim  rl.Vector2
  Hitbox                    rl.Rectangle
  Skills                    []Skill
  SkillEffects              []SkillEffect
}
