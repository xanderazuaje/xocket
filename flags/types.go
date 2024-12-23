package flags

import (
	"errors"
	"strings"
)

type RunType uint8

const (
	RunAll RunType = iota
	RunDebug
)

type RunTypeArr []RunType

func (r *RunTypeArr) String() string {
	var strs []string
	var mapRun = map[RunType]string{
		RunAll:   "all",
		RunDebug: "debug",
	}
	for _, v := range *r {
		strs = append(strs, mapRun[v])
	}
	return strings.Join(strs, ",")
}

func (r *RunTypeArr) Set(value string) error {
	var runMap = map[string]RunType{
		"all":   RunAll,
		"debug": RunDebug,
	}
	strs := strings.Split(value, ",")
	for _, v := range strs {
		val, ok := runMap[v]
		if !ok {
			return errors.New("unsupported run type: " + v)
		}
		*r = append(*r, val)
	}
	return nil
}

func (r *RunTypeArr) Contains(run RunType) bool {
	for _, v := range *r {
		if v == run {
			return true
		}
	}
	return false
}
