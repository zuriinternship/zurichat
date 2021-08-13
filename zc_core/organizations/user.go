package organizations

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"zuri.chat/zccore/utils"
)

func GetUserOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	organisation, err := utils.GetMongoDbCollection("organizations")
	if err != nil{
		utils.GetError(err, w)
	}

  organisationResult, err := organisation.InsertOne(context.Background(), bson.D{
		{"name": "Zuri" },
		{"description": "Zuri organization"},
	})
	if err != nil {
    log.Fatal(err)
	}
	fmt.Fprint(w, len(organisationResult.InsertedIDs))
}
