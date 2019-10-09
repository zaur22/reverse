package main

import (
	"server/serv"
	"server/router"
	"server/util"
)

func main() {
	var utilServ = util.NewUtilService()
	var router = router.NewRouter(utilServ)

	serv.NewServer(router).Serve()
}