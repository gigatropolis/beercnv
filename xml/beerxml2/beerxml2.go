// Package tar implements a way to read BeerXML files
// It aims to cover most of the variations

// References:
// http://www.beerxml.com/

package beerxml2

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type MinScale struct {
	XMLName xml.Name `xml:"minimum"`
	Scale   string   `xml:"scale,attr"`
	Minimum float32  `xml:",chardata"`
}

type MaxScale struct {
	XMLName xml.Name `xml:"maximum"`
	Scale   string   `xml:"scale,attr"`
	Maximum float32  `xml:",chardata"`
}

type ColorScale struct {
	Minimum MinScale `xml:"minimum"`
	Maximum MaxScale `xml:"maximum"`
}

type VolAmount struct {
	XMLName xml.Name `xml:"amount"`
	Volume  string   `xml:"volume,attr"`
	Amount  float32  `xml:",chardata"`
}

type WeightAmount struct {
	XMLName xml.Name `xml:"amount_as_weight"`
	Mass    string   `xml:"mass,attr"`
	Weight  float32  `xml:",chardata"`
}

type OriginalGravity struct {
	XMLName xml.Name `xml:"original_gravity"`
	Density string   `xml:"density,attr"`
	Og      float32  `xml:",chardata"`
}

type FinalGravity struct {
	XMLName xml.Name `xml:"final_gravity"`
	Density string   `xml:"density,attr"`
	Fg      float32  `xml:",chardata"`
}

// Recipes holds a slice of Rrecipes
type BeerXml2 struct {
	XMLName      xml.Name      `xml:"beer_xml"`
	Version      string        `xml:"version"`
	HopVarieties []Hop         `xml:"hop_varieties>hop"`
	Fermentables []Fermentable `xml:"fermentables>fermentable"`
	Miscs        []Misc        `xml:"miscellaneous_ingredients>miscellaneous"`
	Cultures     []Yeast       `xml:"cultures>yeast"`
	Styles       []Style       `xml:"styles>style"`
	Profiles     []Water       `xml:"profiles>water"`
	Procedures   []Mash        `xml:"procedure>mash"`
	Recipes      []Recipe      `xml:"recipes>recipe"`
}

type RecIngredients struct {
	Hops         []HopAddition    `xml:"hop_bill>hop"`
	Fermentables []FermAddition   `xml:"grain_bill>fermentable"`
	Miscs        []MiscAdditions  `xml:"adjuncts>miscellaneous"`
	Yeasts       []YeastAdditions `xml:"yeast_additions>yeast"`
	Waters       []WaterAddition  `xml:"water_profile>water"`
	Equipment    []EquipmentUsed  `xml:"Equipment,omitempty"`
}

// Recipe implements a BeerXML recipe including the different childs.
type Recipe struct {
	XMLName         xml.Name        `xml:"recipe"`
	Name            string          `xml:"name"`
	Type            string          `xml:"type"`
	Brewer          string          `xml:"brewer"`
	AssistantBrewer string          `xml:"assistant_brewer"`
	BatchSize       string          `xml:"batch_size"`
	BoilSize        string          `xml:"boil_size"`
	BoilTime        string          `xml:"boil_time"`
	Efficiency      float32         `xml:"efficiency"`
	Style           StyleAddition   `xml:"style"`
	Ingredients     RecIngredients  `xml:"ingredients"`
	Mash            Mash            `xml:"mash"`
	Notes           string          `xml:"notes"`
	Og              OriginalGravity `xml:"original_gravity"`
	Fg              FinalGravity    `xml:"final_gravity"`
}

type InvLeaf struct {
	XMLName xml.Name `xml:"leaf"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardate"`
}

type InvPellet struct {
	XMLName xml.Name `xml:"pellet"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardate"`
}

type InvPlug struct {
	XMLName xml.Name `xml:"plug"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardate"`
}

type InventoryHop struct {
	Leaf   InvLeaf   `xml:"leaf"`
	Pellet InvPellet `xml:"pellet"`
	Plug   InvPlug   `xml:"plug"`
}

