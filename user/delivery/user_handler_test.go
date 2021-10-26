package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gorilla/mux"
	"github.com/h4yfans/case-study/common"
	"github.com/h4yfans/case-study/domain"
	mocks "github.com/h4yfans/case-study/mocks/domain"
	"github.com/h4yfans/case-study/models"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		// Mock body
		userBody := models.User{
			Name:     "Kaan",
			Email:    "kaan@test.com",
			Password: "123123",
		}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/users", strings.NewReader(string(r)))
		assert.NoError(t, err)

		// Mock response
		var userResponse *domain.UserResponse
		err = faker.FakeData(&userResponse)
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Create", req.Context(), &userBody).Return(userResponse, nil)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.Create(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)

	})

	t.Run("should return 400", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/users", strings.NewReader(string(r)))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Create", req.Context(), &userBody).Return(nil, common.BadRequest)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.Create(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 500", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/users", strings.NewReader(string(r)))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Create", req.Context(), &userBody).Return(nil, common.ServerError)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.Create(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 403", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/users", strings.NewReader(string(r)))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Create", req.Context(), &userBody).Return(nil, common.UserAlreadyExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.Create(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 404", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/users", strings.NewReader(string(r)))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Create", req.Context(), &userBody).Return(nil, common.UserNotExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.Create(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		// Mock body
		userBody := models.User{
			Name:  "Kaan",
			Email: "kaan@test.com",
		}

		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(string(r)))
		assert.NoError(t, err)

		// Mock response
		userResponse := &domain.UserResponse{
			ID:    1,
			Name:  "Kaan",
			Email: "kaan@test.com",
		}

		userBody.ID = 1
		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Update", context.Background(), &userBody).Return(userResponse, nil)

		rec := httptest.NewRecorder()
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)
		handler := UserHandler{usecase: mockUCase}

		handler.Update(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)

	})

	t.Run("should return 400", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(string(r)))
		assert.NoError(t, err)

		userBody.ID = 1
		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Update", context.Background(), &userBody).Return(nil, common.BadRequest)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Update(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 500", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(string(r)))
		assert.NoError(t, err)

		userBody.ID = 1
		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Update", req.Context(), &userBody).Return(nil, common.ServerError)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Update(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 403", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(string(r)))
		assert.NoError(t, err)

		userBody.ID = 1
		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Update", req.Context(), &userBody).Return(nil, common.UserAlreadyExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Update(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 404", func(t *testing.T) {
		// Mock body
		userBody := models.User{}
		r, err := json.Marshal(&userBody)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/users/1", strings.NewReader(string(r)))
		assert.NoError(t, err)

		userBody.ID = 1
		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Update", req.Context(), &userBody).Return(nil, common.UserNotExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Update(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return 204", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Delete", context.Background(), 1).Return(nil)

		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)
		handler := UserHandler{usecase: mockUCase}

		rec := httptest.NewRecorder()
		handler.Delete(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)
		mockUCase.AssertExpectations(t)

	})

	t.Run("should return 400", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Delete", context.Background(), 1).Return(common.BadRequest)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Delete(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 500", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Delete", req.Context(), 1).Return(common.ServerError)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Delete(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 403", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Delete", req.Context(), 1).Return(common.UserAlreadyExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Delete(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("Delete", req.Context(), 1).Return(common.UserNotExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.Delete(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		userResponse := &domain.UserResponse{}

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetByID", context.Background(), 1).Return(userResponse, nil)

		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)
		handler := UserHandler{usecase: mockUCase}

		rec := httptest.NewRecorder()
		handler.GetByID(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)

	})

	t.Run("should return 400", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetByID", context.Background(), 1).Return(nil, common.BadRequest)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.GetByID(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 500", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetByID", req.Context(), 1).Return(nil, common.ServerError)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.GetByID(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 403", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetByID", req.Context(), 1).Return(nil, common.UserAlreadyExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.GetByID(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetByID", req.Context(), 1).Return(nil, common.UserNotExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		handler.GetByID(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}

func TestGetAllUser(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		assert.NoError(t, err)

		var userResponse []domain.UserResponse

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetAllUser", context.Background()).Return(userResponse, nil)

		handler := UserHandler{usecase: mockUCase}

		rec := httptest.NewRecorder()
		handler.GetAllUser(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)

	})

	t.Run("should return 400", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetAllUser", context.Background()).Return(nil, common.BadRequest)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.GetAllUser(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 500", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetAllUser", req.Context()).Return(nil, common.ServerError)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.GetAllUser(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 403", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetAllUser", req.Context()).Return(nil, common.UserAlreadyExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.GetAllUser(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("should return 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		assert.NoError(t, err)

		mockUCase := new(mocks.UserUsecase)
		mockUCase.On("GetAllUser", req.Context()).Return(nil, common.UserNotExist)

		rec := httptest.NewRecorder()
		handler := UserHandler{usecase: mockUCase}

		handler.GetAllUser(rec, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}
