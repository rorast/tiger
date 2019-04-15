package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	htmlEngine := iris.HTML("./iris/", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello world! -- from iris")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "測試頁面")
		ctx.ViewData("Content", "Hello world! -- iris")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))

}
