package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/form_common/app/model"
	"github.com/zngue/form_common/app/router"
	"github.com/zngue/go_helper/pkg/common_run"
	"github.com/zngue/go_helper/pkg/response"
	"gorm.io/gorm"
)

func main() {
	run()

}

func run() {
	common_run.CommonGinRun(
		common_run.FnRouter(func(engine *gin.Engine) {
			engine.NoRoute(func(c *gin.Context) {
				response.HttpFailWithCodeAndMessage(404, "路由不存在", c)
			})
			groups := engine.Group("form")
			fmt.Println(groups)
			router.Router(groups)
		}),
		common_run.IsRegisterCenter(true),
		common_run.MysqlConn(func(db *gorm.DB) {
			//db.AutoMigrate(new(model.ZngUser), new(model.ZngKm), new(model.ZngOrder))
			db.AutoMigrate(model.NewCategory())
		}),
	)

}
