package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal_mid_parametros/controllers:AgendaPersonalParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal_mid_parametros/controllers:AgendaPersonalParametrosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
