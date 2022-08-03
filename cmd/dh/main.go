package main

import (
	"github.com/ekholme/dinner-helper/frstr"
	"github.com/ekholme/dinner-helper/srvr"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func init() {
	r.LoadHTMLGlob("srvr/html/templates/*.html")
}

func main() {
	ms := frstr.NewMealService()
	mh := srvr.NewMealHandler(ms)
	s := srvr.NewServer(r, mh)

	s.Run()


}
