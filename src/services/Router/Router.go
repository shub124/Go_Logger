package router

import (
	"net/http"
	controller "services/Controller"
)

var Routemapping = map[string]func(_ http.ResponseWriter, req *http.Request){
	"user":    controller.UserController,
	"cricket": controller.CricketController,
}
