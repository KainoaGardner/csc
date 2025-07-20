package api

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/auth"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/user"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func (h *Handler) registerUserRoutes(r chi.Router) {
	r.Get("/user/{userID}", h.getUser)

	r.Get("/user", h.getAuthUser)
	r.Post("/user", h.createUser)
	r.Delete("/user", h.deleteUser)

	r.Post("/user/login", h.loginUser)
	r.Post("/user/logout", h.logoutUser)

	r.Get("/user/all", h.getAllUsers)
	r.Delete("/user/all", h.deleteAllUsers)

	r.Post("/auth/refresh", h.refreshToken)

	r.Post("/user/password/forgot", h.forgotPassword)
	r.Post("/user/password/reset", h.resetPassword)

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

	err = user.CheckUniqueLogin(h.client, h.config.DB, *newUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = db.CreateUser(h.client, h.config.DB, newUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userStats := user.SetupUserStats(*newUser)
	_, err = db.CreateUserStats(h.client, h.config.DB, userStats)

	data := types.UserResponse{
		ID:          newUser.ID,
		Username:    newUser.Username,
		Email:       newUser.Email,
		CreatedTime: newUser.CreatedTime,
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("User created"), data)

}

// admin
func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.config.DB, h.config.JWT.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	userID := chi.URLParam(r, "userID")

	dbUser, err := db.FindUser(h.client, h.config.DB, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "User", dbUser)
}

// auth
func (h *Handler) getAuthUser(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.config.DB, h.config.JWT.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	dbUser, err := db.FindUser(h.client, h.config.DB, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := types.UserResponse{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Email:       dbUser.Email,
		CreatedTime: dbUser.CreatedTime,
	}

	utils.WriteResponse(w, http.StatusOK, "User", data)
}

// auth
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.config.DB, h.config.JWT.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	count, err := db.DeleteUser(h.client, h.config.DB, claims.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Deleted count", count)
}

// admin
func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.config.DB, h.config.JWT.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	users, err := db.ListAllUsers(h.client, h.config.DB)
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

// admin
func (h *Handler) deleteAllUsers(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.config.DB, h.config.JWT.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	amount, err := db.DeleteAllUsers(h.client, h.config.DB)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"count": amount}

	utils.WriteResponse(w, http.StatusOK, "Users deleted", data)

}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	var postLogin types.PostLogin
	err := utils.ParseJSON(r, &postLogin)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	dbUser, err := db.FindUserFromUsername(h.client, h.config.DB, postLogin.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	correctLogin := auth.CheckPasswordHash(postLogin.Password, dbUser.PasswordHash)
	if !correctLogin {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Incorrect login information"))
		return
	}

	accessExpireTime := 24 * 14 * time.Hour
	accessToken, err := auth.CreateToken(h.config.JWT.AccessKey, dbUser.ID.Hex(), accessExpireTime)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	refreshExpireTime := 24 * 14 * time.Hour
	refreshToken, err := auth.CreateToken(h.config.JWT.RefreshKey, dbUser.ID.Hex(), refreshExpireTime)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/auth/refresh",
		HttpOnly: true,
		Secure:   false, //change for prod
		SameSite: http.SameSiteStrictMode,
		MaxAge:   24 * 14 * 60 * 60,
	})

	data := types.TokenResponse{
		AccessToken: accessToken,
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Logged in"), data)
}

func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/auth/refresh",
		HttpOnly: true,
		Secure:   false, //change for prod
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Logged out"))
}

// auth
func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	refreshToken := cookie.Value

	claims, err := auth.ParseToken(h.config.JWT.RefreshKey, refreshToken)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	expired := auth.CheckExpiredToken(claims)
	if expired {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	accessExpireTime := 24 * 14 * time.Hour
	accessToken, err := auth.CreateToken(h.config.JWT.AccessKey, claims.UserID, accessExpireTime)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	refreshExpireTime := 24 * 14 * time.Hour
	refreshToken, err = auth.CreateToken(h.config.JWT.RefreshKey, claims.UserID, refreshExpireTime)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/auth/refresh",
		HttpOnly: true,
		Secure:   false, //change for prod
		SameSite: http.SameSiteStrictMode,
		MaxAge:   24 * 14 * 60 * 60,
	})

	data := types.TokenResponse{
		AccessToken: accessToken,
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Updated access token"), data)
}

func (h *Handler) forgotPassword(w http.ResponseWriter, r *http.Request) {
	var postForgot types.PostForgotPassword
	err := utils.ParseJSON(r, &postForgot)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	dbUser, err := db.FindUserFromEmail(h.client, h.config.DB, postForgot.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resetExpireTime := 15 * time.Minute
	resetToken, err := auth.CreateToken(h.config.JWT.PasswordRefreshKey, dbUser.ID.Hex(), resetExpireTime)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.SendResetPasswordEmail(h.config, postForgot.Email, resetToken)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"email": postForgot.Email}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Email sent"), data)
}

func (h *Handler) resetPassword(w http.ResponseWriter, r *http.Request) {
	claims, statusCode, err := auth.CheckValidAuth(h.client, h.config.DB, h.config.JWT.PasswordRefreshKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	var postReset types.PostResetPassword
	err = utils.ParseJSON(r, &postReset)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	passwordHash, err := auth.HashPassword(postReset.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = db.UpdateUserPassword(h.client, h.config.DB, claims.ID, passwordHash)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"_id": claims.ID}
	utils.WriteResponse(w, http.StatusOK, "Password updated", data)
}
