package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	meta "services/Common/Meta"
	model "services/Model"
)

func CricketController(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		metaval := new(meta.Cricketmeta)
		res, err := ioutil.ReadAll(req.Body)
		fmt.Println(res)
		if err != nil {
			fmt.Println("Error in reading request body:", err.Error())
			return
		}
		defer req.Body.Close()
		err = json.Unmarshal([]byte(res), &metaval)
		if err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, "Bad Request", 400)
			return
		}
		fmt.Println(metaval)
		model.Cricketlog(metaval, req.RemoteAddr)

	}

}

func UserController(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		metaval := new(meta.Usermeta)
		res, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error in reading request body", err.Error())
			return
		}
		defer req.Body.Close()
		err = json.Unmarshal([]byte(res), &metaval)
		if err != nil {
			fmt.Println("Error:", err.Error())
			http.Error(w, "Bad Request", 400)
			return
		}
		model.Userlog(metaval, req.RemoteAddr)

	}

}