type Hop struct {
	XMLName        xml.Name     `xml:"hop"`
	Name           string       `xml:"name"`
	Origin         string       `xml:"origin"`
	AlphaAcidUnits float32      `xml:"alpha_acid_units"`
	BetaAcidUnits  float32      `xml:"beta_acid_units"`
	Type           string       `xml:"type"`
	Notes          string       `xml:"notes"`
	PercentLost    float32      `xml:"percent_lost"`
	Substitutes    string       `xml:"substitutes"`
	Humulene       float32      `xml:"humulene"`
	Caryophyllene  float32      `xml:"caryophyllene"`
	Cohumulone     float32      `xml:"cohumulone"`
	Myrcene        float32      `xml:"myrcene"`
	Inventory      InventoryHop `xml:"inventory"`
}

type MassAmount struct {
	XMLName xml.Name `xml:"amount"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardata"`
}

type HopTime struct {
	XMLName  xml.Name `xml:"time"`
	Duration string   `xml:"duration,attr"`
	Time     float32  `xml:",chardata"`
}

type HopAddition struct {
	XMLName        xml.Name   `xml:"hop"`
	Name           string     `xml:"name"`
	Origin         string     `xml:"origin"`
	AlphaAcidUnits float32    `xml:"alpha_acid_units"`
	BetaAcidUnits  float32    `xml:"beta_acid_units"`
	Form           string     `xml:"form"`
	Use            string     `xml:"use"`
	Amount         MassAmount `xml:"amount"`
	Time           HopTime    `xml:"time"`
}

type Yield struct {
	FineGrind      string `xml:"fine_grind"`
	CoarseFineDiff string `xml:"fine_coarse_difference"`
}

type Fermentable struct {
	XMLName        xml.Name   `xml:"fermentable"`
	Name           string     `xml:"name"`
	Type           string     `xml:"type"`
	Color          ColorScale `xml:"color"`
	Origin         string     `xml:"origin"`
	Supplier       string     `xml:"supplier"`
	YieldDryBasis  Yield      `xml:"yield_dry_basis"`
	Notes          string     `xml:"notes"`
	Moisture       float32    `xml:"moisture"`
	DiastaticPower float32    `xml:"diastatic_power"`
	Protein        float32    `xml:"protien"`
	MaxInBatch     float32    `xml:"max_in_batch"`
	RecommendMash  bool       `xml:"recommended_mash"`
	IbuGalPerLb    float32    `xml:"ibu_gal_per_lb"`
	Inventory      string     `xml:"inventory"`
	Potential      float32    `xml:"potential"`
}

type FermAddition struct {
	XMLName      xml.Name   `xml:"fermentable"`
	Name         string     `xml:"name"`
	Type         string     `xml:"type"`
	Color        ColorScale `xml:"color"`
	Origin       string     `xml:"origin"`
	Supplier     string     `xml:"supplier"`
	Amount       MassAmount `xml:"amount"`
	AddAfterBoil bool       `xml:"add_after_boil"`
}

type AlcTolerence struct {
	Minimum float32 `xml:"minimum"`
	Maximum float32 `xml:"maximum"`
}

type LiquidAmount struct {
	XMLName xml.Name `xml:"liquid"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardata"`
}

type DryAmount struct {
	XMLName xml.Name `xml:"dry"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardata"`
}

type SlantAmount struct {
	XMLName xml.Name `xml:"slant"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardata"`
}

type CultureAmount struct {
	XMLName xml.Name `xml:"culture"`
	Mass    string   `xml:"mass,attr"`
	Amount  float32  `xml:",chardata"`
}

type YeastInventory struct {
	Liquid  LiquidAmount  `xml:"liquid"`
	Dry     DryAmount     `xml:"dry"`
	Slant   SlantAmount   `xml:"slant"`
	Culture CultureAmount `xml:"culture"`
}

type MinTemp struct {
	XMLName xml.Name `xml:"minimum"`
	Temp    string   `xml:"Temp,attr"`
	Minimum float32  `xml:",chardata"`
}

