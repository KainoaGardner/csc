package api

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/user"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) registerUserRoutes(r chi.Router) {
	r.Post("/user/", h.createUser)
	r.Get("/user/{userID}", h.getUser)
	r.Delete("/user/{userID}", h.deleteUser)

	// r.Post("/user/login", h.loginUser)
	// r.Post("/user/logout", h.loginUser)

	r.Get("/user", h.getAllUsers)
	r.Delete("/user", h.deleteAllUsers)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var postUser types.PostUser
	err := utils.ParseJSON(r, &postUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := user.SetupNewUser(postUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = user.CheckUniqueLogin(h.client, h.dbConfig, *newUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = db.CreateUser(h.client, h.dbConfig, newUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userStats := user.SetupUserStats(*newUser)
	_, err = db.CreateUserStats(h.client, h.dbConfig, userStats)

	data := types.UserResponse{
		ID:          newUser.ID,
		Username:    newUser.Username,
		Email:       newUser.Email,
		CreatedTime: newUser.CreatedTime,
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("User created"), data)

}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.ListAllUsers(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	result := []types.UserResponse{}
	for _, dbUser := range users {
		userResponse := types.UserResponse{
			ID:          dbUser.ID,
			Username:    dbUser.Username,
			Email:       dbUser.Email,
			CreatedTime: dbUser.CreatedTime,
		}

		result = append(result, userResponse)
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("%d users found", len(result)), result)
}

func (h *Handler) deleteAllUsers(w http.ResponseWriter, r *http.Request) {
	amount, err := db.DeleteAllUsers(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"count": amount}

	utils.WriteResponse(w, http.StatusOK, "Users deleted", data)

}
