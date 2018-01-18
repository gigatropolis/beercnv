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
	//filename := "Recipies\\xml\\nhc_2015.xml"

	err := beerxml2.AddFromBeerXMLFile(&beer2, filename)

	if err != nil {
		panic(err)
	}

	recStyle := StyleAddition{}

	recStyle.Name = recipe.Style.Name
	recStyle.Category = recipe.Style.Category
	recStyle.CategoryNumber = recipe.Style.Name
	recStyle.Name = recipe.Style.CategoryNumber
	recStyle.StyleLetter = recipe.Style.StyleLetter
	recStyle.StyleGuide = recipe.Style.StyleGuide
	recStyle.Type = recipe.Style.Type

	recIng.Style = recStyle

	var pInvStyle *Style = nil
	pInvStyle = getInventoryStyle(beer2.Styles, yeast.Name)

	if pInvStyle == nil {
		pInvStyle = new(Style)
		pInvStyle.Name = recipe.Style.Name
		pInvStyle.Category = recipe.Style.Category
		pInvStyle.CategoryNumber = recipe.Style.Name
		pInvStyle.Name = recipe.Style.CategoryNumber
		pInvStyle.StyleLetter = recipe.Style.StyleLetter
		pInvStyle.StyleGuide = recipe.Style.StyleGuide
		pInvStyle.Type = recipe.Style.Type

		pInvStyle.Og.Minimum.Density = "sg"
		pInvStyle.Og.Minimum.Minimum = recipe.Style.OgMin
		pInvStyle.Og.Maximum.Density = "sg"
		pInvStyle.Og.Maximum.Maximum = recipe.Style.OgMax
		pInvStyle.Fg.Minimum.Density = "sg"
		pInvStyle.Fg.Minimum.Minimum = recipe.Style.FgMin
		pInvStyle.Fg.Maximum.Density = "sg"
		pInvStyle.Fg.Maximum.Maximum = recipe.Style.FgMax
	}

	output, err := xml.MarshalIndent(beer2, "  ", "   ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)
	fmt.Println(" ")
}
