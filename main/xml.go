package main

import (
	"encoding/xml"
)

type Node struct {
	XMLName xml.Name `xml:"xml"`
	Name string
	Value string
}

func main() {
	x := Node{Name: "12321", Value: "zzzzzzzzzz"}
	//var b bytes.Buffer
	//enc := xml.NewEncoder(&b)
	xml.MarshalIndent(x, "    ", "          ")
}




