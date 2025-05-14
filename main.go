package main

import (
	"github.com/dirien/pulumi-talos-go-component/pkg/talos"
	"github.com/pulumi/pulumi-go-provider/infer"
)

func main() {
	err := infer.NewProviderBuilder().
		WithName("talos-go-component").
		WithNamespace("ediri").
		WithComponents(
			infer.Component(talos.NewTalosCluster),
		).
		BuildAndRun()

	if err != nil {
		panic(err)
	}
}
