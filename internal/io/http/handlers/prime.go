package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"

	"awesomeProject/internal/domain/number"
)

func FindPrimeNumbers(ctx *fasthttp.RequestCtx) {
	result, err := number.FindPrimeNumbers(ctx.PostBody())
	if err != nil {
		ctx.Response.SetStatusCode(400)

		json.NewEncoder(ctx.Response.BodyWriter()).Encode(map[string]string{
			"Error": fmt.Sprintf("error on request processing. Err: %s", err.Error()),
		})
		return
	}

	ctx.Response.SetStatusCode(200)
	json.NewEncoder(ctx.Response.BodyWriter()).Encode(result)
}
