package main

import (
	"fmt"
	"html"
	"html/template"
)

// isEthAddr reports whether s is a 42-char 0x-prefixed hex string (Ethereum address).
func isEthAddr(s string) bool {
	if len(s) != 42 {
		return false
	}
	if s[0] != '0' || (s[1] != 'x' && s[1] != 'X') {
		return false
	}
	for i := 2; i < 42; i++ {
		c := s[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// ethAddrLink renders a link to /black/cosmicgame/user/info/{addr} only for valid Ethereum addresses;
// otherwise it renders escaped plain text (e.g. aggregate prize labels like "(All CS NFT Stakers)").
func ethAddrLink(addr string) template.HTML {
	if !isEthAddr(addr) {
		return template.HTML(html.EscapeString(addr))
	}
	esc := html.EscapeString(addr)
	return template.HTML(fmt.Sprintf(`<a href="/black/cosmicgame/user/info/%s">%s</a>`, esc, esc))
}

// ethAddrLinkTo renders a link to pathPrefix+addr when addr is a valid Ethereum address; otherwise escaped text.
func ethAddrLinkTo(pathPrefix, addr string) template.HTML {
	if !isEthAddr(addr) {
		return template.HTML(html.EscapeString(addr))
	}
	esc := html.EscapeString(addr)
	return template.HTML(fmt.Sprintf(`<a href="%s%s">%s</a>`, pathPrefix, esc, esc))
}
