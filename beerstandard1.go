// Package tar implements a way to read BeerXML files
// It aims to cover most of the variations

// References:
// http://www.beerxml.com/

package beercnv

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// BeerXML holds a slice of Recipes
type BeerXML struct {
	XMLName xml.Name `xml:"RECIPES"`
	Recipes []Recipe `xml:"RECIPE"`
}

// Recipe implements a BeerXML recipe including the different childs.
type Recipe struct {
	XMLName              xml.Name      `xml:"RECIPE"`
	Name                 string        `xml:"NAME"`
	Version              string        `xml:"VERSION"`
	Type                 string        `xml:"TYPE"`
	Brewer               string        `xml:"BREWER"`
	AssistantBrewer      string        `xml:"ASST_BREWER"`
	BatchSize            float32       `xml:"BATCH_SIZE"`
	BoilSize             float32       `xml:"BOIL_SIZE"`
	BoilTime             int           `xml:"BOIL_TIME"`
	Efficiency           float32       `xml:"EFFICIENCY"`
	Hops                 []Hop         `xml:"HOPS>HOP"`
	Fermentables         []Fermentable `xml:"FERMENTABLES>FERMENTABLE"`
	Miscs                []Misc        `xml:"MISCS>MISC"`
	Yeasts               []Yeast       `xml:"YEASTS>YEAST"`
	Waters               []Water       `xml:"WATERS>WATER"`
	Style                Style         `xml:"STYLE"`
	Equipment            Equipment     `xml:"EQUIPMENT"`
	Mash                 Mash          `xml:"MASH"`
	Notes                string        `xml:"NOTES"`
	TasteNotes           string        `xml:"TASTE_NOTES"`
	TasteRating          float32       `xml:"TASTE_RATING"`
	Og                   float32       `xml:"OG"`
	Fg                   float32       `xml:"FG"`
	Carbonation          float32       `xml:"CARBONATION"`
	FermentationStages   int           `xml:"FERMENTATION_STAGES"`
	PrimaryAge           float32       `xml:"PRIMARY_AGE"`
	PrimaryTemp          float32       `xml:"PRIMARY_TEMP"`
	SecondaryAge         float32       `xml:"SECONDARY_AGE"`
	SecondaryTemp        float32       `xml:"SECONDARY_TEMP"`
	TertiaryAge          float32       `xml:"TERTIARY_AGE"`
	Age                  string        `xml:"AGE"`
	AgeTemp              float32       `xml:"AGE_TEMP"`
	CarbonationUsed      string        `xml:"CARBONATION_USED"`
	PrimingSugarName     string        `xml:"PRIMING_SUGAR_NAME"`
	PrimingSugarEquiv    float32       `xml:"PRIMING_SUGAR_EQUIV"`
	KegPrimingFactor     float32       `xml:"KEG_PRIMING_FACTOR"`
	CarbonationTemp      float32       `xml:"CARBONATIOn_TEMP"`
	DisplayCarbTemp      string        `xml:"DISPLAY_CARB_TEMP"`
	Date                 string        `xml:"DATE"`
	EstOg                string        `xml:"EST_OG"`
	EstFg                string        `xml:"EST_FG"`
	EstColor             string        `xml:"EST_COLOR"`
	Ibu                  string        `xml:"IBU"`
	IbuMethod            string        `xml:"Tinseth"`
	EstAbv               string        `xml:"EST_ABV"`
	Abv                  string        `xml:"ABV"`
	ActualEfficiency     string        `xml:"ACTUAL_EFFICIENCY"`
	Calories             string        `xml:"CALORIES"`
	DisplayBatchSize     string        `xml:"DISPLAY_BATCH_SIZE"`
	DisplayBoilSize      string        `xml:"DISPLAY_BOIL_SIZE"`
	DisplayOg            string        `xml:"DISPLAY_OG"`
	DisplayFg            string        `xml:"DISPLAY_FG"`
	DisplayPrimaryTemp   string        `xml:"DISPLAY_PRIMARY_TEMP"`
	DisplaySecondaryTemp string        `xml:"DISPLAY_SECONDARY_TEMP"`
	DisplayTertiaryTemp  string        `xml:"DISPLAY_TERTIARY_TEMP"`
	DisplayAgeTemp       string        `xml:"DISPLAY_AGE_TEMP"`
}

