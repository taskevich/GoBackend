package v1

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"main/internal/dto"
	"main/internal/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUserById
// @Summary      Get User by ID
// @Description  Retrieve a user by their unique ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string  "Invalid user ID"
// @Failure      500  {object}  map[string]string  "Internal Server Error"
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers
// @Summary      Get All Users
// @Description  Retrieve a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Failure      500  {object}  map[string]string  "Internal Server Error"
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UsersResponse{Users: users})
}

// ExportUsersCsv
// @Summary      Get All Users In Csv
// @Description  Retrieve a Csv file with all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Failure      500  {object}  map[string]string  "Internal Server Error"
// @Router       /users/export [get]
func (h *UserHandler) ExportUsersCsv(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", `attachment; filename="users.csv"`)

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	headers := []string{"Id", "Login", "Password", "Role", "CreatedAt", "UpdatedAt"}
	if err := writer.Write(headers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		record := []string{
			strconv.Itoa(int(user.Id)),
			user.Login,
			user.Password,
			user.Role.Name,
			user.CreatedAt.String(),
			user.UpdatedAt.String(),
		}
		if err := writer.Write(record); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

// DeleteUserById
// @Summary      Delete User by ID
// @Description  Delete a user by their unique ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]string  "OK"
// @Failure      400  {object}  map[string]string  "Invalid user ID"
// @Failure      500  {object}  map[string]string  "Internal Server Error"
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	err = h.userService.DeleteUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// AddUser @Summary      Add a New User
// @Description  Add a new user to the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserRequest  true  "New User Data"
// @Success      200   {object}  map[string]string      "User ID"
// @Failure      400   {object}  map[string]string      "Bad Request"
// @Failure      500   {object}  map[string]string      "Internal Server Error"
// @Router       /users [post]
func (h *UserHandler) AddUser(c *gin.Context) {
	var newUser dto.CreateUserRequest
	if err := c.ShouldBindBodyWithJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := h.userService.AddUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": userId})
}
