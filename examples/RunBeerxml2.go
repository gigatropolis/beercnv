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
	"../../beercnv"
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
	beer2 := beercnv.BeerXml2{}
	filename := "Recipies\\xml\\include-hops.xml"
	//filename := "..\\Recipies\\xml\\include-hops.xml"
	//filename := "Recipies\\xml\\nhc_2015.xml"

	err := beercnv.AddFromBeerXMLFile(&beer2, filename)

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
