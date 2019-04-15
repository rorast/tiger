package main

import (
	"github.com/kataras/iris"
)


func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello world! -- from iris")
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))

}
