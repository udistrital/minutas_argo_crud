// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/minutas_argo_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/parametro_minuta",
			beego.NSInclude(
				&controllers.ParametroMinutaController{},
			),
		),

		beego.NSNamespace("/parametro_minuta_tipo_contrato",
			beego.NSInclude(
				&controllers.ParametroMinutaTipoContratoController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/plantilla_minuta",
			beego.NSInclude(
				&controllers.PlantillaMinutaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
