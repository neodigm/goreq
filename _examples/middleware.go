package main

import (
	"fmt"
	"github.com/zhshch2002/goreq"
)

func main() {
	// you can config `req.DefaultClient.Use()` to set global middleware
	c := req.NewClient()                                              // create a new client
	c.Use(req.WithRandomUA())                                         // Add a builtin middleware
	c.Use(func(client *req.Client, handler req.Handler) req.Handler { // Add another middleware
		return func(r *req.Request) *req.Response {
			fmt.Println("this is a middleware")
			r.Header.Set("req", "goreq")
			return handler(r)
		}
	})

	txt, err := req.Get("https://httpbin.org/get").SetClient(c).Do().Txt()
	fmt.Println(txt, err)
}
