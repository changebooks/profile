package profile

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Schema struct {
	name   string            // 配置名，如：database-w
	active string            // 当前环境，如：.../profile/dev/database-w.ini
	global string            // 全局，如：.../profile/database.ini
	data   map[string]string // 键 => 值
}

func (x *Schema) Get(key string) string {
	return x.data[key]
}

func (x *Schema) IsExist(key string) bool {
	if _, ok := x.data[key]; ok {
		return true
	} else {
		return false
	}
}

func (x *Schema) ToString() string {
	return fmt.Sprintf("name:   %v\n"+
		"active: %v\n"+
		"global: %v\n"+
		"data:   %v\n",
		x.GetName(), x.GetActive(), x.GetGlobal(), x.GetData(),
	)
}

func (x *Schema) GetName() string {
	return x.name
}

func (x *Schema) GetActive() string {
	return x.active
}

func (x *Schema) GetGlobal() string {
	return x.global
}

func (x *Schema) GetData() map[string]string {
	return x.data
}

type SchemaBuilder struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	name   string     // 配置名，如：database-w
	active string     // 当前环境，如：.../profile/dev/database-w.ini
	global string     // 全局，如：.../profile/database.ini
}

func (x *SchemaBuilder) Build() (*Schema, error) {
	if x.name == "" {
		return nil, errors.New("name can't be empty")
	}

	if x.active == "" && x.global == "" {
		return nil, errors.New("active or global can't be empty")
	}

	var active map[string]string
	var global map[string]string
	var err error

	if x.active != "" {
		if active, err = ParseFile(x.active); err != nil {
			return nil, err
		}
	}

	if x.global != "" {
		if global, err = ParseFile(x.global); err != nil {
			return nil, err
		}
	}

	data := Merge(global, active)

	return &Schema{
		name:   x.name,
		active: x.active,
		global: x.global,
		data:   data,
	}, nil
}

func (x *SchemaBuilder) SetName(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.name = s
	return x
}

func (x *SchemaBuilder) SetActive(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.active = s
	return x
}

func (x *SchemaBuilder) SetGlobal(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.global = s
	return x
}
