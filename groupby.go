package go_funcs

import (
	"encoding/json"
)

type (
	Group struct {
		data  []map[string]interface{}
		group map[string][]interface{}
	}
)

func NewGroup(data interface{}) *Group {
	b, _ := json.Marshal(data)
	var d []map[string]interface{}
	_ = json.Unmarshal(b, &d)

	return &Group{
		data: d,
	}
}

func (g *Group) GroupBy(key string) *Group {
	if len(g.group) == 0 {
		g.group = make(map[string][]interface{}, 0)
	}

	for _, d := range g.data {
		for k, v := range d {
			if k != key {
				continue
			}
			inx, ok := v.(string)
			if !ok {
				panic("key is not a string")
			}
			g.group[inx] = append(g.group[inx], d)
		}
	}
	return g
}

func (g *Group) Collect() map[string][]interface{} {
	return g.group
}
