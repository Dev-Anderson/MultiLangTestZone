package handler

import (
	"api-dindin/internal/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBankHandler(ctx *gin.Context) {
	r := CreateBankRequest{}

	ctx.BindJSON(&r)

	if err := r.Validate(); err != nil {
		logger.Errorf("validation err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	bank := entity.Bank{
		NameBank:  r.NameBank,
		OwnerName: r.OwnerName,
	}

	if err := db.Create(&bank).Error; err != nil {
		logger.Errorf("error creating bank: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating bank on database")
		return
	}

	sendSuccess(ctx, "create-bank", bank)
}

func GetBankHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	bank := entity.Bank{}
	if err := db.First(&bank, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "bank not found")
		return
	}

	sendSuccess(ctx, "show-bank", bank)
}

func EditBankHandler(ctx *gin.Context) {
	r := UpdateBankRequest{}
	ctx.BindJSON(&r)
	if err := r.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	bank := entity.Bank{}
	if err := db.First(&bank, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "card not found")
		return
	}
	if r.NameBank != "" {
		bank.NameBank = r.NameBank
	}
	if r.OwnerName != "" {
		bank.OwnerName = r.OwnerName
	}

	if err := db.Save(&bank).Error; err != nil {
		logger.Errorf("error updating bank: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating bank")
		return
	}

	sendSuccess(ctx, "updating-bank", bank)
}

func DeleteBankHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	bank := entity.Bank{}
	if err := db.First(&bank, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("bank with id: %s not found", id))
		return
	}

	if err := db.Delete(&bank).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting bank with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-bank", bank)
}

func GetAllBankHandler(ctx *gin.Context) {
	banks := []entity.Bank{}
	if err := db.Find(&banks).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining banks")
		return
	}
	sendSuccess(ctx, "list-banks", banks)
}
