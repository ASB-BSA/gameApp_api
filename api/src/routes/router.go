package routes

import (
	"boomin_game_api/src/controllers"
	"boomin_game_api/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	v1 := api.Group("v1")

	v1.Post("user", controllers.PostUser)

	admin := v1.Group("admin")
	admin.Post("login", controllers.AdminLogin)

	isadmin := admin.Use(middlewares.IsAdmin)
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
	authen.Get("user", controllers.GetUser)
	authen.Get("character", controllers.GetCharacter)
	authen.Get("room", controllers.GetRoom)
	authen.Post("room", controllers.PostRoom)

	battle := authen.Use(middlewares.IsBattle)
	battle.Get("battle", controllers.GetBattle)
}
