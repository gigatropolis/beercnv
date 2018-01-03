package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	//"github.com/beerxml"
	//"github.com/stone/beerxml2"

	"../xml/beerxml"
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
	beer2.Version = "2.0"
	//hop := beerxml2.Hop{}
	//beer2.HopVarieties = append(beer2.HopVarieties, hop)

	// scale := beerxml2.MinScale{}
	// scale.Minimum = 0

	beer := beerxml.BeerXml{}

	//filename := "include-hops.xml"
	//filename := "2017-09-24 UB Blonde Test #3_xml.xml"
	filename := "Recipies\\xml\\nhc_2015.xml"
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(buf, &beer)

	if err != nil {
		panic(err)
	}

	for _, recipe := range beer.Recipes {

		rec := beerxml2.Recipe{}
		recIng := beerxml2.RecIngredients{}

		rec.Name = recipe.Name
		rec.Type = recipe.Type
		rec.Brewer = recipe.Brewer
		rec.AssistantBrewer = recipe.AssistantBrewer
		rec.BatchSize = recipe.BatchSize
		rec.BoilSize = recipe.BoilSize
		rec.BoilTime = recipe.BoilTime
		rec.Efficiency = recipe.Efficiency
		rec.Notes = recipe.Notes

		recOg := beerxml2.OriginalGravity{}
		recOg.Density = "sg"
		recOg.Og = recipe.Og
		rec.Og = recOg

		recFg := beerxml2.FinalGravity{}
		recFg.Density = "sg"
		recFg.Fg = recipe.Fg
		rec.Fg = recFg

		for _, hop := range recipe.Hops {

			invHop := beerxml2.Hop{}
			invHop.Inventory = beerxml2.InventoryHop{}
			recHop := beerxml2.HopAddition{}

			recHop.Name = hop.Name
			recHop.Origin = hop.Origin
			recHop.AlphaAcidUnits = hop.Alpha
			recHop.Use = hop.Use
			recHop.Form = hop.Form

			hopTime := beerxml2.HopTime{}
			hopTime.Duration = "min"
			fTime, err := strconv.ParseFloat(hop.Time, 32)
			if err != nil {
				fmt.Println(err)
			}
			hopTime.Time = float32(fTime)
			recHop.Time = hopTime

			recMass := beerxml2.MassAmount{}
			recMass.Mass = "Kg"
			recMass.Amount = hop.Amount
			recHop.Amount = recMass

			recIng.Hops = append(recIng.Hops, recHop)

			recHop.Name = hop.Name
			invHop.Origin = hop.Origin
			invHop.AlphaAcidUnits = hop.Alpha
			invHop.BetaAcidUnits = hop.Bets
			invHop.Form = hop.Form
			invHop.Type = hop.Type
			invHop.Notes = hop.Notes
			invHop.Humulene = hop.Humulene
			invHop.Caryophyllene = hop.Caryophyllene
			invHop.Cohumulone = hop.Cohumulone
			invHop.Myrcene = hop.Myrcene

			fmt.Printf("HOP:%s amt:%f t: %s\n", hop.Name, hop.Amount, hop.Time)
		}

		fmt.Printf("HopCount = %d", len(recIng.Hops))
		rec.Ingredients = recIng

		beer2.Recipes = append(beer2.Recipes, rec)
	}

	output, err := xml.MarshalIndent(beer2, "  ", "   ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)
	fmt.Println(" ")
}
