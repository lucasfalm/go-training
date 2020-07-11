package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	fmt.Printf("Starting application...\n\n\n")
	time.Sleep(time.Second)

	// create flowers examples
	createExamples()
}

type (
	flowers struct {
		name   string
		thc    int
		origin map[string]int
	}
	cannabis struct {
		name    string
		flowers []flowers
	}
	flowerCollection interface {
		showAllFlowers()
	}
)

var (
	cSativa           = cannabis{name: "Sativa"}
	cIndica           = cannabis{name: "Indica"}
	flowersCollection []flowerCollection
	args              = os.Args[1:]
	divisionLine      = func() { fmt.Println(strings.Repeat("-", 50)) }
)

func main() {
	switch len(args) {
	case 0:
		break
	case 5:
		createFlowerFromUser()
	default:
		fmt.Println("[type] [flowers] [thc] [country] [qtde]")
		return
	}

	flowersCollection = append(flowersCollection, &cSativa, &cIndica) // pass pointer because who implements printFlower() is a pointer to *cannabis

	fmt.Printf("\t\t C A N N A B I S:\n")
	divisionLine()
	for _, cannabis := range flowersCollection {
		cannabis.showAllFlowers()
	}
}

// Method to add new flowers for cannabis type
func (c *cannabis) updateFlower(n string, thc int, origin map[string]int) error {
	var flag bool

	if thc <= 0 {
		return errors.New("THC must be bigger or equal than 0")
	}

	newFlower := flowers{
		name:   n,
		thc:    thc,
		origin: make(map[string]int),
	}
	newFlower.origin = origin

	if flag = c.flowerExists(n); !flag {
		c.flowers = append(c.flowers, newFlower)
	} else {
		return errors.New("Flower alreay exists")
	}

	return nil // no errors, everything if fine to go
}

func (c *cannabis) flowerExists(flowerName string) bool {
	for _, f := range c.flowers {
		if f.name != flowerName {
			continue
		} else {
			return true // if flowers exists update flag to true
		}
	}
	return false
}

func (c *cannabis) showAllFlowers() {
	fmt.Printf("\nThe %s flowers are:\n", c.name)
	for _, f := range c.flowers {
		fmt.Printf("Name: %v \t ------- \tTHC: %v\n", f.name, f.thc)
		for cannabis, qtde := range f.origin {
			fmt.Printf("Country: %v \t ------- \tQtde: %v\n\n", cannabis, qtde)
		}
	}
	divisionLine()
}

// Function to create examples of origin (maps)
func createGeo(c string, q int) (map[string]int, error) {
	if q < 0 {
		return nil, errors.New("Quantity must be bigger or equal than 0") // create error and pass nil for map
	}

	return map[string]int{
		c: q, // country and qtde of this flowers at this country
	}, nil // error
}

func createExamples() {
	// Create examples of origin
	eX, _ := createGeo("Brazil", 15) // ignoring the errors because its just an example
	eY, _ := createGeo("France", 200)

	// Add new sativa and indicas flowers
	_ = (&cSativa).updateFlower("Gorilla Haze", 27, eX) // go automatic use cSativa pointer instead the copy of objects itself (&cSativa) OR cSativa

	_ = cIndica.updateFlower("Notherland", 22, eY) // ignoring the err

}
func createFlowerFromUser() {
	// get input from user (or request)
	cType, flowers, country, q := args[0], args[1], args[3], args[4]

	thc, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalln(err)
		return
	}

	// could be a search in the db for all *cannabis types
	// need to implement some logic to clear out this
	switch cType {
	case "Sativa":
		qtde, err := strconv.Atoi(q)
		if err != nil {
			log.Fatalln(err)
			return
		}

		origin, err := createGeo(country, qtde)
		if err != nil { // check if number is negative
			log.Fatalln(err)
			return
		}

		err = cSativa.updateFlower(flowers, thc, origin)
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

		origin, err := createGeo(country, qtde)
		if err != nil {
			log.Fatalln(err)
			return
		}

		err = cIndica.updateFlower(flowers, thc, origin)
		if err != nil {
			log.Fatalln(err)
			return
		}
	default:
		// if type alreay exists, it is create
		// implement logic
		// flowersCollection = append(flowersCollection, newType) for insert in menu
	}
}
