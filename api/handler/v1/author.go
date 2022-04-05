package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/perfectogo/api-gateway/api/handler/models"
	"github.com/perfectogo/api-gateway/genproto/catalog"
	"github.com/perfectogo/api-gateway/pkg/logger"
	"github.com/perfectogo/api-gateway/pkg/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

//  CreateAuthor ...
// @Summary CreateAuthor
// @Router /v1/authors/ [post]
// @Description This API for creating a new author
// @Tags author
// @Accept  json
// @Produce  json
// @Param Author request body models.Author true "authorCreateRequest"
// @Success 200 {object} models.AuthorResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel

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

// GetAuthors ...
// @Router /v1/authors [get]
// @Summary GetAuthors
// @Description This API for getting list of authors
// @Tags author
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListAuthors
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) GetAuthors(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.CatalogService().GetAuthors(
		ctxt, &catalog.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", logger.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetAuthor
// @Router /v1/authors/{id} [get]
// @Summary GetAuthor
// @Description This API for getting author Author by id
// @Tags author
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.AuthorResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) GetAuthor(ctx *gin.Context) {
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

// UpdateAuthor ...
// @Router /v1/authors/{id} [put]
// @Summary UpdateAuthor
// @Description This API for updating author
// @Tags author
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Author request body models.UpdateAuthor true "authorUpdateRequest"
// @Success 200 {object} models.AuthorResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) UpdateAuthor(ctx *gin.Context) {
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
	body.AuthorId = ctx.Param("id")

	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().UpdateAuthor(ctxt, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(403, err)
		h.log.Error("failed to update author", logger.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// DeleteAuthor ...
// @Router /v1/authors/{id} [delete]
// @Summary DeleteAuthor
// @Description This API for deleting author
// @Tags author
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DeleteAuthor(ctx *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := ctx.Param("id")

	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().DeleteAuthor(ctxt, &catalog.ByIdReq{Id: guid})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete auhor", logger.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
