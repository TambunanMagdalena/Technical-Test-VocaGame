package controllers

import (
	"net/http"
	"template-go/app/helpers"
	"template-go/app/models"

	"github.com/ezartsh/inrequest"
	"github.com/ezartsh/validet"
	"github.com/labstack/echo/v4"

	e "template-go/pkg/customerror"
)

type itemController controller

type ItemInterface interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Export(ctx echo.Context) error
}

func (c *itemController) Create(ctx echo.Context) error {
	var (
		reqBody models.ImportItem
		resBody *models.MasterItemResponse
		err     error
	)

	req, err := inrequest.Json(ctx.Request())
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	mapReq := req.ToMap()
	schema := validet.NewSchema(
		mapReq,
		map[string]validet.Rule{
			"code":     validet.String{Required: true},
			"name":     validet.String{Required: true},
			"category": validet.String{Required: true},
			//	"stock":    validet.String{Required: true},
		},
		validet.Options{},
	)

	errorBags, err := schema.Validate()
	if err != nil {
		err := e.NewBadRequestError(err.Error())
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), errorBags.Errors, nil, nil)
	}

	err = req.ToBind(&reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	// err = c.Options.UseCases.Validate.IsValidRequestItem(ctx.Request().Context(), mapReq, "create")
	// if err != nil {
	// 	return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	// }
	if reqBody.Stock < 0 {
		return helpers.StandardResponse(ctx, http.StatusBadRequest, []string{"stock harus >= 0"}, nil, nil)
	}
	resBody, err = c.Options.UseCases.Item.CreateMasterItem(ctx.Request().Context(), reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Success"}, resBody, nil)
}

func (c *itemController) Get(ctx echo.Context) error {
	filter := models.FilterRequest{
		Code:     ctx.QueryParam("code"),
		Name:     ctx.QueryParam("name"),
		Category: ctx.QueryParam("category"),
	}
	resBody, err := c.Options.UseCases.Item.GetMasterItemByFilter(ctx.Request().Context(), filter)
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Success"}, resBody, nil)
}

func (c *itemController) Update(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	name := ctx.QueryParam("name")
	if code == "" || name == "" {
		return helpers.StandardResponse(ctx, http.StatusBadRequest, []string{"code and name are required"}, nil, nil)
	}
	var reqBody models.ImportItem
	req, err := inrequest.Json(ctx.Request())
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}
	err = req.ToBind(&reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}
	if reqBody.Stock < 0 {
		return helpers.StandardResponse(ctx, http.StatusBadRequest, []string{"stock harus >= 0"}, nil, nil)
	}
	err = c.Options.UseCases.Item.UpdateMasterItem(ctx.Request().Context(), code, name, reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Success"}, nil, nil)
}

func (c *itemController) Delete(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	name := ctx.QueryParam("name")
	if code == "" || name == "" {
		return helpers.StandardResponse(ctx, http.StatusBadRequest, []string{"code and name are required"}, nil, nil)
	}
	err := c.Options.UseCases.Item.DeleteMasterItem(ctx.Request().Context(), code, name)
	if err != nil {
		return helpers.StandardResponse(ctx, e.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Success"}, nil, nil)
}

func (c *itemController) Export(ctx echo.Context) error {
	filter := models.FilterRequest{
		Code:     ctx.QueryParam("code"),
		Name:     ctx.QueryParam("name"),
		Category: ctx.QueryParam("category"),
	}
	data, err := c.Options.UseCases.Item.ExportData(ctx.Request().Context(), filter)
	if err != nil {
		return helpers.StandardResponse(ctx, http.StatusInternalServerError, []string{err.Error()}, nil, nil)
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Success"}, data, nil)
}
