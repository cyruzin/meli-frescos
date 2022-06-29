package main

import (
	"database/sql"
	"log"

	"github.com/cyruzin/meli-frescos/internal/section/controller"
	"github.com/cyruzin/meli-frescos/internal/section/repository/mariadb"
	"github.com/cyruzin/meli-frescos/internal/section/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := mariadb.NewMariaDBRepository(&sql.DB{})
	srv := service.NewSection(repo)
	sectionController, err := controller.NewSectionControler(srv)

	if err != nil {
		log.Fatal(err)
	}

	router.GET("/api/v1/sections", sectionController.GetAll())
	router.GET("/api/v1/sections/:id", sectionController.GetById())
	router.POST("/api/v1/sections", sectionController.Post())
	router.PATCH("/api/v1/sections/:id", sectionController.Patch())
	router.DELETE("/api/v1/sections/:id", sectionController.Delete())

	if err := router.Run(); err != nil {
		log.Fatal("failed to start the server. err:", err)
	}
}
