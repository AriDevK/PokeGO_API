package main

import (
	m "PokeGO/models"
)

// GetPokemons Get all pokemon from database with limit and offset
// params for pagination purposes and returns a slice of pokemon
func GetPokemons(limit int, offset int) []m.Pokemon {
	var pokemon []m.Pokemon

	db.Preload("Generation").
		Preload("FirstType").
		Preload("SecondType").
		Limit(limit).
		Offset(offset).
		Find(&pokemon)

	return pokemon
}

// GetPokemonByType Get all pokemon from database by type name
// with limit and offset params for pagination purposes and returns a slice of pokemon
func GetPokemonByType(typeName string, limit int, offset int) []m.Pokemon {
	var pokemon []m.Pokemon
	var typeData m.Type

	db.Select("ID").
		Where(&m.Type{Name: typeName}).
		First(&typeData)

	db.Preload("Generation").
		Preload("FirstType").
		Preload("SecondType").
		Limit(limit).
		Offset(offset).
		Where(&m.Pokemon{FirstTypeId: typeData.ID}).
		Find(&pokemon)

	return pokemon
}

// GetPokemonByGeneration Get all pokemon from database by generation number
// with limit and offset params for pagination purposes and returns a slice of pokemon
func GetPokemonByGeneration(generation int, limit int, offset int) []m.Pokemon {
	var pokemon []m.Pokemon
	var generationData m.Generation

	db.Select("ID").
		Where(&m.Generation{Number: generation}).
		First(&generationData)

	db.Preload("Generation").
		Preload("FirstType").
		Preload("SecondType").
		Limit(limit).
		Offset(offset).
		Where(&m.Pokemon{GenerationId: generationData.ID}).
		Find(&pokemon)

	return pokemon
}

// GetPokemonByRegion Get all pokemon from database by region name
// with limit and offset params for pagination purposes and returns a slice of pokemon
func GetPokemonByRegion(regionName string, limit int, offset int) []m.Pokemon {
	var pokemon []m.Pokemon
	var generationData m.Generation

	db.Select("ID").
		Where(&m.Generation{Name: regionName}).
		First(&generationData)

	db.Preload("Generation").
		Preload("FirstType").
		Preload("SecondType").
		Limit(limit).
		Offset(offset).
		Where(&m.Pokemon{GenerationId: generationData.ID}).
		Find(&pokemon)

	return pokemon
}

// GetPokemonByNumber Get pokemon from database by pokedex number
// with limit and offset params for pagination purposes and returns a slice of pokemon
func GetPokemonByNumber(number int) m.Pokemon {
	var pokemon m.Pokemon
	db.Preload("Generation").
		Preload("FirstType").
		Preload("SecondType").
		Limit(1).
		Offset(0).
		Where(&m.Pokemon{PokedexNumber: uint(number)}).
		Take(&pokemon)

	return pokemon
}
