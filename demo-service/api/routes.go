/*
 * @Author: chuyuan.xu
 * @Date: 2020-03-09 16:28:57
 * @LastEditTime: 2020-03-09 17:14:17
 * @FilePath: /laracom/demo-service/api/routes.go
 */
package api

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"rpc_name",
		"GET",
		"/hello",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
			dict := map[string]string{
				"message": "你好 go rpc!",
			}
			data, _ := json.Marshal(dict)
			writer.Write(data)
		},
	},
}
