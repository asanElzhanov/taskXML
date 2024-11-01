package entities

import (
	"encoding/xml"
)

type RequestXML struct {
	XMLName xml.Name `xml:"request"`
	Header  struct {
		Login    string `xml:"login"`
		Password string `xml:"password"`
		Hash     string `xml:"hash,omitempty"`
	} `xml:"header"`
	Body struct {
		Message string `xml:"message"`
	} `xml:"body"`
}

type ResponseXML struct {
	XMLName xml.Name `xml:"response"`
	Header  struct {
		Hash string `xml:"hash,omitempty"`
	} `xml:"header"`
	Body struct {
		Message string `xml:"message"`
		Info    string `xml:"info,omitempty"`
	} `xml:"body"`
}
