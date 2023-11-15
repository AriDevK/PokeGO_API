package main

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

// GetAllPokemonHandler godoc
// @Summary Get all pokemon
// @Description  Get all pokemon
// @Tags pokemon
// @Produce json
// @Router / [get]
// @Param limit query int false "int valid" default(1) minimum(1) maximum(1000) example(1) "Limit Number"
// @Param offset query int false "int valid" default(1) minimum(0) maximum(1000) example(1) "Offset Number"
func GetAllPokemonHandler(c echo.Context) error {
	var limit = 10000
	var offset = 0

	if qLimit := c.QueryParam("limit"); qLimit != "" {
		l, err := strconv.Atoi(qLimit)
		if err != nil {
			return c.JSON(400, err)
		}
		limit = l
	}

	if qOffset := c.QueryParam("offset"); qOffset != "" {
		o, err := strconv.Atoi(qOffset)
		if err != nil {
			return c.JSON(400, err)
		}
		offset = o
	}

	pokemon := GetPokemons(limit, offset)
	return c.JSON(200, pokemon)
}

// GetPokemonByPokedexHandler godoc
// @Summary Get all pokemon by pokedex number
// @Description Get all pokemon by pokedex number
// @Tags pokemon
// @Produce json
// @Router /pokedex/{pokedexNumber} [get]
// @Param pokedexNumber path int false "int valid" default(1) minimum(1) maximum(898) example(1) "Pokedex Number
func GetPokemonByPokedexHandler(c echo.Context) error {
	var pokedexNumber = -1

	if c.Param("pokedexNumber") != "" {
		pkNumber, err := strconv.Atoi(c.Param("pokedexNumber"))
		if err != nil {
			return c.JSON(400, err)
		}
		pokedexNumber = pkNumber
	}

	pokemon := GetPokemonByNumber(pokedexNumber)
	return c.JSON(200, pokemon)
}

// GetPokemonByTypeHandler godoc
// @Summary Get all pokemon by type
// @Description Get all pokemon by type
// @Tags pokemon
// @Produce json
// @Router /type/{name} [get]
// @Param name path string false "string enums" Enums(Fire, Water, Grass, Electric, Ice, Fighting, Poison, Ground, Flying, Psychic, Bug, Rock, Ghost, Dark, Dragon, Steel, Fairy) "Type Name"
// @Param limit query int false "int valid" default(1) minimum(1) maximum(1000) example(1) "Limit Number"
// @Param offset query int false "int valid" default(1) minimum(0) maximum(1000) example(0) "Offset Number"
func GetPokemonByTypeHandler(c echo.Context) error {
	var limit = 10000
	var offset = 0

	if qLimit := c.QueryParam("limit"); qLimit != "" {
		l, err := strconv.Atoi(qLimit)
		if err != nil {
			return c.JSON(400, err)
		}
		limit = l
	}

	if qOffset := c.QueryParam("offset"); qOffset != "" {
		o, err := strconv.Atoi(qOffset)
		if err != nil {
			return c.JSON(400, err)
		}
		offset = o
	}

	pokemons := GetPokemonByType(c.Param("name"), limit, offset)
	return c.JSON(200, pokemons)
}

// GetPokemonByGenerationHandler godoc
// @Summary Get all pokemon by generation number
// @Description Get all pokemon by generation number
// @Tags pokemon
// @Produce json
// @Router /generation/{number} [get]
// @Param number path int false "int valid" default(1) minimum(1) maximum(8) example(1) "Generation Number"
// @Param limit query int false "int valid" default(1) minimum(1) maximum(1000) example(1) "Limit Number"
// @Param offset query int false "int valid" default(1) minimum(0) maximum(1000) example(1) "Offset Number"
func GetPokemonByGenerationHandler(c echo.Context) error {
	var limit = 10000
	var offset = 0
	var generation = -1

	if qLimit := c.QueryParam("limit"); qLimit != "" {
		l, err := strconv.Atoi(qLimit)
		if err != nil {
			return c.JSON(400, err)
		}
		limit = l
	}

	if qOffset := c.QueryParam("offset"); qOffset != "" {
		o, err := strconv.Atoi(qOffset)
		if err != nil {
			return c.JSON(400, err)
		}
		offset = o
	}

	if c.Param("number") != "" {
		gen, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			return c.JSON(400, err)
		}
		generation = gen
	}

	pokemons := GetPokemonByGeneration(generation, limit, offset)
	return c.JSON(200, pokemons)
}

// GetPokemonByRegionHandler godoc
// @Summary Get all pokemon by region name
// @Description Get all pokemon by region name
// @Tags pokemon
// @Produce json
// @Router /region/{name} [get]
// @Param name path string false "string enums" Enums(Kanto, Johto, Hoenn, Sinnoh, Unova, Kalos, Alola, Galar) "Region Name"
// @Param limit query int false "int valid" default(1) minimum(1) maximum(1000) example(1) "Limit Number"
// @Param offset query int false "int valid" default(1) minimum(0) maximum(1000) example(1) "Offset Number"
func GetPokemonByRegionHandler(c echo.Context) error {
	var limit = 10000
	var offset = 0

	if qLimit := c.QueryParam("limit"); qLimit != "" {
		l, err := strconv.Atoi(qLimit)
		if err != nil {
			return c.JSON(400, err)
		}
		limit = l
	}

	if qOffset := c.QueryParam("offset"); qOffset != "" {
		o, err := strconv.Atoi(qOffset)
		if err != nil {
			return c.JSON(400, err)
		}
		offset = o
	}

	pokemons := GetPokemonByRegion(c.Param("name"), limit, offset)
	return c.JSON(200, pokemons)
}
