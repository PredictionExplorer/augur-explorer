package v2

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// undeclaredLengthReader hides its size from httptest.NewRequest so the
// request carries no Content-Length and only the read-time cap can fire.
type undeclaredLengthReader struct{ r io.Reader }

func (u undeclaredLengthReader) Read(p []byte) (int, error) { return u.r.Read(p) }

// TestOversizedBodyProblemMatchesMiddleware proves the two 413 paths render
// byte-identical problems: the middleware's Content-Length pre-check (which
// answers without entering the generated code) and the generated JSON
// decoder failing on the MaxBytesReader wrapper (mapped by
// writeRequestError). Clients cannot tell where the cap fired.
func TestOversizedBodyProblemMatchesMiddleware(t *testing.T) {
	t.Parallel()
	const limit = 128
	router := httpx.NewRouter()
	router.Use(common.MaxRequestBody(limit))
	newTestServer(t, fakeBidReader{}).RegisterRoutes(router)

	oversized := `{"nft1":1,"nft2":2,"winner":1,"chainId":1,"nonce":"` +
		strings.Repeat("x", limit) + `","signature":"0xff"}`

	declared := httptest.NewRequest(http.MethodPost, "/api/v2/randomwalk/ranking/votes",
		strings.NewReader(oversized))
	if declared.ContentLength <= limit {
		t.Fatalf("test body must exceed the %d-byte cap, got %d", limit, declared.ContentLength)
	}
	declaredRec := httptest.NewRecorder()
	router.ServeHTTP(declaredRec, declared)

	undeclared := httptest.NewRequest(http.MethodPost, "/api/v2/randomwalk/ranking/votes",
		undeclaredLengthReader{strings.NewReader(oversized)})
	if undeclared.ContentLength != -1 {
		t.Fatalf("ContentLength = %d, want -1 (undeclared)", undeclared.ContentLength)
	}
	undeclaredRec := httptest.NewRecorder()
	router.ServeHTTP(undeclaredRec, undeclared)

	for name, rec := range map[string]*httptest.ResponseRecorder{
		"declared length (middleware pre-check)": declaredRec,
		"undeclared length (decoder read)":       undeclaredRec,
	} {
		assertProblem(t, rec, http.StatusRequestEntityTooLarge)
		var problem Problem
		decodeResponse(t, rec, &problem)
		if problem.Type != problemTypeBase+"request-too-large" ||
			problem.Detail == nil || *problem.Detail != common.RequestBodyTooLargeDetail(limit) ||
			problem.Instance == nil || *problem.Instance != "/api/v2/randomwalk/ranking/votes" {
			t.Fatalf("%s: problem = %+v", name, problem)
		}
	}
	if declaredRec.Body.String() != undeclaredRec.Body.String() {
		t.Fatalf("the two 413 paths must render byte-identical problems:\nmiddleware: %q\ndecoder:    %q",
			declaredRec.Body.String(), undeclaredRec.Body.String())
	}
}

// TestSpecDeclares413ExactlyOnBodyOperations pins the 413 inventory: every
// operation accepting a request body declares the shared RequestTooLarge
// response, and no body-less operation does (their global-middleware 413 is
// a transport concern documented in operations.md, like the global 429).
func TestSpecDeclares413ExactlyOnBodyOperations(t *testing.T) {
	t.Parallel()
	spec, err := GetSpec()
	if err != nil {
		t.Fatalf("GetSpec: %v", err)
	}
	bodyOperations := 0
	for path, item := range spec.Paths.Map() {
		for method, operation := range item.Operations() {
			hasBody := operation.RequestBody != nil
			declares413 := operation.Responses.Value("413") != nil
			if hasBody != declares413 {
				t.Errorf("%s %s: requestBody=%v but 413 declared=%v", method, path, hasBody, declares413)
			}
			if hasBody {
				bodyOperations++
			}
		}
	}
	if bodyOperations != 3 {
		t.Errorf("body-accepting operations = %d, want 3 (bid-ban create, ranking votes and matches)", bodyOperations)
	}
}
