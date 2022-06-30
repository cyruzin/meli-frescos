package main

import (
	"database/sql"
	"log"

	"github.com/cyruzin/meli-frescos/internal/section/controller"
	"github.com/cyruzin/meli-frescos/internal/section/repository/mariadb"
	"github.com/cyruzin/meli-frescos/internal/section/service"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSource := "root:root@tcp(localhost:3306)/bootcamp?parseTime=true"

	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("failed to connect to mariadb")
	}

	router := gin.Default()

	repo := mariadb.NewMariaDBRepository(conn)
	srv := service.NewSection(repo)
	controller.NewSectionControler(router, srv)

	if err := router.Run(); err != nil {
		log.Fatal("failed to start the server. err:", err)
	}
}
