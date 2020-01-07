package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", index)

	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>Index Page</h1>")
}
