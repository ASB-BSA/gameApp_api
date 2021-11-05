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

	authen := v1.Use(middlewares.IsAuthenticated)
	authen.Get("user", controllers.GetUser)
	authen.Get("character", controllers.GetCharacter)
	authen.Get("room", controllers.GetRoom)
	authen.Post("room", controllers.PostRoom)

	battle := authen.Use(middlewares.IsBattle)
	battle.Get("battle", controllers.GetBattle)
}
