package main

import (
	"github.com/cozkul/umai/server/handlers"
	"github.com/cozkul/umai/server/middleware"
	"github.com/gofiber/fiber/v2" // Fiber is an Express.js styled HTTP web framework running on Fasthttp
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/auth/signup", handlers.SignUpUser)
	app.Post("/api/auth/signin", handlers.SignInUser)
	app.Post("/api/auth/logout", middleware.DeserializeUser, handlers.LogoutUser)

	app.Get("/api/user/getSystems", middleware.DeserializeUser, handlers.GetUserSystems)
	app.Get("/api/user/getPlanets", middleware.DeserializeUser, handlers.GetUserPlanets)

	app.Get("/api/systemTime", middleware.DeserializeUser, handlers.GetSystemTime)

	app.Get("/api/universe/getSystem/:system", middleware.DeserializeUser, handlers.GetUniverseSystem)

	app.Get("/api/planet/resources/:planet", middleware.DeserializeUser, handlers.GetPlanetResources)
	app.Get("/api/planet/buildings/:planet", middleware.DeserializeUser, handlers.GetPlanetBuildings)

	// app.Get("api/celestialBody/getResources/:celestialBody", middleware.DeserializeUser, handlers.GetResources)
	// app.Get("/api/celestialBody/getBuildings/:celestialBody", middleware.DeserializeUser, handlers.Buildings)
	// app.Get("api/celestialBody/getPopulation/:celestialBody", middleware.DeserializeUser, handlers.GetPopulation)
	// app.Get("api/celestialBody/getCities/:celestialBody", middleware.DeserializeUser, handlers.GetCities)

	// Route::get('city/getResources/{city}', 'Game\CityController@getResources');
	// Route::get('city/getPopulation/{city}', 'Game\CityController@getPopulation');
	// Route::get('city/getActionPoint/{city}', 'Game\CityController@getActionPoint');
	// Route::get('city/getCities', 'Game\CityController@getCities');
	// Route::post('city/setScientists/{city}', 'Game\CityController@setScientists');
	// Route::post('city/setWine/{city}', 'Game\CityController@setWine');
	// Route::post('city/setName/{city}', 'Game\CityController@setName');

	// Route::post('user/getMessages', 'Game\UserController@getMessages');
	// Route::post('user/sendMessage/{city}', 'Game\UserController@sendMessage');
	// Route::post('user/deleteMessage', 'Game\UserController@deleteMessage');
	// Route::put('user/readMessages', 'Game\UserController@readMessages');
	// Route::put('user/readMessage/{message}', 'Game\UserController@readMessage');

	// Route::get('research', 'Game\ResearchController@getData');
	// Route::post('research/{research}', 'Game\ResearchController@create');
}
