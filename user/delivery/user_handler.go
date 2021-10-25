package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/h4yfans/case-study/common"
	"github.com/h4yfans/case-study/domain"
	"github.com/h4yfans/case-study/models"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(usecase domain.UserUsecase, r *mux.Router) {
	handler := UserHandler{usecase: usecase}

	r.HandleFunc("/users", handler.Create).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", handler.Update).Methods(http.MethodPatch)
	r.HandleFunc("/users/{id}", handler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/users/{id}", handler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/users", handler.GetAllUser).Methods(http.MethodGet)
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})
		return
	}

	userData, err := u.usecase.Create(r.Context(), &user)
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})
		return
	}

	common.RespondWithJSON(w, http.StatusOK, userData)
	return
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})
		return
	}

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	user.ID = userID
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})

	}

	userData, err := u.usecase.Update(r.Context(), &user)
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})
		return
	}

	common.RespondWithJSON(w, http.StatusOK, userData)
	return
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})

	}

	err = u.usecase.Delete(r.Context(), userID)
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})
		return
	}

	common.RespondWithJSON(w, http.StatusNoContent, nil)
	return
}

func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})

	}

	user, err := u.usecase.GetByID(r.Context(), userID)
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})
		return
	}

	common.RespondWithJSON(w, http.StatusOK, user)
	return
}

func (u *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.usecase.GetAllUser(r.Context())
	if err != nil {
		common.RespondWithJSON(w, common.GetStatusCode(err), common.ResponseError{Error: err.Error()})
		return
	}

	common.RespondWithJSON(w, http.StatusOK, users)
	return
}
