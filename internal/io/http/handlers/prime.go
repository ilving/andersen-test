package handlers

import (
	"awesomeProject/internal/domain/number"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func FindPrimeNumbers(ctx *fasthttp.RequestCtx) {
	result, err := number.FindPrimeNumbers(ctx.PostBody())
	ctx.Response.SetStatusCode(200)
	if err != nil {
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(map[string]string{
			"Error": fmt.Sprintf("error on request processing. Err: %s", err.Error()),
		})
	} else {
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(result)
	}
}
