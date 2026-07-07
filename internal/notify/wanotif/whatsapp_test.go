package wanotif

import (
	"bytes"
	"encoding/json"
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
