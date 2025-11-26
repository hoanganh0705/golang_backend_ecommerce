package controller

import (
	"GolangBackendEcommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

// use this to make our routers acknowledge which controller we are calling
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	// if err != nil {
	// 	response.ErrorResponse(c, 20003, "Invalid parameter")
	// 	return
	// }
	// response.SuccessResponse(c, 20001, []string{"User1", "User2"})

}
