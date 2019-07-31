package main

import (
	//"github.com/gigatropolis/beercnv"

	"encoding/xml"
	"io/ioutil"
	"log"

	"fmt"
	"os"
	"path/filepath"
	//"strconv"
	//"github.com/beerxml"
	//"github.com/stone/beerxml2"

	//"../xml/beerxml"
	"../../beercnv"
)

// $.map($('.recipe-link'), m=>m.href).forEach(recipe=>{window.open(`${recipe}.xml`, '_blank')})

func main() {
	allXML := beercnv.BeerXML{}
	//beer2 := beercnv.BeerXml2{}
	//path := "../recipes/public"
	//path := "../recipes/recipe3/recipe3"
	path := "../recipes/all"
	//path := "../recipes/home"

	outName := "AllRecipes1.xml"
	outName2 := "AllRecipes2.xml"
	//outName := "home.xml"
	//outName2 := "home2.xml"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".xml" {

			//fmt.Println(file.Name())
			pathXML := path + "/" + file.Name()

			xmlFile, err := os.Open(pathXML)

			if err != nil {
				panic(err)
			}
			beer, err := beercnv.NewBeerXMLFromFile(pathXML)
			if err != nil {
				fmt.Printf("no xml object from %s", file.Name())
				continue
			}
			for _, b := range beer.Recipes {

				allXML.Recipes = append(allXML.Recipes, b)
			}
			xmlFile.Close()

		}
	}

	output, err := xml.MarshalIndent(allXML, "", "   ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	out, err3 := os.Create(outName)
	if err3 != nil {
		panic(err3.Error())
	}
	out.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))
	out.Write(output)

	out.Close()

	beer2 := beercnv.BeerXML2{}

	xmlFile, err := os.Open(outName)

	if err != nil {
		panic(err)
	}

	err = beercnv.ConvertXML1to2(xmlFile, &beer2)

	if err != nil {
		panic(err)
	}

	output2, err := xml.MarshalIndent(beer2, "  ", "   ")

	if err != nil {
		panic(err.Error())
	}

	out2, err4 := os.Create(outName2)
	if err4 != nil {
		panic(err4.Error())
	}

	out2.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))
	out2.Write(output2)
	defer out2.Close()

}
