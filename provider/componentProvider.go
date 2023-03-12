package provider

import "log"

type ComponentProvider struct {
	ComponentMap map[string]interface{}
}

func (b *ComponentProvider) Required(name string) interface{} {
	component, exist := b.ComponentMap[name]

	if !exist {
		log.Fatalf("Required %s but does not exist\n", name)
	}

	return component
}
