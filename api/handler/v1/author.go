package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/perfectogo/api-gateway/genproto/catalog"
	"github.com/perfectogo/api-gateway/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *handlerV1) PostAuthor(ctx *gin.Context) {
	var (
		body        catalog.Author
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}

	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().CreateAuthor(ctxt, &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create author", logger.Error(err))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *handlerV1) GetAuthors(ctx *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := ctx.Param("id")
	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().GetAuthor(
		ctxt, &catalog.ByIdReq{
			Id: guid,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get author", logger.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *handlerV1) GetAuthor(ctx *gin.Context) {

}

func (h *handlerV1) UpdateAuthor(ctx *gin.Context) {

}

func (h *handlerV1) DeleteAuthor(ctx *gin.Context) {

}
