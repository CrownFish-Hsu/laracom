/*
 * @Author: chuyuan.xu
 * @Date: 2020-03-09 16:29:50
 * @LastEditTime: 2020-03-09 16:29:51
 * @FilePath: /laracom/demo-service/api/route.go
 */
package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}