package config

type QueryHandlers map[string]*QueryHandler

type QueryHandler struct {
	Name     string   `yaml:"name"`
	OnEvents OnEvents `yaml:"onEvents"`
}

type OnEvents map[string]*OnEvent

type OnEvent struct {
	Name string `yaml:"name"`
}

func (q *QueryHandlers) init() {
	if q == nil {
		return
	}
	for name, item := range *q {
		if item != nil {
			item.init(name)
		}
	}
}

func (q *QueryHandler) init(name string) {
	q.Name = name
	q.OnEvents.init()
}

func (e *OnEvents) init() {
	for name, item := range *e {
		if item != nil {
			item.init(name)
		}
	}
}

func (q *OnEvent) init(name string) {
	q.Name = name
}
