package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Site Armando Couto", "http://armandocouto.com.br", 800, 600, true)
}