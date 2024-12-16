package parsing

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetFilter(str string) (*Field, error) {
	field := Field{}
	regex := regexp.MustCompile("<[a-z]\\w+(?::[a-z]\\w+=?[-a-z0-9]\\w?|:r='.*')*?>")
	data := regex.Split(str, -1)
	field.Data = data
	filterStrs := regex.FindAllString(str, -1)
	for _, v := range filterStrs {
		var filter Filter
		v = strings.Trim(v, "<>")
		sets := strings.Split(v, ":")
		filter.Type = FilterType(sets[0])
		if !filter.Type.IsValid() {
			return nil, errors.New("invalid filter type: " + v)
		}
		sets = sets[1:]
		for _, v := range sets {
			specs := strings.Split(v, "=")
			if len(specs) > 2 {
				return nil, errors.New("ambiguous filter specification: " + v)
			}
			err := setFilterSpecs(specs, &filter)
			if err != nil {
				return nil, err
			}
		}
		field.Filter = append(field.Filter, filter)
	}
	return &field, nil
}

func setFilterSpecs(specs []string, filter *Filter) error {
	switch specs[0] {
	case "min":
		fmin, err := strconv.ParseFloat(specs[1], 64)
		if err != nil {
			return errors.New("error at: " + specs[1] + " -> " + err.Error())
		}
		if filter.Type == "string" && fmin != math.Trunc(fmin) {
			return errors.New("decimal values are not valid when argument is 'string'")
		}
		if filter.Type == "integer" && fmin != math.Trunc(fmin) {
			return errors.New("decimal values are not valid when argument is 'integer'")
		}
		filter.Min = fmin
	case "max":
		fmax, err := strconv.ParseFloat(specs[1], 64)
		if err != nil {
			return errors.New("error at: " + specs[1] + " -> " + err.Error())
		}
		if filter.Type == "string" && fmax != math.Trunc(fmax) {
			return errors.New("decimal values are not valid when argument is 'string'")
		}
		if filter.Type == "integer" && fmax != math.Trunc(fmax) {
			return errors.New("decimal values are not valid when argument is 'integer'")
		}
		filter.Max = fmax
	case "len":
		length, err := strconv.Atoi(specs[1])
		if err != nil {
			return errors.New("error at: " + specs[1] + " -> " + err.Error())
		}
		filter.Len = length
	case "r":
		r, err := regexp.Compile(specs[1][1 : len(specs[1])-1])
		if err != nil {
			return errors.New("error at: " + specs[1] + " -> " + err.Error())
		}
		filter.Regex = r
	default:
		return errors.New("unknown filter: " + specs[0])
	}
	return nil
}
