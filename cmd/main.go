package main

import (
	"main.go/api"
	"main.go/dao"
)

func main() {
	dao.InitDatabase()
	api.InitRouter()
}
