package main

import "github.com/pengdafu/sim/connectivity/router"

func main()  {
	r := router.Router()
	_ = r.Run(":8080")
}
