package organisations

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"zuri.chat/zccore/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	db, collection := "zurichat", "organizations"

	// validate required fields
	// add required params into required array and make an empty array to hold error strings
	required, empty := []string{"user_id"}, make([]string, 0)
	// get the form params
	form_params := r.URL.Query()
	// loop through and check for empty required params
	for _, value := range required {
		if strings.TrimSpace(form_params.Get(value)) == "" {
			empty = append(empty, strings.Join(strings.Split(value, "_"), " "))
		}

		if len(empty) > 0 {
			utils.GetError(errors.New(strings.Join(empty, ", ") + " required"), w)
		}
	}

	// check if organization name is already taken
	org, err := utils.GetMongoDbCollection(db, collection)
	if err != nil {
		utils.GetError(err, w)
	}
	if org != nil {
		utils.GetError(errors.New("Organization name is already taken"), w)
	}

	// confirm if user_id exists
	user, err := utils.GetMongoDbCollection(db, "users", form_params.Get("user_id"))
	if err != nil {
		utils.GetError(err, w)
	}
	if user == nil {
		utils.GetError(errors.New("Invalid user id"), w)
	}
	// save organization
	save, err := utils.CreateMongoDbCollection(db, collection)

	if err != nil {
		utils.GetError(err, w)
	}

}
