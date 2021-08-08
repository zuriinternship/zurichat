package marketplace

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
	"zuri.chat/zccore/utils"
)

type Category struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Slug string             `json:"slug" bson:"slug"`
}

type Categories []Category
type Features []string

type App struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Description  string             `json:"description" bson:"description"`
	InstallCount int                `json:"install_count,omitempty" bson:"install___count,omitempty"`
	Categories   *Categories        `bson:",omitempty" bson:",omitempty"`
	Features     *Features          `bson:",omitempty" bson:",omitempty"`
	ApprovedAt   bool               `json:"approved_at,omitempty" bson:"approved___at,omitempty"`
}

func Apps(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	var apps []App

	response.Header().Set("Content-Type", "application/json")

	collection, err := utils.GetMongoDbCollection("myFirstDatabase", "apps")
	if err != nil {
		utils.GetError(err, response)
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		utils.GetError(err, response)
		return
	}

	defer cursor.Close(ctx)

	apps = make([]App, 0)
	for cursor.Next(ctx) {
		var app App
		cursor.Decode(&app)
		apps = append(apps, app)
	}

	json.NewEncoder(response).Encode(apps)
}
