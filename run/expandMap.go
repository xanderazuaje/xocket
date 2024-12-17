package run

import (
	"github.com/xanderazuaje/xocket/parsing"
	"github.com/xanderazuaje/xocket/random"
	"os"
)

func expandMap(m map[string][]string) (map[string][]string, error) {
	for k, v := range m {
		k2, err := formatFilter(k)
		if err != nil {
			return nil, err
		}
		k2 = os.ExpandEnv(k2)
		for i, v2 := range v {
			v2, err := formatFilter(v2)
			if err != nil {
				return nil, err
			}
			v[i] = os.ExpandEnv(v2)
		}
		if k2 != k {
			delete(m, k)
		}
		(m)[k2] = v
	}
	return m, nil
}
func formatFilter(k string) (string, error) {
	var k2 string
	field, err := parsing.GetFilterField(k)
	if err != nil {
		return "", err
	}
	if len(field.Data) == 1 {
		k2 = k
	} else {
		for i, v := range field.Data {
			k2 += v
			if i < len(field.Filter) {
				k2 += random.Filter(&field.Filter[i])
			}
		}
	}
	return k2, nil
}
