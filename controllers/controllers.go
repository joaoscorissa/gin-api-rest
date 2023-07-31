package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaoscorissa/gin-api-rest/database"
	"github.com/joaoscorissa/gin-api-rest/models"
)

func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "aluno não encontrado"})
	} else {
		c.JSON(http.StatusOK, aluno)
	}
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAluno(c *gin.Context) {
	var aluno models.Aluno
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao processar o JSON"})
		return
	}

	tipo, hasTipo := jsonData["tipo"].(string)
	valor, hasName := jsonData["valor"].(string)

	if !hasTipo || !hasName {
		c.JSON(400, gin.H{"error": "Campos inválidos"})
		return
	}
	switch tipo {
	case "cpf":
		database.DB.Where(&models.Aluno{Cpf: valor}).Find(&aluno)
	case "rg":
		database.DB.Where(&models.Aluno{Rg: valor}).Find(&aluno)
	}

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
