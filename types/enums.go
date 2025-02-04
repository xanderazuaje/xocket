package types

type ExpectedType string

const (
	ExpectedTypeJSON    ExpectedType = "json"
	ExpectedTypeText    ExpectedType = "text"
	ExpectedTypeBinary  ExpectedType = "bin"
	ExpectedTypeNothing ExpectedType = "nil"
)

var validExpectedType = map[ExpectedType]struct{}{
	ExpectedTypeJSON:    {},
	ExpectedTypeText:    {},
	ExpectedTypeBinary:  {},
	ExpectedTypeNothing: {},
}

func (e *ExpectedType) IsValid() bool {
	_, ok := validExpectedType[*e]
	return ok
}

type ExpectedShape string

const (
	ExpectedShapeArraySoft    ExpectedShape = "array-soft"
	ExpectedShapeArrayStrict  ExpectedShape = "array-strict"
	ExpectedShapeArrayWithout ExpectedShape = "array-without"
	ExpectedShapeSingle       ExpectedShape = "single"
	ExpectedShapeWithout      ExpectedShape = "without"
)

var validExpectedShape = map[ExpectedShape]struct{}{
	ExpectedShapeArraySoft:    {},
	ExpectedShapeArrayStrict:  {},
	ExpectedShapeArrayWithout: {},
	ExpectedShapeSingle:       {},
	ExpectedShapeWithout:      {},
}

func (e *ExpectedShape) IsValid() bool {
	_, ok := validExpectedShape[*e]
	return ok
}

type FilterType string

const (
	FilterInteger FilterType = "integer"
	FilterFloat   FilterType = "float"
	FilterString  FilterType = "string"
	FilterNil     FilterType = "nil"
)

var validFilterType = map[FilterType]struct{}{
	FilterInteger: {},
	FilterFloat:   {},
	FilterString:  {},
	FilterNil:     {},
}

func (e *FilterType) IsValid() bool {
	_, ok := validFilterType[*e]
	return ok
}
