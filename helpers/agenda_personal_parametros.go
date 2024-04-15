package helpers

import (
	"fmt"

	"github.com/udistrital/agenda_personal_mid_parametros/models"
)

func ListarAgendaPersonalParametros() (agendaPersonalParametros models.AgendaPersonalParametros, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "ListarAgendaPersonalParametros", "err": err, "status": "500"}
			panic(outputError)
		}
	}()

	var agendaPersonal []models.Contacto
	var parametros []models.TipoParametro

	urlAgendaPersonal := "contacto?&limit=0"
	if err := GetRequestNew("UrlCrudAgendaPersonal", urlAgendaPersonal, &agendaPersonal); err != nil {
		panic(err.Error())
	}
	fmt.Println("AGENDA_PERSONAL ", agendaPersonal)

	url := "tipo_parametro?query=nombre:Cargos&limit=0"
	if err := GetRequestNew("UrlCrudParametros", url, &parametros); err != nil {
		panic(err.Error())
	}
	fmt.Println("PARAMETROS ", parametros)

	agendaPersonalParametros.AgendaPersonal = agendaPersonal
	agendaPersonalParametros.Parametros = parametros

	fmt.Println(agendaPersonalParametros)

	return agendaPersonalParametros, outputError
}
