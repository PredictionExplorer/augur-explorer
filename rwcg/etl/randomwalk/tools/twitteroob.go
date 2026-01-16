// Copyright 2013 Gary Burd
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"encoding/json"
	"net/url"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	//"github.com/gomodule/oauth1/oauth"
)

var oauthClient = Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

var credPath = flag.String("config", "config.json", "Path to configuration file containing the application's credentials.")

func readCredentials() error {
	b, err := ioutil.ReadFile(*credPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &oauthClient.Credentials)
}

func main() {
	if err := readCredentials(); err != nil {
		log.Fatal(err)
	}
	oauthClient.APIKey="GfCPjysetCkHJhujwodYsOQyl"	// the client
	tempCred, err := oauthClient.RequestTemporaryCredentials(nil, "oob", nil)
	if err != nil {
		log.Fatal("RequestTemporaryCredentials:", err)
	}
	fmt.Printf("Credentials:\n")
	fmt.Printf("\tToken: %v\n",tempCred.Token)
	fmt.Printf("\tSecret: %v\n",tempCred.Secret)
	u := oauthClient.AuthorizationURL(tempCred, nil)

	fmt.Printf("1. Go to %s\n2. Authorize the application\n3. Enter verification code:\n", u)

	var code string
	fmt.Scanln(&code)

	tokenCred, _, err := oauthClient.RequestToken(nil, tempCred, code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token credentials:\n")
	fmt.Printf("\tToken: %v\n",tokenCred.Token)
	fmt.Printf("\tToken Secret: %v\n",tokenCred.Secret)

	form := url.Values{"status": {"got authorization from account owner"}}
	resp, err := oauthClient.Post(nil, tokenCred,
		"https://api.twitter.com/1.1/statuses/update.json", form)
	if err != nil {
		fmt.Printf("Get error: %v\n",err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("body = %+v\n",resp.Body)
	b, err := io.ReadAll(resp.Body)
	fmt.Printf("dump body:\n")
	fmt.Println(string(b))
/*
	fmt.Printf("Second request...\n\n")
	form = url.Values{"status": {"second test tweet"}}
	resp, err = oauthClient.Post(nil, tokenCred,
		"https://api.twitter.com/1.1/statuses/update.json", form)
	if err != nil {
		fmt.Printf("Get error: %v\n",err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("body = %+v\n",resp.Body)
	b, err = io.ReadAll(resp.Body)
	fmt.Printf("dump body:\n")
	fmt.Println(string(b))
*/
}
