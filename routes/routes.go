package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoscorissa/gin-api-rest/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.POST("/alunos", controllers.CriarAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/busca", controllers.BuscaAluno)
	r.Run()
}
