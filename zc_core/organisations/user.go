package organisations

import (
	"net/http"
	"zuri.chat/zccore/utils"
)

func GetUserOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection, err := utils.GetMongoDbCollection("", "")

	if err != nil{
		utils.GetError(err, w)
	}

	
}
