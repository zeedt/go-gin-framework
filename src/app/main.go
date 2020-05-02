package main

func main() {
	route := registerRoute()
	route.Run(":3003")
}
