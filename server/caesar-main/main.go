package main

import (
  "github.com/kataras/iris/v12"
)

func main() {
  app := iris.Default()
  app.Logger().SetLevel("debug")
  app.Get("/", func (ctx iris.Context)  {
    ctx.JSON(iris.Map{
      "Msg": "Hi Caesar!",
    })
  })
  app.Listen(":8080")
}
