package main

import (
	"log"
	"plugin"
)

func loadPlugin() func() {
	plugin, err := plugin.Open("plugins/simple-plugin.so")
	if err != nil {
		log.Fatal(err)
	}

	simplePluginFunc, err := plugin.Lookup("SimplePluginFunc")
	if err != nil {
		log.Fatal(err)
	}

	f, ok := simplePluginFunc.(func())
	if !ok {
		log.Fatal("Tipo nao suportado")
	}
	return f
}

func main() {
	simplePlugin := loadPlugin()
	simplePlugin()
}
