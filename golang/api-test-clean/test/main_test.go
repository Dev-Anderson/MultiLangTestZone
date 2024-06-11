package main_test

import (
	"api-test/cmd/handler"
	"api-test/internal/database"
	"api-test/internal/repository"
	"api-test/internal/services"
	"api-test/pkg/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os/user"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	router := gin.New()

	dbConn := database.InitDB()
	defer dbConn.Close()
	userService := &services.UserService{UserRepository: &repository.UserRepository{DB: dbConn}}
	userHandler := &handler.UserHandler{
		UserService: userService,
	}

	router.GET("/user/:id", userHandler.GetUserByID)

	req, err := http.NewRequest("GET", "user/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var gotUser models.User
	err = json.Unmarshal(w.Body.Bytes(), &gotUser)
	assert.NoError(t, err)
}

type mockUserService struct{}

func (s *mockUserService) GetUserByID(id int) (*user.User, error) {
	// Aqui você pode retornar um usuário falso para testar seu handler
	if id == 1 {
		return &models.User{ID: 1, Name: "Test User", Email: "teste@example.com", Password: "123"}, nil
	}
	return nil, nil
}

func TestUserHandler_GetUserByID(t *testing.T) {
	dbConn := database.InitDB()
	defer dbConn.Close()
	userService := &services.UserService{UserRepository: &repository.UserRepository{DB: dbConn}}
	userHandler := &handler.UserHandler{
		UserService: userService,
	}

	handler := &userHandler{
		UserService: &mockUserService{},
	}

	// Crie um contexto gin de teste
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Defina o parâmetro de ID no contexto de teste
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Chame o método GetUserByID do handler com o contexto de teste
	handler.GetUserByID(c)

	// Verifique se o status da resposta é 200 OK
	assert.Equal(t, http.StatusOK, c.Writer.Status())

	// Verifique se o corpo da resposta contém o usuário esperado
	var gotUser user.User
	err := json.Unmarshal(c.Writer.Body.Bytes(), &gotUser)
	assert.NoError(t, err)

	wantUser := &user.User{ID: 1, Name: "Test User"}
	assert.True(t, reflect.DeepEqual(wantUser, &gotUser))
}
