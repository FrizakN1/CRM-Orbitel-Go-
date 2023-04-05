package settings

import (
	"CRM-test/structures"
	"CRM-test/utils"
	"encoding/json"
)

var settings structures.Setting

func Load(filename string) *structures.Setting {
	bytes, e := utils.LoadFile(filename)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}
	e = json.Unmarshal(bytes, &settings)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}
	return &settings
}
