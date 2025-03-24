package gcode_parser

import "context"

type Repository interface {
	Create(ctx context.Context, data *GCodeMetadata) (*GCodeMetadata, error)
}
