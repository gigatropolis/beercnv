package main

import (
	"encoding/xml"

	"fmt"
	//"io/ioutil"
	"os"
	//"strconv"
	//"github.com/beerxml"
	//"github.com/stone/beerxml2"

	//"../xml/beerxml"
	"../xml/beerxml2"
)

func main() {

	/*
		bxml, err := beerxml.NewBeerXmlFromFile("testfiles/recipes.xml")
		if err != nil {
			fmt.Printf(" %s", err.Error())
		}

		if bxml.Recipes[0].Name != "Burton Ale" {
			fmt.Printf("Recipe.Name: %s", err.Error())
		}
	*/
	beer2 := beerxml2.BeerXml2{}
	filename := "Recipies\\xml\\include-hops.xml"
	//filename := "..\\Recipies\\xml\\include-hops.xml"
	//filename := "Recipies\\xml\\nhc_2015.xml"

	err := beerxml2.AddFromBeerXMLFile(&beer2, filename)

	if err != nil {
		panic(err)
	}

	output, err := xml.MarshalIndent(beer2, "  ", "   ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)
	fmt.Println(" ")
}
