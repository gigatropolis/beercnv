// Package tar implements a way to read BeerXML files
// It aims to cover most of the variations

// References:
// http://www.beerxml.com/

package beercnv

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	//"io/ioutil"
	//"os"
	"regexp"
	"strconv"
)

func ConvertXML1to2(xml1 io.Reader, beer2 *BeerXml2) error {

	//beer2 := beerxml2.BeerXml2{}
	beer := BeerXml{}

	//filename := "Recipies\\xml\\nhc_2015.xml"
	//buf, err := ioutil.ReadFile(filename)

	//decoder := xml.NewDecoder(reader)
	//decoder.CharsetReader = charset.NewReaderLabel
	//err = decoder.Decode(&parsed)

	p := xml.NewDecoder(xml1)
	p.CharsetReader = charset.NewReaderLabel
	err := p.Decode(&beer)

	//err = xml.Unmarshal(buf, &beer)

	if err != nil {
		panic(err)
	}

	for _, recipe := range beer.Recipes {

		rec := BeerRecipe{}
		recIng := RecIngredients{}

		rec.Name = recipe.Name
		rec.Type = recipe.Type
		rec.Brewer = recipe.Brewer
		rec.AssistantBrewer = recipe.AssistantBrewer
		rec.BatchSize = recipe.BatchSize
		rec.BoilSize = recipe.BoilSize
		rec.BoilTime = recipe.BoilTime
		rec.Efficiency = recipe.Efficiency
		rec.Notes = recipe.Notes

		recOg := OriginalGravity{}
		recOg.Units = "sg"
		recOg.Og = recipe.Og
		rec.Og = recOg

		recFg := FinalGravity{}
		recFg.Units = "sg"
		recFg.Fg = recipe.Fg
		rec.Fg = recFg

		rec.DisplayBatchSize = recipe.DisplayBatchSize
		rec.DisplayBoilSize = recipe.DisplayBoilSize
		rec.DisplayOg = recipe.DisplayOg
		rec.DisplayFg = recipe.DisplayFg
		rec.DisplayPrimaryTemp = recipe.DisplayPrimaryTemp
		rec.DisplaySecondaryTemp = recipe.DisplaySecondaryTemp
		rec.DisplayTertiaryTemp = recipe.DisplayTertiaryTemp
		rec.DisplayAgeTemp = recipe.DisplayAgeTemp

		for _, hop := range recipe.Hops {

			recHop := HopAddition{}

			recHop.Name = hop.Name
			recHop.Origin = hop.Origin
			recHop.AlphaAcidUnits = hop.Alpha
			recHop.Use = hop.Use
			recHop.Form = hop.Form

			hopTime := UseTime{}
			hopTime.Units = "min"
			fTime, err := strconv.ParseFloat(hop.Time, 32)
			if err != nil {
				fmt.Println(err)
			}
			hopTime.Time = float32(fTime)
			recHop.Time = hopTime

			recMass := MassAmount{}
			recMass.Units = "Kg"
			recMass.Amount = hop.Amount
			recHop.Amount = recMass

			recHop.DisplayAmount = hop.DisplayAmount
			recHop.Inventory = hop.Inventory
			recHop.DisplayTime = hop.DisplayTime

			recIng.Hops = append(recIng.Hops, recHop)

			var pInvHop *InvHop
			pInvHop = getInventoryHop(beer2.HopVarieties, hop.Name)

			if pInvHop == nil {
				pInvHop = new(InvHop)
				pInvHop.Name = hop.Name
				pInvHop.Origin = hop.Origin
				pInvHop.AlphaAcidUnits = hop.Alpha
				pInvHop.BetaAcidUnits = hop.Beta

				pInvHop.Inventory.AddHopAmount(hop.Amount, "Kg", hop.Form)
				pInvHop.Type = hop.Type
				pInvHop.Notes = hop.Notes
				pInvHop.Humulene = hop.Humulene
				pInvHop.Caryophyllene = hop.Caryophyllene
				pInvHop.Cohumulone = hop.Cohumulone
				pInvHop.Myrcene = hop.Myrcene

				pInvHop.Inventory.AddHopAmount(hop.Amount, "Kg", hop.Form)
				beer2.HopVarieties = append(beer2.HopVarieties, *pInvHop)
			} else {

				pInvHop.Inventory.AddHopAmount(hop.Amount, "Kg", hop.Form)
			}

		}

		for _, ferm := range recipe.Fermentables {

			recFerm := FermAddition{}

			recFerm.Name = ferm.Name
			recFerm.Type = ferm.Type
			recFerm.Origin = ferm.Origin
			recFerm.Supplier = ferm.Supplier
			recFerm.AddAfterBoil = ferm.AddAfterBoil

			if ferm.Type == "Extract" {
				recFerm.Color.Units = "SRM"
			} else {
				recFerm.Color.Units = "L"
			}
			recFerm.Color.Color = ferm.Color

			recFerm.Amount.Units = "Kg"
			recFerm.Amount.Amount = ferm.Amount

			recFerm.DisplayAmount = ferm.DisplayAmount
			recFerm.Inventory = ferm.Inventory
			recFerm.DisplayColor = ferm.DisplayColor

			recIng.Fermentables = append(recIng.Fermentables, recFerm)

			var pInvFerm *InvFermentable
			pInvFerm = getInventoryFermentable(beer2.Fermentables, ferm.Name)

			if pInvFerm == nil {
				pInvFerm = new(InvFermentable)

				pInvFerm.Name = ferm.Name
				pInvFerm.Type = ferm.Type
				pInvFerm.Origin = ferm.Origin
				pInvFerm.Supplier = ferm.Supplier
				pInvFerm.Notes = ferm.Notes

				pInvFerm.Color.Units = recFerm.Color.Units
				pInvFerm.Color.Color = ferm.Color

				pInvFerm.YieldDryBasis.FineGrind = ferm.Yield
				pInvFerm.YieldDryBasis.CoarseFineDiff = ferm.CoarseFineDiff
				pInvFerm.Notes = ferm.Notes
				pInvFerm.Moisture = ferm.Moisture
				pInvFerm.DiastaticPower = ferm.DiastaticPower
				pInvFerm.Protein = ferm.Protein
				pInvFerm.MaxInBatch = ferm.MaxInBatch
				pInvFerm.RecommendMash = ferm.RecommendMash
				pInvFerm.IbuGalPerLb = ferm.IbuGalPerLb
				pInvFerm.Potential = ferm.Potential

				pInvFerm.Inventory.AddFermentationAmount(ferm.Amount, "Kg")
				beer2.Fermentables = append(beer2.Fermentables, *pInvFerm)

			} else {
				pInvFerm.Inventory.AddFermentationAmount(ferm.Amount, "Kg")
			}
		}

		for _, misc := range recipe.Miscs {

			recMisc := MiscAdditions{}

			recMisc.Name = misc.Name
			recMisc.Type = misc.Type
			recMisc.Use = misc.Use

			if misc.AmountIsWeight {
				recMisc.Amount.Units = "Kg"
				recMisc.Amount.Amount = misc.Amount
			} else {
				recMisc.AmountAsWeight.Units = "l"
				recMisc.AmountAsWeight.Weight = misc.Amount
			}

			recMisc.Time.Units = "min"
			recMisc.Time.Time = misc.Time

			recMisc.DisplayAmount = misc.DisplayAmount
			recMisc.Inventory = misc.Inventory
			recMisc.DisplayTime = misc.DisplayTime

			recIng.Miscs = append(recIng.Miscs, recMisc)

			var pInvMisc *InvMisc
			pInvMisc = getInventoryMisc(beer2.Miscs, misc.Name)

			if pInvMisc == nil {

				pInvMisc = new(InvMisc)

				pInvMisc.Name = misc.Name
				pInvMisc.Type = misc.Type
				pInvMisc.Use = misc.Use
				pInvMisc.UseFor = misc.UseFor
				pInvMisc.Notes = misc.Notes
				if misc.AmountIsWeight {
					pInvMisc.Inventory.AddMiscMassAmount(misc.Amount, "Kg")
				} else {
					pInvMisc.Inventory.AddMiscVolAmount(misc.Amount, "l")
				}

				beer2.Miscs = append(beer2.Miscs, *pInvMisc)
			} else {
				if misc.AmountIsWeight {
					pInvMisc.Inventory.AddMiscMassAmount(misc.Amount, "Kg")
				} else {
					pInvMisc.Inventory.AddMiscVolAmount(misc.Amount, "l")
				}
			}

		}

		for _, water := range recipe.Waters {

			recWater := WaterAddition{}

			recWater.Name = water.Name
			recWater.Calcium = water.Calcium
			recWater.Bicarbonate = water.Bicarbonate
			recWater.Sulfate = water.Sulfate
			recWater.Chloride = water.Chloride
			recWater.Sodium = water.Sodium
			recWater.Magnesium = water.Magnesium
			recWater.Amount.Units = "l"
			recWater.Amount.Amount = water.Amount

			recWater.DisplayAmount = water.DisplayAmount

			recIng.Waters = append(recIng.Waters, recWater)

			var pInvWater *WaterProfile
			pInvWater = getInventoryWater(beer2.Profiles, water.Name)

			if pInvWater == nil {

				pInvWater = new(WaterProfile)
				pInvWater.Name = water.Name
				pInvWater.Calcium = water.Calcium
				pInvWater.Bicarbonate = water.Bicarbonate
				pInvWater.Sulfate = water.Sulfate
				pInvWater.Chloride = water.Chloride
				pInvWater.Sodium = water.Sodium
				pInvWater.Magnesium = water.Magnesium
				pInvWater.Ph = water.Ph
				pInvWater.Notes = water.Notes
				beer2.Profiles = append(beer2.Profiles, *pInvWater)
			}
		}

		for _, yeast := range recipe.Yeasts {

			recYeast := YeastAdditions{}

			recYeast.Name = yeast.Name
			recYeast.Type = yeast.Type
			recYeast.Form = yeast.Form
			recYeast.Laboratory = yeast.Laboratory
			recYeast.ProductID = yeast.ProductID
			recYeast.AddToSecondary = yeast.AddToSecondary

			if yeast.AmountIsWeight {
				recYeast.AmountAsWeight.Units = "Kg"
				recYeast.AmountAsWeight.Weight = yeast.Amount
			} else {
				recYeast.Amount.Units = "l"
				recYeast.Amount.Amount = yeast.Amount
			}

			recYeast.TimesCultured = yeast.TimesCultured

			recYeast.DisplayAmount = yeast.DisplayAmount
			recYeast.DispMinTemp = yeast.DispMinTemp
			recYeast.DispMaxTemp = yeast.DispMaxTemp
			recYeast.Inventory = yeast.Inventory
			recYeast.CultureDate = yeast.CultureDate

			recIng.Yeasts = append(recIng.Yeasts, recYeast)

			var pInvYeast *InvYeast
			pInvYeast = getInventoryYeast(beer2.Cultures, yeast.Name)

			if pInvYeast == nil {

				pInvYeast = new(InvYeast)

				pInvYeast.Name = yeast.Name
				pInvYeast.Type = yeast.Type
				pInvYeast.Form = yeast.Form
				pInvYeast.Laboratory = yeast.Laboratory
				pInvYeast.ProductID = yeast.ProductID

				pInvYeast.TemperatureRange.Minimum.Degrees = "C"
				pInvYeast.TemperatureRange.Minimum.Temp = yeast.MaxTemperature
				pInvYeast.TemperatureRange.Maximum.Degrees = "C"
				pInvYeast.TemperatureRange.Maximum.Temp = yeast.MaxTemperature

				pInvYeast.Flocculation = yeast.Flocculation
				pInvYeast.Attenuation = yeast.Attenuation
				pInvYeast.Notes = yeast.Notes
				pInvYeast.BestFor = yeast.BestFor
				pInvYeast.MaxReuse = yeast.MaxReuse
				pInvYeast.BestFor = yeast.BestFor

				if yeast.CultureDate == "" {
					pInvYeast.Inventory.Culture.Date = yeast.CultureDate
				}

				if yeast.AmountIsWeight {
					pInvYeast.Inventory.Dry.Units = "Kg"
					pInvYeast.Inventory.Dry.Amount += yeast.Amount
				} else {
					pInvYeast.Inventory.Liquid.Units = "l"
					pInvYeast.Inventory.Liquid.Amount += yeast.Amount
				}
				beer2.Cultures = append(beer2.Cultures, *pInvYeast)
			} else {

				if yeast.AmountIsWeight {
					pInvYeast.Inventory.Dry.Units = "Kg"
					pInvYeast.Inventory.Dry.Amount += yeast.Amount
				} else {
					pInvYeast.Inventory.Liquid.Units = "l"
					pInvYeast.Inventory.Liquid.Amount += yeast.Amount
				}
			}
		}

		recStyle := StyleAddition{}

		recStyle.Name = recipe.Style.Name
		recStyle.Category = recipe.Style.Category
		recStyle.CategoryNumber = recipe.Style.CategoryNumber
		recStyle.StyleLetter = recipe.Style.StyleLetter
		recStyle.StyleGuide = recipe.Style.StyleGuide
		recStyle.Type = recipe.Style.Type

		recStyle.DisplayOgMin = recipe.Style.DisplayOgMin
		recStyle.DisplayOgMax = recipe.Style.DisplayOgMax
		recStyle.DisplayFgMin = recipe.Style.DisplayFgMin
		recStyle.DisplayFgMax = recipe.Style.DisplayFgMax
		recStyle.DisplayColorMin = recipe.Style.DisplayColorMin
		recStyle.DisplayColorMax = recipe.Style.DisplayColorMax

		rec.Style = recStyle

		var pInvStyle *StyleProfile
		pInvStyle = getInventoryStyle(beer2.Styles, recipe.Style.Name)

		if pInvStyle == nil {
			pInvStyle = new(StyleProfile)
			pInvStyle.Name = recipe.Style.Name
			pInvStyle.Category = recipe.Style.Category
			pInvStyle.CategoryNumber = recipe.Style.CategoryNumber
			pInvStyle.StyleLetter = recipe.Style.StyleLetter
			pInvStyle.StyleGuide = recipe.Style.StyleGuide
			pInvStyle.Type = recipe.Style.Type

			pInvStyle.Og.Minimum.Units = "sg"
			pInvStyle.Og.Minimum.Minimum = recipe.Style.OgMin
			pInvStyle.Og.Maximum.Units = "sg"
			pInvStyle.Og.Maximum.Maximum = recipe.Style.OgMax
			pInvStyle.Fg.Minimum.Units = "sg"
			pInvStyle.Fg.Minimum.Minimum = recipe.Style.FgMin
			pInvStyle.Fg.Maximum.Units = "sg"
			pInvStyle.Fg.Maximum.Maximum = recipe.Style.FgMax

			pInvStyle.IBU.Minimum = recipe.Style.IbuMin
			pInvStyle.IBU.Maximum = recipe.Style.IbuMax
			pInvStyle.Color.Minimum.Units = "SRM"
			pInvStyle.Color.Maximum.Units = "SRM"
			pInvStyle.Color.Minimum.Color = recipe.Style.ColorMin
			pInvStyle.Color.Maximum.Color = recipe.Style.ColorMax
			pInvStyle.Carbonation.Minimum = recipe.Style.CarbMin
			pInvStyle.Carbonation.Maximum = recipe.Style.CarbMax
			pInvStyle.ABV.Minimum = recipe.Style.AbvMin
			pInvStyle.ABV.Maximum = recipe.Style.AbvMax

			pInvStyle.Notes = recipe.Style.Notes
			pInvStyle.Profile = recipe.Style.Profile
			pInvStyle.Ingredients = recipe.Style.Ingredients
			pInvStyle.Examples = recipe.Style.Examples

			beer2.Styles = append(beer2.Styles, *pInvStyle)
		}

		mash := recipe.Mash

		rec.Mash.Name = mash.Name
		rec.Mash.GrainTemp.Units = "C"
		rec.Mash.GrainTemp.Degrees = mash.GrainTemp
		rec.Mash.SpargeTemp.Units = "C"
		rec.Mash.SpargeTemp.Degrees = mash.SpargeTemp
		rec.Mash.Ph = mash.Ph
		rec.Mash.Notes = mash.Notes

		rec.Mash.DisplayGrainTemp = mash.DisplayGrainTemp
		rec.Mash.DisplayTunTemp = mash.DisplayTunTemp
		rec.Mash.DisplaySpargeTemp = mash.DisplaySpargeTemp
		rec.Mash.DisplayTunWeight = mash.DisplayTunWeight

		for _, mashStep := range recipe.Mash.MashSteps {

			recMashStep := RecMashStep{}

			recMashStep.Name = mashStep.Name
			recMashStep.Type = mashStep.Type
			recMashStep.InfuseAmount.Units = "L"
			recMashStep.InfuseAmount.Amount = mashStep.InfuseAmount
			recMashStep.StepTemp.Units = "C"
			recMashStep.StepTemp.Degrees = mashStep.StepTemp
			recMashStep.StepTime.Units = "min"
			recMashStep.StepTime.Time = mashStep.StepTime
			recMashStep.RampTime.Units = "min"
			recMashStep.RampTime.Time = mashStep.RampTime
			recMashStep.EndTemp.Units = "C"
			recMashStep.EndTemp.Degrees = mashStep.EndTemp
			recMashStep.Description = mashStep.Description
			recMashStep.WaterGrainRatio = mashStep.WaterGrainRatio

			re := regexp.MustCompile(`([0-9]+\.[0-9]+)\s*(\w+)`)
			infuseMatch := re.FindAllStringSubmatch(mashStep.InfuseTemp, 1)
			decoctionMatch := re.FindAllStringSubmatch(mashStep.DecotionAmt, 1)

			if infuseMatch != nil {
				recMashStep.InfuseTemp.Units = infuseMatch[0][2]
				iInfuse, err := strconv.ParseFloat(infuseMatch[0][1], 32)
				if err != nil {
					recMashStep.InfuseTemp.Degrees = float32(iInfuse)
				}
			}

			if decoctionMatch != nil {
				recMashStep.DecotionAmt.Units = decoctionMatch[0][2]
				iDecoction, err := strconv.ParseFloat(decoctionMatch[0][1], 32)
				if err != nil {
					recMashStep.DecotionAmt.Amount = float32(iDecoction)
				}
			}

			recMashStep.DisplayStepTemp = mashStep.DisplayStepTemp
			recMashStep.DisplayInfuseAmt = mashStep.DisplayInfuseAmt

			rec.Mash.MashSteps = append(rec.Mash.MashSteps, recMashStep)
		}

		rec.Ingredients = recIng

		beer2.Recipes = append(beer2.Recipes, rec)
	}

	return nil
}
