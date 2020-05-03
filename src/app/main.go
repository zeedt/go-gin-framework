package main

func main() {
	db :=setUpDb()
	route := registerRoute(db)
	route.Run(":3003")
}
