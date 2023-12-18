package main

import (
	"WebBack/Router"
	"WebBack/dao"
)

func main() {
	dao.InitDataBase()
	r := Router.InitUserRouter()
	r.Run()

}
