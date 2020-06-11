package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	fmt.Printf("Starting application...\n\n\n")
	time.Sleep(time.Second)
}

type (
	// Define flowers struct
	flowers struct {
		name string
		thc  int
		geo  map[string]int // declare a map inside the struct
	}
	// Define cannabis struct
	cannabis struct {
		name    string
		flowers []flowers
	}
)

var (
	args = os.Args[1:]
	// Create new types of cannabis using struct type
	cSativa = cannabis{name: "Sativa"}
	cIndica = cannabis{name: "Indica"}
)

func main() {
	// create flowers examples
	createExamples()

	// check if user input any new flower
	// its like routing, if get (just read, without parameters, just show)
	// if it has arguments (POST), try to create a new flower
	// need to implement PATCH and PUT "routes", for updating an existent flower
	switch len(args) {
	case 0:
		break
	case 5:
		createFlowerFromUser()
	default:
		fmt.Println("[type] [flower] [thc] [country] [qtde]")
		return
	}
	// Printing the flowers for each type
	cSativa.printFlowers()
	fmt.Println("----------------------------")
	cIndica.printFlowers()
}

// Method to add new flowers for cannabis type
func (c *cannabis) updateFlower(n string, thc int, g map[string]int) error {
	var flag bool

	nF := flowers{
		name: n,
		thc:  thc,
		geo:  make(map[string]int), // Initializing map
	}
	nF.geo = g // Give a value to the map

	// check if flowers exists
	flag = c.checkFlowers(n)

	if flag == false {
		c.flowers = append(c.flowers, nF)
	} else {
		return errors.New("Flower alreay exists")
	}
	return nil // no errors, everything if fine to go
}

func (c *cannabis) checkFlowers(fName string) bool {
	for _, f := range c.flowers {
		if f.name != fName {
			continue
		} else {
			return true // if flower exists update flag to true
		}
	}
	return false
}

// Method to print all flowers of Cannabis
func (c *cannabis) printFlowers() {
	fmt.Printf("\nThe %s flowers are:\n", c.name)
	for _, f := range c.flowers {
		fmt.Printf("Name: %v ------- THC: %v\n", f.name, f.thc)
		for c, q := range f.geo {
			fmt.Printf("Country: %v ----- Qtde: %v\n\n", c, q)
		}
	}
}

// Function to create examples of geo (maps)
func createGeo(c string, q int) (map[string]int, error) {
	if q < 0 {
		return nil, errors.New("Number must be bigger or equal than 0") // create error and pass nil for map
	}

	return map[string]int{
		c: q, // country and qtde of this flower at this country
	}, nil // error
}

func createExamples() {
	// Create examples of geo
	eX, _ := createGeo("Brazil", 15) // ignoring the errors because its just an example
	eY, _ := createGeo("France", 200)

	// Add new sativa and indicas flowers
	_ = (&cSativa).updateFlower("Gorilla Haze", 27, eX) // go automatic use cSativa pointer instead the copy of objects itself (&cSativa) OR cSativa

	_ = cIndica.updateFlower("Notherland", 22, eY) // ignoring the err

}
func createFlowerFromUser() {
	cType, flower, country, q := args[0], args[1], args[3], args[4]

	thc, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalln(err)
		return
	}

	switch cType {
	// could be a search into db for all *cannabis types
	case "Sativa":
		qtde, err := strconv.Atoi(q)
		if err != nil {
			log.Fatalln(err)
			return
		}
		geo, _ := createGeo(country, qtde) // ignoring the error because at conversion I've already check the err
		err = cSativa.updateFlower(flower, thc, geo)
		if err != nil {
			log.Fatalln(err)
			return
		}
	case "Indica":
		qtde, err := strconv.Atoi(q)
		if err != nil {
			log.Fatalln(err)
			return
		}
		geo, _ := createGeo(country, qtde)
		err = cIndica.updateFlower(flower, thc, geo)
		if err != nil {
			log.Fatalln(err)
			return
		}
	default:
		// if type alreay exists, it is create
		// implement logic
	}
}
