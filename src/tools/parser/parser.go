package parser

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"logging"
	"github.com/go-errors/errors"
)

var logger = logging.NewLogger("parser")

type parser struct {
	matches [][]string
	re      *regexp.Regexp
	Err     error
}

func New(re *regexp.Regexp, input string) *parser {
	return &parser{
		matches: re.FindAllStringSubmatch(input, -1),
		re:      re,
	}
}

// Result fetches a single result from a regex submatch
func (p *parser) NextResult(result interface{}) bool {
	logger.Debug("Result", "matches", fmt.Sprintf("%#v", p.matches))
	if len(p.matches) < 1 {
		return false
	}

	submatches := p.matches[0]
	p.matches = p.matches[1:]

	values := matchMap(p.re, submatches)
	if values == nil {
		p.Err = errors.Errorf("Could not parse input")
		return false
	}

	if reflect.TypeOf(result).Kind() != reflect.Ptr {
		p.Err = errors.Errorf("result parameter should be a pointer to a struct and not a %v", reflect.TypeOf(result))
		return false
	}

	resultValue := reflect.ValueOf(result).Elem()
	resultType := resultValue.Type()

	if resultType.Kind() != reflect.Struct {
		p.Err = errors.Errorf("result parameter should be a pointer to a struct and not a %v", reflect.TypeOf(result))
		return false
	}

	for i := 0; i < resultValue.NumField(); i++ {
		field := resultValue.Field(i)
		key := resultType.Field(i).Name

		stringValue, ok := values[key]
		if !ok {
			// Some values might be calculated and not appear in the regex
			continue
		}

		switch field.Type().Kind() {
		case reflect.Float64:
			value, err := strconv.ParseFloat(stringValue, 64)
			if err != nil {
				p.Err = err
				return false
			}
			field.SetFloat(value)
		case reflect.Int:
			value, err := strconv.ParseInt(stringValue, 10, 64)
			if err != nil {
				p.Err = err
				return false
			}
			field.SetInt(value)
		case reflect.String:
			field.SetString(stringValue)
		default:
			p.Err = errors.Errorf("unsupported type: %v", field.Type().Name())
			logger.Error("parser error", "err", p.Err)
			return false
		}
	}

	return true
}

func matchMap(re *regexp.Regexp, submatches []string) map[string]string {
	if submatches == nil {
		return nil
	}

	result := map[string]string{}
	for i, name := range re.SubexpNames() {
		result[name] = submatches[i]
	}
	return result
}
