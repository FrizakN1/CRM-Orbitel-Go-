package main

import (
	"CRM-test/database"
	"CRM-test/router"
	"CRM-test/settings"
)

func main() {
	config := settings.Load("settings/settings.json")

	database.Connection(config)

	_ = router.Initialization().Run("127.0.0.1:8080")
}