// Hop definition
type Hop struct {
	XMLName       xml.Name `xml:"HOP"`
	Name          string   `xml:"NAME"`
	Version       string   `xml:"VERSION"`
	Origin        string   `xml:"ORIGIN"`
	Alpha         float32  `xml:"ALPHA"`
	Beta          float32  `xml:"BETA"`
	Amount        float32  `xml:"AMOUNT"`
	Use           string   `xml:"USE"`
	Time          string   `xml:"TIME"`
	Notes         string   `xml:"NOTES"`
	Type          string   `xml:"TYPE"`
	Form          string   `xml:"FORM"`
	Hsi           string   `xml:"HSI"`
	Humulene      float32  `xml:"HUMULENE"`
	Caryophyllene float32  `xml:"CARYOPHYLLENE"`
	Cohumulone    float32  `xml:"COHUMULONE"`
	Myrcene       float32  `xml:"MYRCENE"`
	DisplayAmount string   `xml:"DISPLAY_AMOUNT"`
	Inventory     string   `xml:"INVENTORY"`
	DisplayTime   string   `xml:"DISPLAY_TIME"`
}

// Fermentable definition
type Fermentable struct {
	XMLName           xml.Name `xml:"FERMENTABLE"`
	Name              string   `xml:"NAME"`
	Version           int      `xml:"VERSION"`
	Type              string   `xml:"TYPE"`
	Amount            float32  `xml:"AMOUNT"`
	Yield             float32  `xml:"YIELD"`
	Color             float32  `xml:"COLOR"`
	AddAfterBoil      bool     `xml:"ADD_AFTER_BOIL"`
	Origin            string   `xml:"ORIGIN"`
	Supplier          string   `xml:"SUPPLIER"`
	Notes             string   `xml:"NOTES"`
	CoarseFineDiff    float32  `xml:"COARSE_FINE_DIFF"`
	Moisture          float32  `xml:"MOISTURE"`
	DiastaticPower    float32  `xml:"DIASTATIC_POWER"`
	Protein           float32  `xml:"PROTEIN"`
	MaxInBatch        float32  `xml:"MAX_IN_BATCH"`
	RecommendMash     bool     `xml:"RECOMMEND_MASH"`
	IbuGalPerLb       float32  `xml:"IBU_GAL_PER_LB"`
	DisplayAmount     string   `xml:"DISPLAY_AMOUNT"`
	Inventory         string   `xml:"INVENTORY"`
	Potential         float32  `xml:"POTENTIAL"`
	DisplayColor      string   `xml:"DISPLAY_COLOR"`
	ExtractSubstitute string   `xml:"EXTRACT_SUBSTITUTE"`
}

// Yeast definition
type Yeast struct {
	XMLName        xml.Name `xml:"YEAST"`
	Name           string   `xml:"NAME"`
	Version        int      `xml:"VERSION"`
	Type           string   `xml:"TYPE"`
	Form           string   `xml:"FROM"`
	Amount         float32  `xml:"AMOUNT"`
	AmountIsWeight bool     `xml:"AMOUNT_IS_WEIGHT"`
	Laboratory     string   `xml:"LABORATORY"`
	ProductID      string   `xml:"PRODUCT_ID"`
	MinTemperature float32  `xml:"MIN_TEMPERATURE"`
	MaxTemperature float32  `xml:"MAX_TEMPERATURE"`
	Flocculation   string   `xml:"FLOCCULATION"`
	Attenuation    float32  `xml:"ATTENUATION"`
	Notes          string   `xml:"NOTES"`
	BestFor        string   `xml:"BEST_FOR"`
	MaxReuse       int      `xml:"MAX_REUSE"`
	TimesCultured  int      `xml:"TIMES_CULTURED"`
	AddToSecondary bool     `xml:"ADD_TO_SECONDARY"`
	DisplayAmount  string   `xml:"DISPLAY_AMOUNT"`
	DispMinTemp    string   `xml:"DISP_MIN_TEMP"`
	DispMaxTemp    string   `xml:"DISP_MAX_TEMP"`
	Inventory      string   `xml:"INVENTORY"`
	CultureDate    string   `xml:"CULTURE_DATE"`
}

