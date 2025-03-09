package main

import (
	"security-bank-password/password"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	p := password.GeneratePermutions()

	r.GET("/pairs", func(c *gin.Context) {
		gp := password.GeneratePair(p)

		c.JSON(200, gin.H{
			"pairs":     gp,
			"validator": password.GenerateValidatorDigit(gp),
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var loginRequest struct {
			Pairs     [][]int `json:"pairs"`
			Validator int     `json:"validator"`
			Password  [][]int `json:"password"`
		}

		if err := ctx.BindJSON(&loginRequest); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		if len(loginRequest.Pairs) == 0 {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if loginRequest.Validator == 0 {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if loginRequest.Validator != password.GenerateValidatorDigit(loginRequest.Pairs) {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if !password.ValidatePasswordPairs(loginRequest.Pairs, loginRequest.Password) {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if loginRequest.Password == nil || len(loginRequest.Password) != 5 {
			ctx.JSON(400, gin.H{"error": "Password is required and must have 5 pairs"})
			return
		}

		ps := "12345"
		for i := range loginRequest.Password {
			pair := loginRequest.Password[i]

			if int(ps[i]-'0') != pair[0] && int(ps[i]-'0') != pair[1] {
				ctx.JSON(401, gin.H{"error": "Invalid password"})
				return
			}
		}

		ctx.JSON(200, gin.H{"message": "Welcome!"})
	})

	r.Run()
}
