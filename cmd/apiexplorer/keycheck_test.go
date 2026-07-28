package main

// TestTemplateKeysAgainstLiveAPI executes every page template against the
// live API server response with missingkey=error, so a template referencing
// a key the API does not return fails loudly instead of rendering an empty
// table. It needs a running API server with indexed data, so it only runs
// when KEYCHECK=1 is set:
//
//	KEYCHECK=1 API_BASE=http://127.0.0.1:9090 go test -run TemplateKeys -v

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

// keycheckParams substitutes production-plausible values for path
// placeholders. RandomWalk and CosmicGame use different token/user id
// spaces, so a few names are resolved per project.
var keycheckParams = map[string]string{
	"offset":        "0",
	"limit":         "5",
	"sort":          "0",
	"order_by":      "0",
	"round_num":     "0",
	"prize_num":     "0",
	"init_ts":       "0",
	"fin_ts":        "2000000000",
	"interval_secs": "86400",
	"evtlog_start":  "1",
	"evtlog_end":    "100",
	"name":          "a",
	"record_id":     "1",
	"action_id":     "1",
	"deposit_id":    "1",
	"id":            "1",
}

var keycheckParamsCG = map[string]string{
	"user_addr": "0x1Ec14aDaf61e27AB339bc590BA4Bf2356Dd7E990",
	"user_aid":  "1028",
	"token_id":  "23",
	"evtlog_id": "25308",
}

var keycheckParamsRW = map[string]string{
	"user_addr":  "0x7BBF44394a23504cbE46b2b2d76929451cb86975",
	"user_aid":   "45",
	"token_id":   "17",
	"token_addr": "0x895a6F444BE4ba9d124F61DF736605792B35D66b",
}

func keycheckParam(api, name string) string {
	proj := keycheckParamsCG
	if strings.Contains(api, "/randomwalk/") {
		proj = keycheckParamsRW
	}
	if val, ok := proj[name]; ok {
		return val
	}
	if val, ok := keycheckParams[name]; ok {
		return val
	}
	return "1"
}

func keycheckSub(api string) string {
	for _, name := range paramNames(api) {
		api = strings.ReplaceAll(api, "{"+name+"}", keycheckParam(api, name))
	}
	return api
}

func TestTemplateKeysAgainstLiveAPI(t *testing.T) {
	if os.Getenv("KEYCHECK") != "1" {
		t.Skip("set KEYCHECK=1 (and API_BASE) to run against a live API server")
	}
	apiBase := strings.TrimSuffix(os.Getenv("API_BASE"), "/")
	if apiBase == "" {
		apiBase = "http://127.0.0.1:9090"
	}

	tmpl, err := loadTemplates()
	if err != nil {
		t.Fatalf("parsing templates: %v", err)
	}
	tmpl.Option("missingkey=error")

	s := &server{
		apiBase:   apiBase,
		templates: tmpl,
		client:    &http.Client{Timeout: apiTimeout},
	}

	all := append(append([]page{}, pages...), generatedPages...)
	seen := map[string]bool{}
	for _, p := range all {
		if p.API == "" {
			continue
		}
		key := p.Template + "|" + p.API
		if seen[key] {
			continue
		}
		seen[key] = true

		apiPath := keycheckSub(p.API)
		ctx, cancel := context.WithTimeout(context.Background(), apiTimeout)
		data, err := s.fetch(ctx, apiPath)
		cancel()
		if err != nil {
			t.Logf("SKIP %s: fetching %s: %v", p.Template, apiPath, err)
			continue
		}
		injectParams(data, p.API, func(name string) string { return keycheckParam(p.API, name) })
		// The dashboard only carries V3Config on V3 deployments; give the
		// template the guard value so the rest of the page is still checked.
		if p.Template == "cosmicsignature/cg_index.html" {
			if _, ok := data["V3Config"]; !ok {
				data["V3Config"] = map[string]any{"IsV3": false}
			}
		}
		if err := tmpl.ExecuteTemplate(io.Discard, p.Template, data); err != nil {
			t.Errorf("%s (api %s): %v", p.Template, p.API, err)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
