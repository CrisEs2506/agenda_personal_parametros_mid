package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/agenda_personal_mid_parametros/helpers"
)

type AgendaPersonalParametrosController struct {
	beego.Controller
}

func (c *AgendaPersonalParametrosController) URLMapping() {

}

// GetAll ...
// @Title GetAll
// @Description get Obtener agenda_personal y parametros
// @Success 200 {object} []models.AgendaPersonalParametros
// @Failure 400 bad request
// @Failure 500 Internal server error
// @router / [get]
func (c *AgendaPersonalParametrosController) GetAll() {
	defer helpers.ErrorController(c.Controller, "AgendaPersonalParametrosController")

	if v, err := helpers.ListarAgendaPersonalParametros(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Listado consultado con éxito", "Data": v}
	} else {
		panic(err)
	}

	c.ServeJSON()
}