//Style defines beer styles
type Style struct {
	XMLName         xml.Name `xml:"STYLE"`
	Name            string   `xml:"NAME"`
	Version         int      `xml:"VERSION"`
	Category        string   `xml:"CATEGORY"`
	CategoryNumber  int      `xml:"CATEGORY_NUMBER"`
	StyleLetter     string   `xml:"STYLE_LETTER"`
	StyleGuide      string   `xml:"STYLE_GUIDE"`
	Type            string   `xml:"TYPE"`
	OgMin           float32  `xml:"OG_MIN"`
	OgMax           float32  `xml:"OG_MAX"`
	FgMin           float32  `xml:"FG_MIN"`
	FgMax           float32  `xml:"FG_MAX"`
	IbuMin          float32  `xml:"IBU_MIN"`
	IbuMax          float32  `xml:"IBU_MAX"`
	ColorMin        float32  `xml:"COLOR_MIN"`
	ColorMax        float32  `xml:"COLOR_MAX"`
	CarbMin         float32  `xml:"CARB_MIN"`
	CarbMax         float32  `xml:"CARB_MAX"`
	AbvMax          float32  `xml:"ABV_MAX"`
	AbvMin          float32  `xml:"ABV_MIN"`
	Notes           string   `xml:"NOTES"`
	Profile         string   `xml:"PROFILE"`
	Ingredients     string   `xml:"INGREDIENTS"`
	Examples        string   `xml:"EXAMPLES"`
	DisplayOgMin    string   `xml:"DISPLAY_OG_MIN"`
	DisplayOgMax    string   `xml:"DISPLAY_OG_MAX"`
	DisplayFgMin    string   `xml:"DISPLAY_FG_MIN"`
	DisplayFgMax    string   `xml:"DISPLAY_FG_MAX"`
	DisplayColorMin string   `xml:"DISPLAY_COLOR_MIN"`
	DisplayColorMax string   `xml:"DISPLAY_COLOR_MAX"`
	OgRange         string   `xml:"OG_RANGE"`
	FgRange         string   `xml:"FG_RANGE"`
	IbuRange        string   `xml:"IBU_RANGE"`
	CarbRange       string   `xml:"CARB_RANGE"`
	ColorRange      string   `xml:"COLOR_RANGE"`
	AbvRange        string   `xml:"ABV_RANGE"`
}