type MaxTemp struct {
	XMLName xml.Name `xml:"maximum"`
	Temp    string   `xml:"Temp,attr"`
	Maximum float32  `xml:",chardata"`
}

type TempRange struct {
	Minimum MinTemp `xml:"minimum"`
	Maximum MaxTemp `xml:"maximum"`
}

type Yeast struct {
	XMLName          xml.Name       `xml:"yeast"`
	Name             string         `xml:"name"`
	Type             string         `xml:"type"`
	Form             string         `xml:"form"`
	Laboratory       string         `xml:"laboratory"`
	ProductID        string         `xml:"product_id"`
	TemperatureRange TempRange      `xml:"temperature_range"`
	Flocculation     string         `xml:"flocculation"`
	Attenuation      float32        `xml:"attenuation"`
	AlcoholTolerece  AlcTolerence   `xml:"alcohol_tolerence"`
	Notes            string         `xml:"notes"`
	BestFor          string         `xml:"best_for"`
	MaxReuse         int            `xml:"max_reuse"`
	Inventory        YeastInventory `xml:"inventory"`
}

type YeastAdditions struct {
	XMLName        xml.Name     `xml:"yeast"`
	Name           string       `xml:"name"`
	Type           string       `xml:"type"`
	Form           string       `xml:"form"`
	Laboratory     string       `xml:"laboratory"`
	ProductID      string       `xml:"product_id"`
	Amount         VolAmount    `xml:"amount"`
	AmountAsWeight WeightAmount `xml:"amount_as_weight"`
	TimesCultured  int          `xml:"times_cultured"`
	AddToSecondary bool         `xml:"add_to_secondary"`
}

type MinDensity struct {
	XMLName xml.Name `xml:"minimum"`
	Density string   `xml:"density,attr"`
	Minimum float32  `xml:",chardata"`
}

type MaxDensity struct {
	XMLName xml.Name `xml:"maximum"`
	Density string   `xml:"density,attr"`
	Maximum float32  `xml:",chardata"`
}

type Gravity struct {
	Minimum MinDensity `xml:"minimum"`
	Maximum MaxDensity `xml:"maximum"`
}

type Bitterness struct {
	Minimum float32 `xml:"minimum"`
	Maximum float32 `xml:"maximum"`
}

type StyleCarb struct {
	Minimum float32 `xml:"minimum"`
	Maximum float32 `xml:"maximum"`
}

type StyleABV struct {
	Minimum float32 `xml:"minimum"`
	Maximum float32 `xml:"maximum"`
}

type Style struct {
	XMLName        xml.Name   `xml:"style"`
	Name           string     `xml:"name"`
	Category       string     `xml:"category"`
	CategoryNumber int        `xml:"category_number"`
	StyleLetter    string     `xml:"style_letter"`
	StyleGuide     string     `xml:"style_guide"`
	Type           string     `xml:"type"`
	Og             Gravity    `xml:"original_gravity"`
	Fg             Gravity    `xml:"final_gravity"`
	IBU            Bitterness `xml:"international_bitterness_units"`
	Color          ColorScale `xml:"color"`
	Carbonation    StyleCarb  `xml:"carbonation"`
	ABV            StyleABV   `xml:"alcohol_by_volume"` // Testfile from beerxml contains 2.1>
	Notes          string     `xml:"notes"`
	Profile        string     `xml:"profile"`
	Ingredients    string     `xml:"ingredients"`
	Examples       string     `xml:"examples"`
}

