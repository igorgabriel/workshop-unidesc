package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igorgabriel/api-workshop/src/controllers"
	"github.com/igorgabriel/api-workshop/src/models"
)

// InitializeRoutes ...
func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", handlePing)
		v1.POST("/workshops", handlePostWorkshop)
		v1.PUT("/workshops/:id", handlePutWorkshop)
		v1.GET("/workshops/:id", handleGetWorkshop)
		v1.GET("/workshops", handleGetWorkshops)
		v1.DELETE("/workshops/:id", handleDeleteWorkshop)
	}
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func handleGetWorkshops(c *gin.Context) {
	var ws []models.Workshop

	ws, err := controllers.GetWorkshops()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"workshops": ws})
}

func handlePostWorkshop(c *gin.Context) {
	var w models.Workshop
	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := controllers.SaveWorkshop(w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Workshop criado com sucesso"})
}

func handlePutWorkshop(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var w models.Workshop
	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	w.ID = id
	if err := controllers.UpdateWorkshop(w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Workshop atualizado com sucesso"})
}

func handleDeleteWorkshop(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err = controllers.DeleteWorkshop(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Workshop removido com sucesso"})
}

func handleGetWorkshop(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	w, err := controllers.GetWorkshopByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"workshop": w})
}
