package hello

import (
	"encoding/json"
	ofctx "github.com/Mershab99/functions-framework-go/context"
	"github.com/Mershab99/functions-framework-go/functions"
)

func init() {
	functions.OpenFunction("Bar", bar,
		functions.WithFunctionPath("/bar/{name}"),
		functions.WithFunctionMethods("GET", "POST"),
	)
}

func bar(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	vars := ofctx.VarsFromCtx(ctx.GetNativeContext())
	response := map[string]string{
		string(in): vars["name"],
	}

	// Test DAPR Client
	ctx.GetDaprClient()
	responseBytes, _ := json.Marshal(response)
	return ctx.ReturnOnSuccess().WithData(responseBytes), nil
}