// Equipment profiles for brewin beer
type Equipment struct {
	XMLName                xml.Name `xml:"EQUIPMENT"`
	Name                   string   `xml:"NAME"`
	Version                int      `xml:"VERSION"`
	BoilSize               float32  `xml:"BOIL_SIZE"`
	BatchSize              float32  `xml:"BATCH_SIZE"`
	TunVolume              float32  `xml:"TUN_VOLUME"`
	TunWeight              float32  `xml:"TUN_WEIGHT"`
	TunSpecificHeat        float32  `xml:"TUN_SPECIFIC_HEAT"`
	TopUpWater             float32  `xml:"TOP_UP_WATER"`
	TrubChillerLoss        float32  `xml:"TRUB_CHILLER_LOSS"`
	EvapRate               float32  `xml:"EVAP_RATE"`
	BoilTime               float32  `xml:"BOIL_TIME"`
	CalcBoilVolume         bool     `xml:"CALC_BOIL_VOLUME"`
	LauterDeadspace        float32  `xml:"LAUTER_DEADSPACE"`
	TopUpKettle            float32  `xml:"TOP_UP_KETTLE"`
	HopUtilization         float32  `xml:"HOP_UTILIZATION"`
	CoolingLossPct         float32  `xml:"COOLING_LOSS_PCT"`
	Notes                  string   `xml:"NOTES"`
	DisplayBoilSize        string   `xml:"DISPLAY_BOIL_SIZE"`
	DisplayBatchSize       string   `xml:"DISPLAY_BATCH_SIZE"`
	DisplayTunVolume       string   `xml:"DISPLAY_TUN_VOLUME"`
	DisplayTunWeight       string   `xml:"DISPLAY_TUN_WEIGHT"`
	DisplayTopUpWater      string   `xml:"DISPLAY_TOP_UP_WATER"`
	DiplayTrubChillerLoss  string   `xml:"DISPLAY_TRUB_CHILLER_LOSS"`
	DisplayLauterDeadspace string   `xml:"DISPLAY_LAUTER_DEADSPACE"`
	DisplayTopUpKettle     string   `xml:"DISPLAY_TOP_UP_KETTLE"`
}

// Mash is struct defining mash type
type Mash struct {
	XMLName           xml.Name   `xml:"MASH"`
	Name              string     `xml:"NAME"`
	Version           int        `xml:"VERSION"`
	GrainTemp         float32    `xml:"GRAIN_TEMP"`
	TunTemp           float32    `xml:"TUN_TEMP"`
	SpargeTemp        float32    `xml:"SPARGE_TEMP"`
	Ph                float32    `xml:"PH"`
	TunWeight         float32    `xml:"TUN_WEIGHT"`
	TunSpecificHeat   float32    `xml:"TUN_SPECIFIC_HEAT"`
	EquipAdjust       bool       `xml:"EQUIP_ADJUST"`
	Notes             string     `xml:"NOTES"`
	DisplayGrainTemp  string     `xml:"DISPLAY_GRAIN_TEMP"`
	DisplayTunTemp    string     `xml:"DISPLAY_TUN_TEMP"`
	DisplaySpargeTemp string     `xml:"DISPLAY_SPARGE_TEMP"`
	DisplayTunWeight  string     `xml:"DISPLAY_TUN_WEIGHT"`
	MashSteps         []MashStep `xml:"MASH_STEPS>MASH_STEP"`
}

// MashStep ia single step for a mash profile
type MashStep struct {
	XMLName          xml.Name `xml:"MASH_STEP"`
	Name             string   `xml:"NAME"`
	Version          int      `xml:"VERSION"`
	Type             string   `xml:"TYPE"`
	InfuseAmount     float32  `xml:"INFUSE_AMOUNT"`
	StepTime         float32  `xml:"STEP_TIME"`
	StepTemp         float32  `xml:"STEP_TEMP"`
	RampTime         float32  `xml:"RAMP_TIME"`
	EndTemp          float32  `xml:"END_TEMP"`
	Description      string   `xml:"DESCRIPTION"`
	WaterGrainRatio  string   `xml:"WATER_GRAIN_RATIO"`
	DecotionAmt      string   `xml:"DECOTION_AMT"`
	InfuseTemp       string   `xml:"INFUSE_TEMP"`
	DisplayStepTemp  string   `xml:"DISPLAY_STEP_TEMP"`
	DisplayInfuseAmt string   `xml:"DISPLAY_INFUSE_AMT"`
}

