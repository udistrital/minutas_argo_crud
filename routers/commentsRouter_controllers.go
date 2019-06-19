package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:ParametroMinutaTipoContratoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:PlantillaMinutaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/minutas_argo_crud/controllers:TipoContratoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
