package gcode_parser

import (
	"context"
	"github.com/google/uuid"
	"github.com/santirufiner/slicerwise/internal/domain/model/gcode_parser"
)

type ParseMetadataCommand struct {
	UserId  uuid.UUID
	Content *string
}

type ParseMetadataHandler struct {
}

func NewParseMetadataHandler() *ParseMetadataHandler {
	return &ParseMetadataHandler{}
}

func (h *ParseMetadataHandler) Exec(ctx context.Context, cmd *ParseMetadataCommand) (*gcode_parser.GCodeMetadata, error) {
	// TODO: implement me
	return nil, nil
}
