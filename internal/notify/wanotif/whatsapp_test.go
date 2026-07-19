package wanotif

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestParseHTTPError(t *testing.T) {
	cases := []struct {
		name string
		body string
		want string
	}{
		{"well-formed", `{"error":{"message":"token expired"}}`, "token expired"},
		{"not json", `<html>boom</html>`, "unparsed error message"},
		{"missing error key", `{"other":{}}`, "%!s(<nil>)"},
		{"empty body", ``, "unparsed error message"},
		{"null", `null`, "%!s(<nil>)"},
	}
	for _, tc := range cases {
		err := parseHTTPError(strings.NewReader(tc.body))
		if err == nil {
			t.Errorf("%s: parseHTTPError returned nil error", tc.name)
			continue
		}
		if err.Error() != tc.want {
			t.Errorf("%s: parseHTTPError = %q, want %q", tc.name, err.Error(), tc.want)
		}
	}
}

func TestCreateSendWithTemplateRequest(t *testing.T) {
	wa := NewWhatsapp("token", "12345")
	req := wa.createSendWithTemplateRequest("15550001111", "alert", LanguageEnglish, []Components{
		{Type: "body", Parameters: wa.GenerateTemplateParameters("", "hello")},
	})
	if req.MessagingProduct != "whatsapp" || req.To != "15550001111" || req.Type != "template" {
		t.Fatalf("unexpected request envelope: %+v", req)
	}
	if req.Template.Name != "alert" || len(req.Template.Components) != 1 {
		t.Fatalf("unexpected template: %+v", req.Template)
	}
	if p := req.Template.Components[0].Parameters; len(p) != 1 || p[0].Type != "text" || p[0].Text != "hello" {
		t.Fatalf("unexpected parameters: %+v", p)
	}
}

// newStubWhatsapp returns a client pointed at a stub Graph API server that
// records the request and answers with the scripted status/body.
func newStubWhatsapp(t *testing.T, status int, respBody string) (*Whatsapp, *http.Request, *[]byte) {
	t.Helper()
	var gotReq http.Request
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotReq = *r
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(r.Body)
		gotBody = buf.Bytes()
		w.WriteHeader(status)
		_, _ = w.Write([]byte(respBody))
	}))
	t.Cleanup(srv.Close)

	wa := NewWhatsapp("secret-token", "12345")
	wa.BaseURL = srv.URL
	wa.HTTPClient = srv.Client()
	return wa, &gotReq, &gotBody
}

func TestSendText(t *testing.T) {
	wa, gotReq, gotBody := newStubWhatsapp(t, http.StatusOK,
		`{"messaging_product":"whatsapp","messages":[{"id":"wamid.X"}]}`)

	res, err := wa.SendText("15550001111", "server down")
	if err != nil {
		t.Fatalf("SendText: %v", err)
	}
	if res["messaging_product"] != "whatsapp" {
		t.Errorf("response not decoded: %v", res)
	}

	if gotReq.URL.Path != "/v14.0/12345/messages" {
		t.Errorf("path = %q, want /v14.0/12345/messages", gotReq.URL.Path)
	}
	if auth := gotReq.Header.Get("Authorization"); auth != "Bearer secret-token" {
		t.Errorf("Authorization = %q", auth)
	}
	if ct := gotReq.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q", ct)
	}
	var payload map[string]any
	if err := json.Unmarshal(*gotBody, &payload); err != nil {
		t.Fatalf("request body is not JSON: %v", err)
	}
	if payload["to"] != "15550001111" || payload["type"] != "text" {
		t.Errorf("payload envelope = %v", payload)
	}
	if text, ok := payload["text"].(map[string]any); !ok || text["body"] != "server down" {
		t.Errorf("payload text = %v", payload["text"])
	}
}

func TestSendWithTemplate(t *testing.T) {
	wa, _, gotBody := newStubWhatsapp(t, http.StatusOK, `{"messages":[{"id":"wamid.Y"}]}`)

	_, err := wa.SendWithTemplate("15550001111", "alert", []Components{
		{Type: "body", Parameters: wa.GenerateTemplateParameters("", "disk full")},
	})
	if err != nil {
		t.Fatalf("SendWithTemplate: %v", err)
	}
	var back SendWithTemplateRequest
	if err := json.Unmarshal(*gotBody, &back); err != nil {
		t.Fatalf("request body is not the template payload: %v", err)
	}
	if back.Template.Name != "alert" || back.Template.Language.Code != "en" {
		t.Errorf("template payload = %+v", back.Template)
	}
}

func TestSendTextGraphAPIError(t *testing.T) {
	wa, _, _ := newStubWhatsapp(t, http.StatusUnauthorized,
		`{"error":{"message":"Invalid OAuth access token"}}`)

	_, err := wa.SendText("15550001111", "x")
	if err == nil || err.Error() != "Invalid OAuth access token" {
		t.Fatalf("want the Graph API error message, got %v", err)
	}
}

func TestSendTextMalformedSuccessBody(t *testing.T) {
	wa, _, _ := newStubWhatsapp(t, http.StatusOK, `{not json`)
	if _, err := wa.SendText("1", "x"); err == nil {
		t.Fatal("want JSON decode error for malformed 200 body")
	}
}

func TestSendTextTransportError(t *testing.T) {
	wa := NewWhatsapp("tok", "42")
	wa.BaseURL = "http://127.0.0.1:1"
	if _, err := wa.SendText("1", "x"); err == nil {
		t.Fatal("want transport error for unreachable endpoint")
	}
}

func TestSendMessageUnmarshalableRequest(t *testing.T) {
	wa := NewWhatsapp("tok", "42")
	if _, err := wa.sendMessage(func() {}); err == nil {
		t.Fatal("want marshal error for unmarshalable request payload")
	}
}

func FuzzWhatsappPayload(f *testing.F) {
	f.Add(`{"error":{"message":"x"}}`, "15550001111", "alert", "hello")
	f.Add(``, "", "", "")
	f.Add(`{"error":[1,2]}`, "+1 (555) 000-1111", "name with spaces", "√"+string(rune(0)))
	f.Fuzz(func(t *testing.T, errBody, to, templateName, paramText string) {
		// parseHTTPError must always yield a non-nil error and never panic.
		if err := parseHTTPError(strings.NewReader(errBody)); err == nil {
			t.Fatalf("parseHTTPError(%q) returned nil error", errBody)
		}

		// The template payload builder must produce valid JSON that carries
		// the inputs through unchanged.
		wa := NewWhatsapp("tok", "42")
		req := wa.createSendWithTemplateRequest(to, templateName, wa.Language, []Components{
			{Type: "body", Parameters: wa.GenerateTemplateParameters("text", paramText)},
		})
		raw, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("marshal payload: %v", err)
		}
		var back SendWithTemplateRequest
		if err := json.NewDecoder(bytes.NewReader(raw)).Decode(&back); err != nil {
			t.Fatalf("payload does not round-trip as JSON: %v", err)
		}
		// JSON strings cannot carry invalid UTF-8 verbatim; the marshaller
		// replaces such bytes, so compare only when the input was valid UTF-8.
		if utf8.ValidString(to) && back.To != to {
			t.Fatalf("payload To changed: %q -> %q", to, back.To)
		}
	})
}
