package handlers

import (
	"context"
	"os"

	"github.com/PercyPenguin-Metadata/app/config"
	"github.com/PercyPenguin-Metadata/constants"
	"github.com/PercyPenguin-Metadata/db"
	"github.com/PercyPenguin-Metadata/structs"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetRarityById endpoints accepts id of a single polymorph and information for a single polymorph.
//
// If no polymorph is found returns empty response
func GetRarityById(id int) structs.RarityServiceResponse {
	godotenv.Load()
	percyPenguinDBName := os.Getenv("PERCYPENGUIN_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")
	collection, err := db.GetMongoDbCollection(percyPenguinDBName, rarityCollectionName)
	if err != nil {
		return stucts.RarityServiceResponse{}
	}

	fincOptions := options.FindOneOptions{}
	removePrivateFieldsSingle(&findOptions)

	var filter bson.M = bson.M{}
	filter = bson.M{constants.MorphFieldNames.TokenId: id}

	var result = structs.RarityServiceResponse{}
	curr := collection.FineOne(context.Background(), filter, &findOptions)

	curr.Decode(&result)

	return result
}

// removePrivateFieldsSingle removes internal fields that are of no interest to the users of the API.
//
// Configuration of these fields can be found in helpers.apiConfig.go
func removePrivateFieldsSingle(findOptions *options.FindOneOptions) {
	noProjectionFields := bson.M{}
	for _, field := range config.Morph_NO_PROJECTION_FIELDS {
		noProjectionFields[field] = 0
	}
	findOptions.SetProjection(noProjectionFields)
}
