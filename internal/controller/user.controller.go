package controller

import (
	"GolangBackendEcommerce/internal/service"
	"GolangBackendEcommerce/internal/vo"
	"GolangBackendEcommerce/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

// type UserController struct {
// 	userService *service.UserService
// }

// // use this to make our routers acknowledge which controller we are calling
// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	// if err != nil {
// 	// 	response.ErrorResponse(c, 20003, "Invalid parameter")
// 	// 	return
// 	// }
// 	// response.SuccessResponse(c, 20001, []string{"User1", "User2"})
// 	result := uc.userService.GetUserByEmail("email", "purpose")
// }

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamValid, err.Error())
		return
	}
	fmt.Printf("Registration params: %+v\n", params)
	fmt.Printf("Email: %s, Purpose: %s\n", params.Email, params.Purpose)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
