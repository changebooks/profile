package profile

import (
	"errors"
	"fmt"
	"path"
	"strings"
	"sync"
)

type Profile struct {
	active    string             // 当前环境，如：dev
	directory string             // 配置目录，如：.../profile
	data      map[string]*Schema // name => *Schema
}

func (x *Profile) Get(name string) *Schema {
	return x.data[name]
}

func (x *Profile) IsExist(name string) bool {
	if _, ok := x.data[name]; ok {
		return true
	} else {
		return false
	}
}

func (x *Profile) ToString() string {
	data := ""
	for name, schema := range x.data {
		data += "==== " + name + " ====\n"
		data += schema.ToString()
	}

	return fmt.Sprintf("active:    %v\n"+
		"directory: %v\n"+
		"%v\n",
		x.GetActive(), x.GetDirectory(), data,
	)
}

func (x *Profile) GetActive() string {
	return x.active
}

func (x *Profile) GetDirectory() string {
	return x.directory
}

func (x *Profile) GetData() map[string]*Schema {
	return x.data
}

type Builder struct {
	mu        sync.Mutex // ensures atomic writes; protects the following fields
	active    string     // 当前环境，如：dev
	directory string     // 配置目录，如：.../profile
	names     []string   // 配置名，如：database-w、database-r、database-b
	global    string     // 全局，如：database
}

func (x *Builder) Build() (*Profile, error) {
	if x.active == "" {
		return nil, errors.New("active can't be empty")
	}

	if x.directory == "" {
		return nil, errors.New("directory can't be empty")
	}

	var global string
	if x.global != "" {
		global = path.Join(x.directory, x.global) + FileExt
	}

	data := make(map[string]*Schema)

	if len(x.names) > 0 {
		for _, n := range x.names {
			if n == "" {
				return nil, errors.New("name can't be empty")
			}

			active := path.Join(x.directory, x.active, n) + FileExt

			b := &SchemaBuilder{}
			schema, err := b.SetName(n).SetActive(active).SetGlobal(global).Build()
			if err != nil {
				return nil, err
			}

			data[n] = schema
		}
	} else {
		if x.global == "" {
			return nil, errors.New("global can't be empty")
		}

		b := &SchemaBuilder{}
		schema, err := b.SetName(x.global).SetGlobal(global).Build()
		if err != nil {
			return nil, err
		}

		data[x.global] = schema
	}

	return &Profile{
		active:    x.active,
		directory: x.directory,
		data:      data,
	}, nil
}

func (x *Builder) SetActive(s string) *Builder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.active = s
	return x
}

func (x *Builder) SetDirectory(s string) *Builder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.directory = s
	return x
}

func (x *Builder) AddName(s string) error {
	if s = strings.TrimSpace(s); s == "" {
		return errors.New("name can't be empty")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.names = append(x.names, s)
	return nil
}

func (x *Builder) SetGlobal(s string) *Builder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.global = s
	return x
}