type StyleAddition struct {
	XMLName        xml.Name `xml:"style"`
	Name           string   `xml:"name"`
	Category       string   `xml:"category"`
	CategoryNumber int      `xml:"category_number"`
	StyleLetter    string   `xml:"style_letter"`
	StyleGuide     string   `xml:"style_guide"`
	Type           string   `xml:"type"`
}
type EquipmentUsed struct {
	XMLName                xml.Name `xml:"equipment"`
	Name                   string   `xml:"name"`
	Version                int      `xml:"version"`
	BoilSize               float32  `xml:"boil_size"`
	BatchSize              float32  `xml:"batch_size"`
	TunVolume              float32  `xml:"tun_volume"`
	TunWeight              float32  `xml:"tun_weight"`
	TunSpecificHeat        float32  `xml:"tun_specific_heat"`
	TopUpWater             float32  `xml:"top_up_water"`
	TrubChillerLoss        float32  `xml:"trub_chiller_loss"`
	EvapRate               float32  `xml:"evap_rate"`
	BoilTime               float32  `xml:"boil_time"`
	CalcBoilVolume         bool     `xml:"calc_boil_volume"`
	LauterDeadspace        float32  `xml:"lauter_deadspace"`
	TopUpKettle            float32  `xml:"top_up_kettle"`
	HopUtilization         float32  `xml:"hop_utilization"`
	CoolingLossPct         float32  `xml:"cooling_loss_pct"`
	Notes                  string   `xml:"notes"`
	DisplayBoilSize        string   `xml:"display_boil_size"`
	DisplayBatchSize       string   `xml:"display_batch_size"`
	DisplayTunVolume       string   `xml:"display_tun_volume"`
	DisplayTunWeight       string   `xml:"display_tun_weight"`
	DisplayTopUpWater      string   `xml:"display_top_up_water"`
	DiplayTrubChillerLoss  string   `xml:"display_trub_chiller_loss"`
	DisplayLauterDeadspace string   `xml:"display_lauter_deadspace"`
	DisplayTopUpKettle     string   `xml:"display_top_up_kettle"`
}

type InfuseVol struct {
	XMLName xml.Name `xml:"infuse_amount"`
	Volume  string   `xml:"volume,attr"`
	Amount  float32  `xml:",chardata"`
}

type DecVol struct {
	XMLName xml.Name `xml:"decoction_volume"`
	Volume  string   `xml:"volume,attr"`
	Amount  float32  `xml:",chardata"`
}

type StepDur struct {
	XMLName  xml.Name `xml:"step_time"`
	Duration string   `xml:"duration,attr"`
	Time     float32  `xml:",chardata"`
}

type RampDur struct {
	XMLName  xml.Name `xml:"step_time"`
	Duration string   `xml:"duration,attr"`
	Time     float32  `xml:",chardata"`
}

type StepDeg struct {
	XMLName xml.Name `xml:"step_tempurature"`
	Degrees string   `xml:"degrees,attr"`
	Time    float32  `xml:",chardata"`
}

type EndDeg struct {
	XMLName xml.Name `xml:"end_tempurature"`
	Degrees string   `xml:"degrees,attr"`
	Time    float32  `xml:",chardata"`
}

type InfuseDeg struct {
	XMLName xml.Name `xml:"infuse_tempurature"`
	Degrees string   `xml:"degrees,attr"`
	Time    float32  `xml:",chardata"`
}

type MashStep struct {
	XMLName         xml.Name  `xml:"mash_step"`
	Name            string    `xml:"name"`
	Type            string    `xml:"type"`
	InfuseAmount    InfuseVol `xml:"infuse_amount"`
	StepTemp        StepDeg   `xml:"step_tempurature"`
	StepTime        StepDur   `xml:"step_time"`
	RampTime        RampDur   `xml:"ramp_time"`
	EndTemp         EndDeg    `xml:"end_tempurature"`
	Description     string    `xml:"description"`
	WaterGrainRatio string    `xml:"water_grain_ratio"`
	DecotionAmt     DecVol    `xml:"decoction_amount"`
	InfuseTemp      InfuseDeg `xml:"infuse_tempurature"`
}

type GrainDeg struct {
	XMLName xml.Name `xml:"grain_tempurature"`
	Degrees string   `xml:"degrees,attr"`
	Time    float32  `xml:",chardata"`
}

type SpargeDeg struct {
	XMLName xml.Name `xml:"sparge_temperature"`
	Degrees string   `xml:"degrees,attr"`
	Time    float32  `xml:",chardata"`
}

