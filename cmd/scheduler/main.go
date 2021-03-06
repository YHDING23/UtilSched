package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

//	"github.com/angao/scheduler-framework-sample/pkg/plugins"
    "github.com/YHDING23/UtilSched/pkg/plugins/register"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	cmd := plugins.Register()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}