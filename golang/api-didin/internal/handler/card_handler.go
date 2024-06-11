package handler

import (
	"api-dindin/internal/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCardHandler(ctx *gin.Context) {
	r := CreateCardRequest{}
	ctx.BindJSON(&r)
	if err := r.Validate(); err != nil {
		logger.Errorf("validation err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	card := entity.Card{
		NameCard:   r.NameCard,
		NameOwner:  r.NameOwner,
		NumberCard: r.NumberCard,
		CvvCard:    r.CvvCard,
	}

	if err := db.Create(&card).Error; err != nil {
		logger.Errorf("error creating card: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating card on database")
		return
	}

	sendSuccess(ctx, "create-card", card)
}

func GetCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	card := entity.Card{}

	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "card not found")
		return
	}

	sendSuccess(ctx, "show-card", card)
}

func EditCardHandler(ctx *gin.Context) {
	request := UpdateCardRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	card := entity.Card{}

	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "card not found")
		return
	}

	if request.NameCard != "" {
		card.NameCard = request.NameCard
	}

	if request.NameOwner != "" {
		card.NameOwner = request.NameOwner
	}

	if request.NumberCard <= 0 {
		card.NumberCard = request.NumberCard
	}

	if request.CvvCard <= 0 {
		card.CvvCard = request.CvvCard
	}

	if err := db.Save(&card).Error; err != nil {
		logger.Errorf("error updating card: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating card")
		return
	}

	sendSuccess(ctx, "updating-card", card)
}

func DeleteCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	card := entity.Card{}

	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("card with id: %s not found", id))
		return
	}

	if err := db.Delete(&card).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting card with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-card", card)
}

func GetAllCardHandler(ctx *gin.Context) {
	cards := []entity.Card{}

	if err := db.Find(&cards).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining cards")
		return
	}

	sendSuccess(ctx, "list-cards", cards)
}
