package enums

type QuestType int

const (
  Main QuestType = iota
  Secondary
  Event
  Daily
  Monthly
)
