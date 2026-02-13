package users

import "github.com/gin-gonic/gin"

type UserDeliveryInterface interface {
	RegisterUser() gin.HandlerFunc
	LoginUser() gin.HandlerFunc
	// GetList() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	// GetById() gin.HandlerFunc
	Update() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	// 	GenerateOTP() gin.HandlerFunc
	// 	VerifyOTP() gin.HandlerFunc
	// 	ResetPassword() gin.HandlerFunc
	// 	ResendOTP() gin.HandlerFunc
}
