package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/santirufiner/slicerwise/internal/domain/actions/gcode_parser"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ParseMetadata(handler *gcode_parser.ParseMetadataHandler, l *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var uri parseMetadataUri
		if err := c.ShouldBindUri(&uri); err != nil {
			l.Debugf("Invalid URI param: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing user_id"})
			return
		}

		var req parseMetadataReq
		if err := c.ShouldBindJSON(&req); err != nil {
			l.Debugf("Invalid JSON: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
			return
		}

		metadata, err := handler.Exec(c.Request.Context(), &gcode_parser.ParseMetadataCommand{
			UserId:  uri.UserId,
			Content: req.Content,
		})
		if err != nil {
			l.Errorf("Handler failed: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, metadata)
	}
}

type parseMetadataUri struct {
	UserId uuid.UUID `uri:"user_id" binding:"required"`
}

type parseMetadataReq struct {
	Content *string `json:"content" binding:"required"`
}
