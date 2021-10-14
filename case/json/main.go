package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wbing441282413/gotest/case/json/entity"
)

func main() {

	fmt.Println("----------------------编码---------------------------")

	//json,多个结构体实体
	var movies = []entity.Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)

	//解码
	fmt.Println("----------------------解码---------------------------")
	var movies2 []entity.Movie
	if err = json.Unmarshal(data, &movies2); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(movies2)

	var actors []struct{ Actors []string }
	if err = json.Unmarshal(data, &actors); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(actors)

	var ree []struct{ released string }
	if err = json.Unmarshal(data, &ree); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(ree)
}
