package controller

import (
	"fmt"
	conv "github.com/capkeik/backend-trainee-assignment-2023/internal/convert/csv"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/repository/pg"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/repository/static"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/request"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type RecordsController struct {
	recordsService *pg.RecordsRepo
	csvRepo        *static.CSVRepo
}

func NewRecords(recordsService *pg.RecordsRepo, csvRepo *static.CSVRepo) RecordsController {
	return RecordsController{recordsService: recordsService, csvRepo: csvRepo}
}

func (c RecordsController) Get(ctx echo.Context) error {
	var req request.RecordsReq

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	records, err := c.recordsService.GetRecords(req.ID, req.From, req.To)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Internal Error"})
	}

	if len(*records) == 0 {
		errStr := fmt.Sprintf("No records for user with id <%v>", req.ID)
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": errStr})
	}

	var data = &[][]string{}
	for _, r := range *records {
		rec, err := conv.Record(*r)
		*data = append(*data, rec)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "csv converting error"})
		}
	}
	filename := fmt.Sprintf("%v-%v.csv", req.ID, time.Now().Unix())
	err = c.csvRepo.SaveCSV(filename, *data)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Error saving csv"})
	}
	downPath := fmt.Sprintf("/records/download/csv/%v", filename)
	return ctx.JSON(http.StatusOK, map[string]interface{}{"link": downPath})
}

func (c RecordsController) Download(ctx echo.Context) error {
	filename := ctx.Param("filename")
	log.Println(filename)
	fileInfo, filePath, err := c.csvRepo.GetCSV(filename)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "File not found"})
	}

	ctx.Response().Header().Set("Content-Disposition", "attachment; filename="+filename)
	ctx.Response().Header().Set("Content-Type", "text/csv")
	ctx.Response().Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	return ctx.File(filePath)
}
