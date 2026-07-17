package config_test

import (
	"fmt"

	"github.com/PredictionExplorer/augur-explorer/internal/config"
)

// ExampleLoad shows the tag-driven loader on a service-style struct: values
// come from the injected getenv, empty variables fall back to `default`
// tags, and every problem — including missing `required` variables — is
// aggregated into one error so a broken deployment reports everything at
// once.
func ExampleLoad() {
	type ServiceConfig struct {
		HTTPPort string `env:"HTTP_PORT" default:"8080"`
		RPCURL   string `env:"RPC_URL" required:"true" secret:"url"`
		Debug    bool   `env:"DEBUG" default:"false"`
	}

	env := map[string]string{
		"RPC_URL": "wss://arb-node.example/ws",
		"DEBUG":   "yes",
	}
	var cfg ServiceConfig
	if err := config.Load(&cfg, func(name string) string { return env[name] }); err != nil {
		fmt.Println("load failed:", err)
		return
	}
	fmt.Println(cfg.HTTPPort, cfg.Debug)

	// An empty environment misses the required variable; the error carries
	// one line per problem.
	err := config.Load(&ServiceConfig{}, func(string) string { return "" })
	fmt.Println(err)
	// Output:
	// 8080 true
	// configuration errors:
	//   RPC_URL: required but not set
}
