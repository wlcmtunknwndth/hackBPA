package main

import (
	"fmt"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%+v", cfg)
}
