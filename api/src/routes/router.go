package routes

import (
	"boomin_game_api/src/controllers"
	"boomin_game_api/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	v1 := api.Group("v1")

	v1.Get("image/:id", controllers.Images)
	v1.Post("user", controllers.PostUser)

	admin := v1.Group("admin")
	admin.Post("login", controllers.AdminLogin)

	isadmin := admin.Use(middlewares.IsAdmin)
	isadmin.Get("character", controllers.GetCharacter)
	isadmin.Put("character/:id", controllers.PutCharacter)
	isadmin.Get("admins", controllers.Admins)
	isadmin.Post("register", controllers.AdminRegister)
	isadmin.Get("characteristic", controllers.GetCharacteristic)
	isadmin.Post("characteristic", controllers.CreateCharacteristic)
	isadmin.Post("settings", controllers.CreateSettingGroup)
	isadmin.Get("settings", controllers.GetSettingGroup)
	isadmin.Put("settings", controllers.PutSettingGroup)
	isadmin.Delete("settings", controllers.DeleteSettingGroup)
	isadmin.Post("settings/:id", controllers.CreateSettingItem)
	isadmin.Put("settings/:id", controllers.PutSettingItem)
	isadmin.Delete("settings/:id", controllers.DeleteSettingItem)

	authen := v1.Use(middlewares.IsAuthenticated)

	authen.Get("settings/:id", controllers.ExportSetting)
	authen.Get("user", controllers.GetUser)
	authen.Get("characteristic", controllers.GetCharacteristic)
	authen.Get("character", controllers.GetCharacter)
	authen.Put("team/:id", controllers.PutTeamCharacter)
	authen.Get("room", controllers.GetRoom)
	authen.Post("room", controllers.PostRoom)
	authen.Put("room", controllers.PostRoom)
	authen.Delete("room/:id", controllers.DeleteRoom)
	authen.Post("battle", controllers.PostBattle)

	battle := authen.Use(middlewares.IsBattle)
	battle.Get("battle", controllers.GetBattle)
	battle.Get("fight", controllers.CreateBattleLog)
}
