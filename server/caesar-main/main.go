package main

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/kataras/iris/v12"
	"github.com/ronbb/caesar/public/models/caesar"
	"github.com/ronbb/caesar/server/caesar-main/config"
	"github.com/ronbb/caesar/server/caesar-main/db"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	app.Post("/sign_in", func(ctx iris.Context) {
		res := caesar.Response{}
		res.Time = timestamppb.Now()
		b, err := ctx.GetBody()
		if err != nil {
			res.Code = caesar.ResponseCode_RC_FAILED
			res.Message = err.Error()
			ctx.Protobuf(&res)
			return
		}
		auth := caesar.Authentication{}
		err = proto.Unmarshal(b, &auth)
		if err != nil {
			res.Code = caesar.ResponseCode_RC_FAILED
			res.Message = err.Error()
			ctx.Protobuf(&res)
			return
		}
		userInfo, err := db.UserSignIn(&auth)
		if err != nil {
			res.Code = caesar.ResponseCode_RC_FAILED
			res.Message = err.Error()
			ctx.Protobuf(&res)
			return
		}
		res.Code = caesar.ResponseCode_RC_SUCCESSFUL
		res.Message = "Success"
		res.Result.MarshalFrom(userInfo)
		ctx.Protobuf(&res)
	})

	app.Post("/sign_up", func(ctx iris.Context) {
		res := caesar.Response{}
		b, err := ctx.GetBody()
		if err != nil {
			res.Code = caesar.ResponseCode_RC_FAILED
			res.Message = err.Error()
			ctx.Protobuf(&res)
			return
		}
		auth := caesar.UserRegistration{}
		err = proto.Unmarshal(b, &auth)
		if err != nil {
			res.Code = caesar.ResponseCode_RC_FAILED
			res.Message = err.Error()
			ctx.Protobuf(&res)
			return
		}
		err = db.UserSignUp(&auth)
		if err != nil {
			res.Code = caesar.ResponseCode_RC_FAILED
			res.Message = err.Error()
			ctx.Protobuf(&res)
			return
		}
		res.Code = caesar.ResponseCode_RC_SUCCESSFUL
		res.Message = "Success"
		ctx.Protobuf(&res)
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