type Mash struct {
	XMLName    xml.Name   `xml:"mash"`
	Name       string     `xml:"name"`
	GrainTemp  GrainDeg   `xml:"grain_tempurature"`
	SpargeTemp SpargeDeg  `xml:"sparge_temperature"`
	Ph         float32    `xml:"pH"`
	Notes      string     `xml:"notes"`
	MashSteps  []MashStep `xml:"mash_steps>mash_step"`
}

type Water struct {
	XMLName     xml.Name `xml:"water"`
	Name        string   `xml:"name"`
	Calcium     float32  `xml:"calcium"`
	Bicarbonate float32  `xml:"bicarbonate"`
	Sulfate     float32  `xml:"sulfate"`
	Chloride    float32  `xml:"chloride"`
	Sodium      float32  `xml:"sodium"`
	Magnesium   float32  `xml:"magnesium"`
	Ph          float32  `xml:"pH"`
	Notes       string   `xml:"notes"`
}

type WaterAddition struct {
	XMLName     xml.Name `xml:"water"`
	Name        string   `xml:"name"`
	Calcium     float32  `xml:"calcium"`
	Bicarbonate float32  `xml:"bicarbonate"`
	Sulfate     float32  `xml:"sulfate"`
	Chloride    float32  `xml:"chloride"`
	Sodium      float32  `xml:"sodium"`
	Magnesium   float32  `xml:"magnesium"`
}

type InventoryMisc struct {
	Amount         float32 `xml:"amount"`
	AmountAsWeight bool    `xml:"amount_as_weight"`
}

type Misc struct {
	XMLName   xml.Name      `xml:"miscellaneous"`
	Name      string        `xml:"name"`
	Type      string        `xml:"type"`
	Use       string        `xml:"use"`
	UseFor    string        `xml:"use_for"`
	Notes     string        `xml:"notes"`
	Inventory InventoryMisc `xml:"inventory"`
}

type MiscAdditions struct {
	XMLName        xml.Name     `xml:"miscellaneous"`
	Name           string       `xml:"name"`
	Type           string       `xml:"type"`
	Use            string       `xml:"use"`
	Amount         VolAmount    `xml:"amount"`
	AmountAsWeight WeightAmount `xml:"amount_as_weight"`
	Time           string       `xml:"time"`
}

// NewBeerXml takes a io.Reader and returns Recipes
func NewBeerXml(r io.Reader) (bxml *BeerXml2, err error) {
	dec := xml.NewDecoder(r)
	//dec.CharsetReader = CharsetReader
	if err := dec.Decode(&bxml); err != nil {
		return nil, err
	}
	return bxml, nil
}

// NewBeerXmlFromFile takes a filename as string and returns Recipes
func NewBeerXmlFromFile(f string) (bxml *BeerXml2, err error) {
	xmlFile, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	return NewBeerXml(xmlFile)
}

// TextSummary returns a string with a summary of Recipes including fermentables and hops
func (bxml *BeerXml2) TextSummary() string {
	buf := ""
	for x := range bxml.Recipes {
		buf += fmt.Sprintf("Recipe (%d) : %s \n", x, bxml.Recipes[x].Name)
		buf += fmt.Sprintf("Type: %s\n", bxml.Recipes[x].Type)
		buf += fmt.Sprintf("Batch Size: %s\n", bxml.Recipes[x].BatchSize)
		buf += fmt.Sprintf("Boil Size: %s\n", bxml.Recipes[x].BoilSize)
		buf += fmt.Sprintf("Boil Time: %s\n", bxml.Recipes[x].BoilTime)
		/*
			for f := range bxml.Recipes[x].Fermentables {
				buf += fmt.Sprintf("Fermentable: %d : %s : %s\n", f, bxml.Recipes[x].Fermentables[f].Name,
					bxml.Recipes[x].Fermentables[f].DisplayAmount)
			}
			for h := range bxml.Recipes[x].Hops {
				buf += fmt.Sprintf("Hops %d : %s : %s\n", h, bxml.Recipes[x].Hops[h].Name,
					bxml.Recipes[x].Hops[h].DisplayAmount)
			}
		*/
		buf += "\n"
	}
	return buf

}
