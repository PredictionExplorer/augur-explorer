// Package wanotif sends WhatsApp notifications via the WhatsApp Cloud API.
//
// Copied from: github.com/febriliankr/whatsapp-cloud-api
// (had to copy because the library has a bug (issue #2 at Github) which I fixed in my copy)
package wanotif

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Whatsapp is a WhatsApp Cloud API client bound to one sending phone
// number.
type Whatsapp struct {
	Token         string
	APIVersion    string
	PhoneNumberID string
	Language      TemplateLanguage

	// BaseURL is the Graph API root; empty selects the production
	// https://graph.facebook.com (tests point it at an httptest server).
	BaseURL string
	// HTTPClient issues the requests; nil selects http.DefaultClient.
	HTTPClient *http.Client
}

// TemplateLanguage selects the language of a WhatsApp message template.
type TemplateLanguage struct {
	Code string `json:"code,omitempty"`
}

// Components for WhatsApp template messages.
type Components struct {
	Type       string               `json:"type,omitempty"`
	Parameters []TemplateParameters `json:"parameters,omitempty"`
}

// TemplateParameters for WhatsApp template messages.
type TemplateParameters struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

// Template for WhatsApp template messages.
type Template struct {
	Name       string           `json:"name,omitempty"`
	Language   TemplateLanguage `json:"language,omitempty"`
	Components []Components     `json:"components,omitempty"`
}

// SendWithTemplateRequest for WhatsApp API.
type SendWithTemplateRequest struct {
	MessagingProduct string   `json:"messaging_product,omitempty"`
	To               string   `json:"to,omitempty"`
	Type             string   `json:"type,omitempty"`
	Template         Template `json:"template,omitempty"`
}

// LanguageEnglish is the default template language.
var LanguageEnglish = TemplateLanguage{
	Code: "en",
}

// NewWhatsapp returns a client for the given bearer token and sending phone
// number, defaulting to English templates and Graph API v14.0.
func NewWhatsapp(token string, phoneNumberID string) *Whatsapp {
	return &Whatsapp{
		Language:      LanguageEnglish,
		Token:         token,
		APIVersion:    "v14.0",
		PhoneNumberID: phoneNumberID,
	}
}

func parseHTTPError(body io.Reader) (err error) {
	var errRes map[string]map[string]interface{}
	err = json.NewDecoder(body).Decode(&errRes)
	if err != nil {
		return fmt.Errorf("unparsed error message")
	}
	msg := fmt.Sprintf("%s", errRes["error"]["message"])
	return errors.New(msg)
}

func (wa *Whatsapp) sendMessage(request interface{}) (res map[string]interface{}, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(jsonRequest)

	baseURL := wa.BaseURL
	if baseURL == "" {
		baseURL = "https://graph.facebook.com"
	}
	endpoint := fmt.Sprintf("%s/%s/%s/messages", baseURL, wa.APIVersion, wa.PhoneNumberID)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, endpoint, body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Authorization", "Bearer "+wa.Token)
	req.Header.Set("Content-Type", "application/json")

	client := wa.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	if resp.StatusCode != http.StatusOK {
		return res, parseHTTPError(resp.Body)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

// SendWithTemplate sends the named message template with the given
// components to toPhoneNumber.
func (wa *Whatsapp) SendWithTemplate(toPhoneNumber string, templateName string, components []Components) (res map[string]interface{}, err error) {
	request := wa.createSendWithTemplateRequest(toPhoneNumber, templateName, wa.Language, components)

	return wa.sendMessage(request)
}

// SendText sends a plain-text message to toPhoneNumber.
//
// Note: the recipient must initiate a conversation with the sending account
// first, otherwise the message will be classified as spam and won't be
// delivered, though no error is reported.
func (wa *Whatsapp) SendText(toPhoneNumber string, text string) (res map[string]interface{}, err error) {
	request := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                toPhoneNumber,
		"type":              "text",
		"text": map[string]string{
			"body": text,
		},
	}
	return wa.sendMessage(request)
}

func (wa *Whatsapp) createSendWithTemplateRequest(receiverPhoneNumber string, templateName string, language TemplateLanguage, components []Components) (res SendWithTemplateRequest) {
	return SendWithTemplateRequest{
		MessagingProduct: "whatsapp",
		To:               receiverPhoneNumber,
		Type:             "template",
		Template: Template{
			Name:     templateName,
			Language: language,
			// Components can be empty if you don't want to send any parameters
			Components: components,
		},
	}
}

// GenerateTemplateParameters builds one template parameter per argument;
// an empty parameterType defaults to "text".
func (wa *Whatsapp) GenerateTemplateParameters(parameterType string, args ...string) (res []TemplateParameters) {
	if parameterType == "" {
		parameterType = "text"
	}

	for _, arg := range args {
		res = append(res, TemplateParameters{
			Type: parameterType,
			Text: arg,
		})
	}
	return
}