// Water definition used in brewing process
type Water struct {
	XMLName       xml.Name `xml:"WATER"`
	Name          string   `xml:"NAME"`
	Version       int      `xml:"VERSION"`
	Amount        float32  `xml:"AMOUNT"`
	Calcium       float32  `xml:"CALCIUM"`
	Bicarbonate   float32  `xml:"BICARBONATE"`
	Sulfate       float32  `xml:"SULFATE"`
	Chloride      float32  `xml:"CHLORIDE"`
	Sodium        float32  `xml:"SODIUM"`
	Magnesium     float32  `xml:"MAGNESIUM"`
	Ph            float32  `xml:"PH"`
	Notes         string   `xml:"NOTES"`
	DisplayAmount string   `xml:"DISPLAY_AMOUNT"`
}

// Misc is miscellaneous ingredients used in beer
type Misc struct {
	XMLName        xml.Name `xml:"MISC"`
	Name           string   `xml:"NAME"`
	Version        int      `xml:"VERSION"`
	Type           string   `xml:"TYPE"`
	Use            string   `xml:"USE"`
	Amount         float32  `xml:"AMOUNT"`
	Time           float32  `xml:"TIME"`
	AmountIsWeight bool     `xml:"AMOUNT_IS_WEIGHT"`
	UseFor         string   `xml:"USE_FOR"`
	Notes          string   `xml:"NOTES"`
	DisplayAmount  string   `xml:"DISPLAY_AMOUNT"`
	Inventory      string   `xml:"INVENTORY"`
	DisplayTime    string   `xml:"DISPLAY_TIME"`
}

// NewBeerXML takes a io.Reader and returns Recipes
func NewBeerXML(r io.Reader, pBXml *BeerXML) error {
	dec := xml.NewDecoder(r)

	// dec.CharsetReader = CharsetReader
	if err := dec.Decode(pBXml); err != nil {
		return err
	}
	return nil
}

// NewBeerXMLFromFile takes a filename as string and returns Recipes
func NewBeerXMLFromFile(f string) (BeerXML, error) {

	beer := BeerXML{}
	xmlFile, err := os.Open(f)
	if err != nil {
		return beer, err
	}
	defer xmlFile.Close()

	err2 := NewBeerXML(xmlFile, &beer)
	if err2 != nil {
		return beer, err2
	}

	for _, b := range beer.Recipes {

		for i, hop := range b.Hops {
			if hop.Type == "" {
				b.Hops[i].Type = "Both"
			}
		}

		for i, ferm := range b.Fermentables {

			fermType := strings.ToLower(ferm.Type)

			if fermType == "liquid extract" {
				b.Fermentables[i].Type = "Extract"
				continue
			}

			if fermType == "base malt" ||
				fermType == "kilned malt" ||
				fermType == "caramel/crystal malt" ||
				fermType == "roasted malt" {
				b.Fermentables[i].Type = "Grain"
				continue
			}
		}
	}

	return beer, nil
}

// TextSummary returns a string with a summary of Recipes including fermentables and hops
func (bxml *BeerXML) TextSummary() string {
	buf := ""
	for x := range bxml.Recipes {
		buf += fmt.Sprintf("Recipe (%d) : %s \n", x, bxml.Recipes[x].Name)
		buf += fmt.Sprintf("Type: %s\n", bxml.Recipes[x].Type)
		buf += fmt.Sprintf("Batch Size: %f\n", bxml.Recipes[x].BatchSize)
		buf += fmt.Sprintf("Boil Size: %f\n", bxml.Recipes[x].BoilSize)
		buf += fmt.Sprintf("Boil Time: %d\n", bxml.Recipes[x].BoilTime)
		for f := range bxml.Recipes[x].Fermentables {
			buf += fmt.Sprintf("Fermentable: %d : %s : %s\n", f, bxml.Recipes[x].Fermentables[f].Name,
				bxml.Recipes[x].Fermentables[f].DisplayAmount)
		}
		for h := range bxml.Recipes[x].Hops {
			buf += fmt.Sprintf("Hops %d : %s : %s\n", h, bxml.Recipes[x].Hops[h].Name,
				bxml.Recipes[x].Hops[h].DisplayAmount)
		}
		buf += "\n"
	}
	return buf

}
