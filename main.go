package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "ncode/docs"
	"ncode/ncode"
	"fmt"
	cors "github.com/rs/cors/wrapper/gin"
)

var router *gin.Engine

// @title NationalCode
// @version 0.1
// @description This API validates Iran's NationalCode

// @contact.name Saeede Moghimi
// @contact.url saeedehmoghimi.com
// @contact.email moghimi.saeedeh@gmail.com

// @host 127.0.0.1:8080
// @BasePath /validation/


func main() {

  router = gin.Default()
  router.Use(cors.Default())
  url := ginSwagger.URL("127.0.0.1:8080/swagger/doc.json")
  router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler,url))

  
  router.GET("/validation/", validateNationalCode)


  router.Run()

}


// validateNationalCode godoc
// @Summary Validates Iranian's NationalCode
// @Description get NationalCode by id
// @Tags NationalCode
// @Accept  text/plain
// @Produce  json
// @Param  id query string true "NationalCode ID"
// @Router / [get]
// @Success 200 {object} ncode.NationalCode
func validateNationalCode(c *gin.Context){
	code := c.Query("id")

	fmt.Println(code)
	nationalCode := ncode.NewNationalCode(code)
	c.JSON(200,nationalCode)

}