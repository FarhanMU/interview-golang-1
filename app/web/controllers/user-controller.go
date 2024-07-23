package controllers

import (
	"net/http"

	"zen-test/app/helpers"
	"zen-test/app/middleware"
	"zen-test/app/web"
	"zen-test/app/web/models"
	"zen-test/app/web/services"
)

type UserControllerImpl struct {
	UserService services.UserService
}

type UserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user account
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserCreate true "User Sign Up"
// @Success 200 {object} web.WebResponse{data=models.UserResponse}
// @Failure 400 {object} web.WebResponse
// @Router /users/signup [post]
func (c *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	userSignUp := models.UserCreate{}
	helpers.ToRequestBody(r, &userSignUp)

	userResponse := c.UserService.Register(r.Context(), userSignUp)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
	http.Redirect(w, r, "/api/login", http.StatusOK)
}

// Login godoc
// @Summary Log in a user
// @Description Authenticate a user and set a session cookie
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User Login"
// @Success 200 {object} web.WebResponse{data=models.UserResponse}
// @Failure 401 {object} web.WebResponse
// @Router /users/login [post]
func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := models.UserLogin{}
	helpers.ToRequestBody(r, &userLogin)

	userResponse, loggedIn := c.UserService.Login(r.Context(), userLogin)
	if loggedIn {
		helpers.SetUserCookie(w, r, userResponse.ID)
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   userResponse,
		}
		helpers.WriteResponseBody(w, webResponse)
	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helpers.WriteResponseBody(w, webResponse)
	}
}

// Update godoc
// @Summary Update user for the user
// @Description Update user for the authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserUpdate true "User update"
// @Param userId path string true "User ID"
// @Success 200 {object} web.WebResponse{data=models.UserResponse}
// @Failure 401 {object} web.WebResponse
// @Router /users/{userId} [put]
// @Security BearerAuth
func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	userUpdateRequest := models.UserUpdate{}
	helpers.ToRequestBody(r, &userUpdateRequest)

	userId := middleware.GetUserID(r)
	c.UserService.Update(r.Context(), userUpdateRequest, userId)
}

// Logout godoc
// @Summary Logout for the user
// @Description Logout for the authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /users/logout [get]
// @Security BearerAuth
func (c *UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	helpers.DeleteCookieHandler(w, r, "refresh_token")
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Logout success",
	}
	helpers.WriteResponseBody(w, response)
}
