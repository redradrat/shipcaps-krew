package main

import (
	"github.com/redradrat/shipcaps-kubectl/cmd/shipcaps/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // required for GKE
)

func main() {
	cli.InitAndExecute()
}
