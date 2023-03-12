package provider

import "log"

type ComponentProvider struct {
	componentMap map[string]interface{}
}

func NewComponentProvider() *ComponentProvider {
	return &ComponentProvider{
		componentMap: make(map[string]interface{}, 1),
	}
}

func (p *ComponentProvider) Add(name string, component interface{}) {
	p.componentMap[name] = component
}

func (p *ComponentProvider) Required(name string) interface{} {
	component, exist := p.componentMap[name]

	if !exist {
		log.Fatalf("Required %s but does not exist\n", name)
	}

	return component
}
