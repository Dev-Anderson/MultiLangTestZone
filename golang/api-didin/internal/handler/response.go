package handler

import (
	"api-dindin/internal/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/type")

	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Context-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateUserResponse struct {
	Message string              `json:"message"`
	Data    entity.UserResponse `json:"data"`
}

type DeleteUserResponse struct {
	Message string              `json:"message"`
	Data    entity.UserResponse `json:"data"`
}

type ShowUserResponse struct {
	Message string              `json:"message"`
	Data    entity.UserResponse `json:"data"`
}

type GetAllUserResponse struct {
	Message string                `json:"message"`
	Data    []entity.UserResponse `json:"data"`
}

type UpdateUserResponse struct {
	Message string              `json:"message"`
	Data    entity.UserResponse `json:"data"`
}

type CreateTransactionResponse struct {
	Message string                     `json:"message"`
	Data    entity.TransactionResponse `json:"data"`
}

type DeleteTransactionResponse struct {
	Message string                     `json:"message"`
	Data    entity.TransactionResponse `json:"data"`
}

type ShowTransactionResponse struct {
	Message string                     `json:"message"`
	Data    entity.TransactionResponse `json:"data"`
}

type GetAllTransactionResponse struct {
	Message string                       `json:"message"`
	Data    []entity.TransactionResponse `json:"data"`
}

type CreateBankResponse struct {
	Message string              `json:"message"`
	Data    entity.BankResponse `json:"data"`
}

type DeleteBankResponse struct {
	Message string              `json:"message"`
	Data    entity.BankResponse `json:"data"`
}

type ShowBankResponse struct {
	Message string              `json:"message"`
	Data    entity.BankResponse `json:"data"`
}

type GetAllBankResponse struct {
	Message string                `json:"message"`
	Data    []entity.BankResponse `json:"data"`
}

type CreateCardResponse struct {
	Message string              `json:"message"`
	Data    entity.CardResponse `json:"data"`
}

type DeleteCardResponse struct {
	Message string              `json:"message"`
	Data    entity.CardResponse `json:"data"`
}

type ShowCardResponse struct {
	Message string              `json:"message"`
	Data    entity.CardResponse `json:"data"`
}

type GetAllCardResponse struct {
	Message string                `json:"message"`
	Data    []entity.CardResponse `json:"data"`
}
