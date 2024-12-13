package parsing

import (
	"errors"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetFilter(str string) (*Field, error) {
	field := Field{}
	err := validateTags(str)
	if err != nil {
		return nil, err
	}
	openTagsIdx := getTagsIdx(str, '<', 1)
	if len(openTagsIdx) == 0 {
		field.Data = append(field.Data, str)
		return &field, nil
	}
	closeTagsIdx := getTagsIdx(str, '>', -1)
	idx := 0
	for i := 0; ; i++ {
		if i > len(openTagsIdx) {
			suffix := str[idx+1:]
			if suffix != "" {
				field.Data = append(field.Data, suffix)
			}
			break
		}
		field.Data = append(field.Data, str[idx:openTagsIdx[i]]) // Save every string before the <> arguments
		filterStr := str[openTagsIdx[i]:closeTagsIdx[i]]         // Save the enclosed argument between <>
		println(filterStr)
		idx = closeTagsIdx[i] // Move the index to the last tag
		filterFields := strings.Split(filterStr, ":")
		var filter Filter
		for i, v := range filterFields {
			if i == 0 {
				err := setFilter(v, &filter)
				if err != nil {
					return nil, err
				}
				continue
			}
			specs := strings.Split(v, "=")
			if len(specs) > 2 {
				return nil, errors.New("ambiguous field specification at: " + v)
			}
			err := setFilterSpecs(specs, filter)
			if err != nil {
				return nil, err
			}
			log.Println(field)
			field.Filter = append(field.Filter, filter)
		}
	}
	return &field, nil
}

func setFilterSpecs(specs []string, filter Filter) error {
	switch specs[0] {
	case "min":
		fmin, err := strconv.ParseFloat(specs[1], 64)
		if err != nil {
			return errors.New("error at: " + specs[1] + " -> " + err.Error())
		}
		if filter.Type == "string" && fmin != math.Trunc(fmin) {
			return errors.New("decimal values are not valid when argument is 'string'")
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
		filter.Max = fmax
	case "len":
		length, err := strconv.Atoi(specs[1])
		if err != nil {
			return errors.New("error at: " + specs[1] + " -> " + err.Error())
		}
		filter.Len = length
	case "r":
		r, err := regexp.Compile(specs[1])
		if err != nil {
			return errors.New("error at: " + specs[1] + " -> " + err.Error())
		}
		filter.Regex = r
	default:
		return errors.New("unknown filter: " + specs[1])
	}
	return nil
}

func setFilter(v string, filter *Filter) error {
	filterType := FilterType(v)
	if !filterType.IsValid() {
		return errors.New("invalid filter type: " + v)
	}
	filter.Type = filterType
	return nil
}

func getTagsIdx(str string, tag rune, padding int) []int {
	var tagsIdx []int
	var j int
	str2 := str
	for {
		str2 = str2[j:]
		println(str2)
		j := strings.IndexRune(str2, tag)
		if j > 0 && str2[j-1] == '\\' {
			str2 = str2[j:]
			continue
		}
		if j == -1 {
			break
		}
		if len(tagsIdx) == 0 {
			tagsIdx = append(
				tagsIdx,
				j+padding,
			)
		} else {
			tagsIdx = append(
				tagsIdx,
				j+tagsIdx[len(tagsIdx)-1],
			)
		}
		str2 = str2[j+1:]
	}
	return tagsIdx
}

func validateTags(str string) error {
	var hasValidTags [2]bool
	var skip bool
	for _, v := range str {
		if skip == true {
			skip = false
			continue
		} else if v == '\\' {
			skip = true
			continue
		}
		if (v == '<' && hasValidTags[0] == true) ||
			(v == '>' && hasValidTags[1] == true) {
			return errors.New(
				"nested <> arguments are not valid",
			)
		}
		if v == '<' {
			hasValidTags[0] = true
		} else if v == '>' {
			hasValidTags[1] = true
		}
		if hasValidTags[0] && hasValidTags[1] {
			hasValidTags[0] = false
			hasValidTags[1] = false
		}
	}
	if hasValidTags[0] != hasValidTags[1] {
		return errors.New("invalid <> arguments syntax")
	}
	return nil
}
