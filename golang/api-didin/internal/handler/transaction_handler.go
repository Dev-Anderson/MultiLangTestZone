package handler

import (
	"api-dindin/internal/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransactionHandler(ctx *gin.Context) {
	r := CreateTransactionRequest{}
	ctx.BindJSON(&r)
	if err := r.Validate(); err != nil {
		logger.Errorf("validation err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transaction := entity.Transaction{
		NameTransaction: r.NameTransaction,
		Value:           r.Value,
		Category:        r.Category,
	}

	if err := db.Create(&transaction).Error; err != nil {
		logger.Errorf("error creating transaction: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating transation on database")
		return
	}
	sendSuccess(ctx, "create-transation", transaction)
}

func GetTransactionHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	transaction := entity.Transaction{}
	if err := db.First(&transaction, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Transaction not found")
		return
	}
	sendSuccess(ctx, "show-transaction", transaction)
}

func EditTransactionHandler(ctx *gin.Context) {
	r := UpdateTransactionRequest{}

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

	transaction := entity.Transaction{}

	if err := db.First(&transaction, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Transaction not found")
		return
	}

	if r.NameTransation != "" {
		transaction.NameTransaction = r.NameTransation
	}

	if r.Value <= 0 {
		transaction.Value = r.Value
	}

	if r.Category != "" {
		transaction.Category = r.Category
	}

	if err := db.Save(&transaction).Error; err != nil {
		logger.Errorf("error updating transaction: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating transaction")
		return
	}

	sendSuccess(ctx, "updating-transaction", transaction)
}

func DeleteTransationHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	transation := entity.Transaction{}

	if err := db.First(&transation, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}

	if err := db.Delete(&transation).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting transation with id: %v", id))
		return
	}

	sendSuccess(ctx, "delete-transation", transation)
}

func GetAllTransactionHandler(ctx *gin.Context) {
	t := []entity.Transaction{}
	if err := db.Find(&t).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining transactions")
		return
	}

	sendSuccess(ctx, "list-transaction", t)
}
