package route

import (
	"github.com/labstack/echo"
	"github.com/yushiga/go_echo/controllers"
)

/*
router
*/
func Init() *echo.Echo {
	e := echo.New()
	controllers.Init()

	api := e.Group("/sample_echo")

	{
		// サンプルデータ登録
		// curl -X POST http://localhost:1323/sample_echo/sample -H 'Content-Type: application/json'
		api.POST("/sample", controllers.CreateSampleData())
		// 全件取得
		// curl -X GET http://localhost:1323/sample_echo/user
		api.GET("/user", controllers.FindUsers())
		// id検索
		// curl -X GET http://localhost:1323/sample_echo/user/1
		api.GET("/user/:id", controllers.GetUser())
		// 登録
		// curl -X POST http://localhost:1323/sample_echo/user -H 'Content-Type: application/json' -d '{"name":"taro", "user_id":"test_id"}'
		api.POST("/user", controllers.AddUser())
		// 更新
		// curl -X PUT http://localhost:1323/sample_echo/user -H 'Content-Type: application/json' -d '{"id":1, "name":"aaaaa"}'
		api.PUT("/user", controllers.UpdateUser())
		// 削除
		//api.DELETE("/user/:id", controllers.DeleteUser())
	}

	return e
}
