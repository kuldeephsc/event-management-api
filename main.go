package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kuldeephsc/api/db"
	"github.com/kuldeephsc/api/routes"
)

func main() {
	fmt.Println("Kuldeep")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8085")
}
