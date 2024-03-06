package controller

import (
	"BDServer/internal/models"
	"BDServer/internal/service"
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	UserService service.UserService
}

func New(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

// @Summary Create
// @Description create user
// @Tags Users
// @Accept json
// @Prodece json
// @Param input body models.UserRequest true "User payload"
// @Success 201 {object} map[string]interface{} "User"
// @Router /users [post]
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := models.User{
		Username:  userRequest.Username,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := c.UserService.CreateUser(context.Background(), &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Get a user by ID
// @Description Retrieve a user by their unique ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Router /users/{id} [get]
func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	user, err := c.UserService.GetUserByID(context.Background(), userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Update a user
// @Description Update a user's details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Router /users [put]
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var updateUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.UserService.UpdateUser(context.Background(), &updateUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Delete a user
// @Description Delete a user by their unique ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Router /users/{id} [delete]
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	if err := c.UserService.DeleteUser(context.Background(), userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary List users
// @Description Retrieve a list of users
// @Tags Users
// @Accept json
// @Produce json
// @Param input body models.Conditions true "User conditions"
// @Router /users [get]
func (c *UserController) List(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	condition := models.Conditions{
		Limit:  limit,
		Offset: offset,
	}

	users, err := c.UserService.ListUsers(context.Background(), condition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
