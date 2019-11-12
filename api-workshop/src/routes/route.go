package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.sandmanbb.com/pivotal/agrows-client/src/controllers"
	"gitlab.sandmanbb.com/pivotal/agrows-client/src/models"
)

// InitializeRoutes ...
func InitializeRoutes(router *gin.Engine) {
	v0 := router.Group("/v0")
	{
		v0.GET("/ping", handlePing)
		v0.POST("/signin", handlePostSignin)
		v0.POST("/produtores", handlePostProdutor)
	}
	{
		auth := v0.Group("/")
		auth.Use(authenticationRequired())
		{
			auth.GET("/forecast", handleGetForecast)
			auth.GET("/produtores/:id/areas", handleGetAreas)
			auth.POST("/produtores/:id/areas", handlePostArea)
			auth.PUT("/produtores/:id/areas/:areaId", handlePutArea)
			auth.DELETE("/produtores/:id/areas/:areaId", handleDeleteArea)
			auth.POST("/produtores/:id/cultivares", handlePostCultivar)
			auth.GET("/produtores/:id/cultivares", handleGetCultivares)
			auth.GET("/produtores/:id/cultivares/:cultivarId", handleGetCultivar)
			auth.PUT("/produtores/:id/cultivares/:cultivarId", handlePutCultivar)
			auth.DELETE("/produtores/:id/cultivares/:cultivarId", handleDeleteCultivar)
			auth.POST("/produtores/:id/lavouras", handlePostLavoura)
			auth.PUT("/produtores/:id/lavouras/:lavouraId", handlePutLavoura)
			auth.GET("/produtores/:id/lavouras/:lavouraId", handleGetLavoura)
			auth.GET("/produtores/:id/lavouras", handleGetLavouras)
			auth.DELETE("/produtores/:id/lavouras/:lavouraId", handleDeleteLavoura)
			auth.POST("/produtores/:id/lavouras/:lavouraId/irrigacoes", handlePostIrrigacao)
			auth.PUT("/produtores/:id/irrigacoes/:irrigacaoId", handlePutIrrigacao)
			auth.GET("/produtores/:id/lavouras/:lavouraId/irrigacoes", handleGetIrrigacoes)
			auth.POST("/produtores/:id", handlePutProdutor)
			auth.POST("/messages/remind", handlePostReminders)
			auth.POST("/messages/suggest", handlePostSuggestions)
			auth.POST("/messages/summary", handlePostIrrigationsSummary)
			auth.POST("/messages/queries", handlePostIrrigationQueries)
			auth.GET("/produtores/:id/dashboard/:lavouraId", handleGetDashboard)
		}
	}
}

func authenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.GetHeader("Authorization")
		logrus.Debugf(a)
		if a == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}

		s := strings.Split(a, " ")
		token := s[1]

		err := controllers.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func handleGetForecast(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "forecast",
	})
}

func handleGetAreas(c *gin.Context) {
	a, err := controllers.GetAreas()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"areas": a,
	})
}

func handlePostArea(c *gin.Context) {
	p := models.Area{Pid: 1}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		logrus.Errorf("Error binding JSON: %s", err)
		c.String(http.StatusBadRequest, "Error binding request body JSON. Err:", err)
		return
	}

	newArea, err := controllers.SaveArea(p)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "area criado com sucesso",
		"content": newArea,
	})
}

func handlePutArea(c *gin.Context) {

	areaID, err := strconv.Atoi(c.Param("areaId"))

	if err != nil {
		logrus.Errorf("Error binding parameters to handler: %s", err)
		c.String(http.StatusBadRequest, "Error binding request. Err:", err)
		return
	}

	a := models.Area{Pid: 1}
	err = c.ShouldBindJSON(&a)
	if err != nil {
		logrus.Errorf("Error binding JSON: %s", err)
		c.String(http.StatusBadRequest, "Error binding request body JSON. Err:", err)
		return
	}
	a.ID = areaID
	err = controllers.UpdateArea(a)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "area alterado com sucesso",
	})
}

func handleDeleteArea(c *gin.Context) {
	areaID, err := strconv.Atoi(c.Param("areaId"))
	if err != nil {
		logrus.Errorf("Error binding parameters to handler: %s", err)
		c.String(http.StatusBadRequest, "Error binding request. Err:", err)
		return
	}

	err = controllers.DeleteArea(areaID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "area deletado com sucesso",
	})
}

