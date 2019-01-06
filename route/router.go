package route

import (
	"github.com/hmhr/go_sample/controllers"
	"github.com/labstack/echo"
)

/*
router
*/
func Init() *echo.Echo {
	e := echo.New()
	controllers.Init()

	api := e.Group("/sample_echo")

	{
		// 全件取得
		// curl -X GET http://localhost:1323/sample_echo/users
		api.GET("/users", controllers.FindUsers())
		// id検索
		// curl -X GET http://localhost:1323/sample_echo/users/1
		api.GET("/users/:id", controllers.GetUser())
		// 登録
		// curl -X POST http://localhost:1323/sample_echo/users -H 'Content-Type: application/json' -d '{"name":"taro", "user_id":"test_id"}'
		api.POST("/users", controllers.AddUser())
		// 更新
		//api.
	}

	return e
}
