package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igorgabriel/api-workshop/src/models"
)

// InitializeRoutes ...
func InitializeRoutes(router *gin.Engine) {
	v0 := router.Group("/v0")
	{
		v0.GET("/ping", handlePing)
		v0.POST("/workshops", handlePostWorkshop)
		v0.PUT("/workshops/:id", handlePutWorkshop)
		v0.GET("/workshops/:id", handleGetWorkshop)
		v0.GET("/workshops", handleGetWorkshops)
		v0.DELETE("/workshops/:id", handleDeleteWorkshop)
	}
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func handleGetWorkshops(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	isTid := c.DefaultQuery("isTelegramId", "false")
	ti, err := strconv.ParseBool(isTid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if ti {
		cliente, err := controllers.GetClienteByTelegramID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		if cliente.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		id = cliente.ID
	}

	acS := c.DefaultQuery("active", "true")
	ac, err := strconv.ParseBool(acS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var ls []models.Lavoura
	if ac {
		ls, err = controllers.GetActiveLavouras(id)
	} else {
		ls, err = controllers.GetLavouras(id)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lavouras": ls})
}

func handlePostWorkshop(c *gin.Context) {
	var l models.Lavoura
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := controllers.SaveLavoura(l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lavoura criada com sucesso"})
}

func handlePutWorkshop(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var l models.Lavoura
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	l.ID = id
	if err := controllers.UpdateLavoura(l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lavoura atualizada com sucesso"})
}

func handleDeleteWorkshop(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err = controllers.DeleteLavoura(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lavoura removida com sucesso"})
}

func handleGetWorkshop(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	l, err := controllers.GetLavouraByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lavoura": l})
}
