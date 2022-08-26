package main

import (
	"fmt"
	"os"

	"github.com/ItsVasconcellos/oficina/controller"
	"github.com/ItsVasconcellos/oficina/database"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := database.Connect(); err != nil {
		panic(err)
	}
	db := database.GrabDB()
	st := db.MustBegin()
	sql := "CREATE TABLE IF NOT EXISTS usuarios (nome VARCHAR(128),ra VARCHAR(12));"
	if err := st.MustExec(sql); err != nil {
		panic(err)
	}
	r := gin.Default()

	r.POST("/Usuario", func(c *gin.Context) {
		if err := controller.PostAluno(c); err != nil {
			panic(err)
		}
	})

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
