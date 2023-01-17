package game

import (
  "time"

  _ "github.com/gen2brain/raylib-go/raylib"
  "seapvnk/mimi/models"
)

type Skill struct {
  CharacterRef    *Character
  Data            models.Skill
  invoke          func(*Character)
  lastTimeUsed    time.Time
}

func (skill *Skill) CanActivate() bool {
  inCooldown := true
  return inCooldown
}

func (skill *Skill) Activate() {
  if skill.CanActivate() {
    skill.invoke(skill.CharacterRef)
  }
}

func SimpleDash(character *Character) func() {
  return func() {
    skillFx := SkillEffect{
      Speed: float32(100),
      Direction: character.Aim,
    }

    character.SkillEffects = append(character.SkillEffects, skillFx)
  }
}