func handlePostCultivar(c *gin.Context) {
	var cu models.Cultivar
	if err := c.ShouldBindJSON(&cu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := controllers.SaveCultivar(cu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cultivar criado com sucesso"})
}

func handleGetCultivares(c *gin.Context) {
	cs, err := controllers.GetCultivares()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cultivares": cs})
}

func handleGetCultivar(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	cu, err := controllers.GetCultivarByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cultivar": cu})
}

func handlePutCultivar(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var cu models.Cultivar
	if err := c.ShouldBindJSON(&cu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	cu.ID = id
	if err := controllers.UpdateCultivar(cu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cultivar atualizado com sucesso"})
}

func handleDeleteCultivar(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err = controllers.DeleteCultivar(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cultivar removido com sucesso"})
}

func handleGetLavouras(c *gin.Context) {
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

func handlePostLavoura(c *gin.Context) {
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

func handlePutLavoura(c *gin.Context) {
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

func handleDeleteLavoura(c *gin.Context) {
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

func handleGetLavoura(c *gin.Context) {
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

func handlePostIrrigacao(c *gin.Context) {
	idS := c.Param("id")
	lID, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var irr models.Irrigacao
	if err := c.ShouldBindJSON(&irr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	irr.LavouraID = lID
	iID, err := controllers.SaveIrrigacao(irr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Irrigação criada com sucesso", "irrigacaoId": iID})
}

func handlePutIrrigacao(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var i models.Irrigacao
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	i.ID = id
	if err := controllers.UpdateIrrigacao(i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Irrigação atualizada com sucesso"})
}

func handleGetIrrigacoes(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	is, err := controllers.GetIrrigacoes(id, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"irrigacoes": is})
}

func handlePostProdutor(c *gin.Context) {
	p := models.Produtor{}

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = controllers.SaveCliente(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "produtor criado com sucesso",
	})
}

func handlePutProdutor(c *gin.Context) {
	p := models.Produtor{}
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = controllers.UpdateCliente(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "produtor atualizado com sucesso",
	})
}

func handlePostIrrigationQueries(c *gin.Context) {

	clientes, err := controllers.GetAllClientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, cliente := range clientes {
		logrus.Debugf("Sending irrigation summary for client %v", cliente)
		irrigations, err := controllers.GetIrrigationSummary(cliente.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = controllers.CallBotIrrigationQueries(cliente, irrigations)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Irrigations summary send to the user succesfully",
		"clientes": clientes,
	})
}

func handlePostIrrigationsSummary(c *gin.Context) {

	clientes, err := controllers.GetAllClientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, cliente := range clientes {
		logrus.Debugf("Sending irrigation summary for client %v", cliente)
		irrigations, err := controllers.GetIrrigationSummary(cliente.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = controllers.CallBotIrrigationsSummary(cliente, irrigations)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Irrigations summary send to the user succesfully",
		"clientes": clientes,
	})
}

func handlePostSuggestions(c *gin.Context) {

	var sr models.UserMessageRequest
	err := c.ShouldBindJSON(&sr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	logrus.Debug(sr)
	irrigations, err := controllers.GetRainAmount(sr.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = controllers.CallBotSuggestions(sr.UserID, irrigations)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Suggestions send to the user succesfully",
		"userID":      sr.UserID,
		"suggestions": irrigations,
	})
}

func handlePostReminders(c *gin.Context) {

	var sr models.UserMessageRequest

	err := c.ShouldBindJSON(&sr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = controllers.CallBotReminder(sr.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Remainders sent to the user succesfully",
		"userID":  sr.UserID,
	})
	return
}

func handleGetDashboard(c *gin.Context) {
	lID, err := strconv.Atoi(c.Param("lavouraId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	dashboard, err := controllers.GetDashboardInfo(lID)
	c.JSON(http.StatusOK, dashboard)
	return
}

func handlePostSignin(c *gin.Context) {
	var cred models.Credentials
	err := c.ShouldBindJSON(&cred)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	usr, siErr := controllers.SignIn(cred)
	if siErr != nil {
		c.JSON(siErr.Code, gin.H{"message": siErr.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
	return
}
