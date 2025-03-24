package gcode_parser

import "github.com/google/uuid"

type GCodeMetadata struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	MachineName  string
	Flavor       string
	TimeSeconds  int
	FilamentUsed float64
	LayerHeight  float64
	BoundingBox  BoundingBox
	Settings     map[string]any
}

type BoundingBox struct {
	MinX, MinY, MinZ float64
	MaxX, MaxY, MaxZ float64
}
