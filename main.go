package main

import (
	"encoding/json"
	"fmt"
	"github.com/aimook/golibs/pkg/api"
)

func main() {
	r := api.ResultCreate().WithCode(200).WithMessage("message").WithData("data").WithSuccess()

	s, _ := json.Marshal(r)
	fmt.Printf("%s", s)

}
