package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/ronbb/caesar/server/caesar-main/config"
)

func main() {
	app := iris.Default()

	if config.Global().Debug {
		app.Logger().SetLevel("debug")
  }
  
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"Msg": "Hi Caesar!",
		})
	})

	https := config.Global().HTTPS
	http := config.Global().HTTP

	if https.Enabled {
		app.Run(iris.AutoTLS(https.Address(), https.Domain, https.Email, iris.AutoTLSNoRedirect(fallbackServer)))
	} else {
		app.Listen(http.Address())
	}

}

// Without this, the ACME HTTP-01 challenge would fail
func fallbackServer(acme func(http.Handler) http.Handler) *http.Server {
	fallback := config.Global().HTTPS.Fallback
	srv := &http.Server{
		Addr: fallback.Address(),
		Handler: acme(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, fallback.Domain, iris.StatusTemporaryRedirect)
		})),
	}
	go srv.ListenAndServe()
	return srv
}
