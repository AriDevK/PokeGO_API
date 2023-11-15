package main

import (
	c "PokeGO/config"
	_ "PokeGO/docs/swagger"
	m "PokeGO/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var configuration c.Configurations

// @title Pokemon API
// @version 1.0
// @description This is a sample server for Pokemon API.
// @termsOfService http://swagger.io/terms/

// @contact.name Ariadne Rangel aka AriDevK
// @contact.url AriDev.me
// @contact.email AriadneFave@hotmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	initConfiguration()
	initDatabase()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "️🐿 method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", GetAllPokemonHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/pokedex/:pokedexNumber", GetPokemonByPokedexHandler)
	e.GET("/generation/:number", GetPokemonByGenerationHandler)
	e.GET("/region/:name", GetPokemonByRegionHandler)
	e.GET("/type/:name", GetPokemonByTypeHandler)

	e.Logger.Fatal(e.Start(configuration.Server.Port))
}

// initConfiguration reads the config file and unmarshal it into the configuration struct
func initConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		panic(err)
	}
}

// initDatabase initializes the database connection and migrates the models
func initDatabase() {
	tx, err := gorm.Open(sqlserver.Open(configuration.Database.ConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = tx
	err = db.AutoMigrate(&m.Pokemon{}, &m.Type{}, &m.Generation{})
	if err != nil {
		log.Fatal(err)
	}
}
