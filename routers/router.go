package routers

import (
	"github.com/komputeko/terminigo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index_:slang.php", &controllers.SearchWord{})
    beego.Router("/:slang/", &controllers.MainController{})
    beego.Router("/:slang/:word/:wlang", &controllers.ShowPage{})
    beego.Router("/s/:slang", &controllers.SearchWord{})
}
