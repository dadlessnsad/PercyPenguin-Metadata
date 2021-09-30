package metadata

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/PercyPernguin-Metadata/app/config"
	"github.com/PercyPernguin-Metadata/structs"
)

const POLYMORPH_IMAGE_URL string = ""
const EXTERNAL_URL string = ""
const GENES_COUNT = 
const BACKGROUND_GENES_COUNT int =
const BASE_GENES_COUNT int = 
const EYEWEAR_GENES_COUNT int =
const FRONTITEM_GENES_COUNT int =
const HANDS_GENES_COUNT int =

type Genome string
type gene int
type StringAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type IntegerAttribute struct {
	TraitType string `json:"trait_type"`
	Value     int    `json:"value"`
}

type FloatAttribute struct {
	TraitType   string  `json:"trait_type"`
	Value       float64 `json:"value"`
	DisplayType string  `json:"display_type"`
}

func (g Gene) toPath() string {
	if g < 10 {
		return fmt.Sprintf("0%s", strconv.Itoa(int(g)))
	}

	return strconv.Itoa(int(g))
}

func getGeneInt(g string, start, end, count int) int {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	return gene % count
}

func getHandsGene(g string) int {
	return getGeneInt(g, -18, -16, HANDS_GENES_COUNT)
}

func getHandsGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getHandsGene(g)
	return StringAttribute{
		TraitType: 	"Hands",
		Value:		configService.Hands[gene],
	}
}

func getHandsGenePath(g string) string {
	gene := getHandsGene(g)
	return Gene(gene).toPath()
}

func getFrontItemGene(g string) string {
	return getGeneInt(g, -14, -12, FRONTITEM_GENES_COUNT)
}

func getFrontItemGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getFrontItemGene(g)
	return StringAttribute{
		TraitType:	"FrontItem"
		Value:		configService.FrontItem[gene],
	}
}

func getFrontItemGenePath(g string) string {
	gene := getFrontItemGene(g)
	return Gene(gene).toPath()
}

func getEyewearGene(g string) int {
	return getGeneInt(g, -12, -10, EYEWEAR_GENES_COUNT)
}

func getEyewearGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getEyewearGene(g)
	return StringAttribute{
		TraitType: "Eyewear",
		Value:     configService.Eyewear[gene],
	}
}

func getEyewearGenePath(g string) string {
	gene := getEyewearGene(g)
	return Gene(gene).toPath()
}

func getBackgroundGene(g string) int {
	return getGeneInt(g, -4, -2, BACKGROUND_GENES_COUNT)
}

func getBackgroundGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBackgroundGene(g)
	return StringAttribute{
		TraitType: "Background",
		Value:     configService.Background[gene],
	}
}

func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}

func getBaseGene(g string) int {
	return getGeneInt(g, -2, 0, BASE_GENES_COUNT)
}

func getBaseGeneAttribute(g string, configService *config.ConfigService) StringAttribute {
	gene := getBaseGene(g)
	return StringAttribute{
		TraitType: "Character",
		Value:     configService.Character[gene],
	}
}

func getBaseGenePath(g string) string {
	gene := getBaseGene(g)
	return Gene(gene).toPath()
}

func (g *Genome) name(configService *config.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getBaseGene(gStr)
	return fmt.Sprintf("%v #%v", configService.Character[gene], tokenId)
}

func (g *Genome) description(configService *config.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getBaseGene(gStr)
	return fmt.Sprintf("The %v named %v #%v is a Male Percy Penguin from the Snow Zone Universe, Scramble your Penguins attributes at anytime,",
	configService.Type[gene], configeService.Character[gene],tokenId)
}

func (g *Genome) gene() []string {
	gStr := string(*g)

	res := make([]string, 0, GENES_COUNT)

	res = append(res, getHandsGenePath(gStr))
	res = append(res, getFrontItemGenePath(gStr))
	res = append(res, getEyewearGeneAttribute(gStr))
	res = append(res, getBaseGenePath(gStr))
	res = append(res, getBackgroundGenePath(gStr))

	return res
}

func getRankAttribute(rank int) IntegerAttribute {
	return IntegerAttribute{
		TraitType:	"Rank",
		Value:		rank,
	}
}

func (g *Genome) attributes(configService *config.ConfigService, rarityResponse structs.RarityServiceResponse) []interface{} {
	gStr := string(*g)

	res := []interface{}{}
	res = append(res, getBaseGeneAttribute(gStr, configService))
	res = append(res, getHandsGeneAttribute(gStr, configService))
	res = append(res, getFrontItemGeneAttribute(gStr, configService))
	res = append(res, getEyewearGeneAttribute(gStr, configService))
	res = append(res, getBackgroundGeneAttribute(gStr, configService))
	res = append(res, getRarityScoreAttribute(rarityResponse.RarityScore))
	res = append(res, getRankAttribute(rarityResponse.Rank))
	
	return res
}

func (g *Genome) Metadata(tokenId string, configService *config.ConfigService, rarityResponse structs.RarityServiceResponse) Metadata {
	var m metadata
	genes := g.genes()

	m.Attributes = g.attributes(configService, rarityResponse)
	m.Name = g.name(configService, tokenId)
	m.Description = g.description(configService, tokenId)
	m.ExternalUrl = fmt.Sprintf("%s%s", EXTERNAL_URL, tokenId)

	b := strings.Builder{}

	b.WriteString(PERCYPENGUIN_IMAGE_URL)//Start with base URL

	for _, gene := range genes {
		b.WriteString(gene)
	}

	b.WriteString(".jpg") // Finish with jpg

	imageURL := b.String()

	imageExists := imageExists(imageURL)

	if !imageExists {
		generateAndSaveImage(genes)
	}

	m.Image = imageURL

}

type Metadata struct {
	Description	string		`json:"description"`
	Name 		string		`json:"name"`
	Image 		string		`json:"image"`
	Attributes	interface	`json:"attributes"`
	ExternalUrl	String		`json:external_url`
}