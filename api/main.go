package main

import (
	database_middleware "api/middleware"
	"api/routes"
	"api/utils"
	_ "github.com/lib/pq"
)

func main() {
	db, err := utils.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := routes.Router(db)
	router.Use(database_middleware.DatabaseMiddleware(db))
	_ = router.Run()
}
