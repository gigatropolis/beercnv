package main

import (
	"encoding/xml"

	"fmt"
	//"io/ioutil"
	"os"
	//"strconv"
	//"github.com/beerxml"
	//"github.com/stone/beerxML2"

	//"../xml/beerxml"
	"../../beercnv"
)

// $.map($('.recipe-link'), m=>m.href).forEach(recipe=>{window.open(`${recipe}.xml`, '_blank')})

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
	beer2 := beercnv.BeerXML2{}
	//filename := "Recipies\\xml\\include-hops.xml"
	filename := "Recipies\\xml\\include-hops.xml"
	//filename := "Recipies\\xml\\nhc_2015.xml"

	xmlFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	err = beercnv.ConvertXML1to2(xmlFile, &beer2)

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
