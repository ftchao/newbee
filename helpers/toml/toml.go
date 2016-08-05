package toml

import (
	"github.com/pelletier/go-toml"
)

type TomlTree struct {
	toml *toml.TomlTree
}

func LoadFile(path string) (t *TomlTree, err error) {

	tolmTree, err := toml.LoadFile(path)
	if err != nil {
		return nil, err
	}

	t = &TomlTree{
		toml: tolmTree,
	}

	return t, nil
}

func (t *TomlTree) Get(key string) interface{} {
	if t != nil {
		return t.toml.Get(key)
	}

	return nil
}

func (t *TomlTree) GetDefault(key string, def interface{}) interface{} {
	if t != nil {
		return t.toml.GetDefault(key, def)
	}

	return nil
}
