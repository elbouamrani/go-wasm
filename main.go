package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WASM!")
	document := js.Global().Get("document")

	paragraph := document.Call("createElement", "p")
	paragraph.Set("innerHTML", "Hello world from wasm, written in go")
	paragraph.Set("className", "block")

	style := document.Call("createElement", "style")
	style.Set("innerHTML", `
		.block {
			border: 1pt solid black;
			background-color: lightgray;
		}
	`)

	document.Get("head").Call("appendChild", style)
	document.Get("body").Call("appendChild", paragraph)

}
