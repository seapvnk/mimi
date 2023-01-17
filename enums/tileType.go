package enums

type TileType int

const (
  Solid TileType = iota
  NonSolid
  Wall
  Tile
  NonSolidWall
  NonSolidTile
)
