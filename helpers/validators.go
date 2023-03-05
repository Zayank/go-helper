package helpers

import (
	"reflect"
	"regexp"
	"strings"
)

type Input struct {
	value interface{}
}
type Args struct {
	value interface{}
}

func SetRules(value interface{}, rules string) string {

	if len(rules) < 1 {
		return "rules not defined"
	}

	var args = make([]reflect.Value, 1)
	m := Input{value}
	var res []reflect.Value

	methods := strings.Split(rules, "|")

	for _, call_this_method := range methods {

		args = nil
		re := regexp.MustCompile(`\[.*?\]`)
		argument_string := re.FindString(call_this_method)

		if len(argument_string) > 0 {
			argument := strings.Trim(argument_string, "[]")
			args = append(args, reflect.ValueOf(VarToInt64(argument)))
			call_this_method = strings.Replace(call_this_method, argument_string, "", -1)
		}

		call_this_method = strings.Title(strings.ToLower(call_this_method))
		method := reflect.ValueOf(m).MethodByName(call_this_method)

		if method.Kind() != 0 {
			res = method.Call(args)

			message := res[0].String()
			if message != "" {
				return message
			}

		}
	}

	return ""

}

func (m Input) Required() (message string) {
	message = ""
	if reflect.ValueOf(m.value).Len() < 1 {
		return "this field is required"
	}
	return
}

func (m Input) Greater_than(v int64) (message string) {
	message = ""
	if VarToInt64(m.value) <= v {
		return "the value should be greater than " + VarToString(v)
	}
	return

}

func (m Input) Less_than(v int64) (message string) {
	message = ""
	if VarToInt64(m.value) >= v {
		message = "the value should be less than " + VarToString(v)
	}
	return

}

func (m Input) Is_array() (message string) {
	message = ""
	if reflect.TypeOf(m.value).Kind() != reflect.Array {
		message = "this field should be an array"
	}
	return

}

func (m Input) Is_alphanumeric() (message string) {
	message = ""
	is_alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(VarToString(m.value))
	if !is_alphanumeric {
		message = "this field should contain only alphanumeric characters"
	}
	return
}

func (m Input) Is_alpha() (message string) {
	message = ""
	is_alpha := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(VarToString(m.value))
	if !is_alpha {
		message = "this field should contain only alphabets"
	}
	return
}

func (m Input) Is_numeric() (message string) {
	message = ""
	is_alphanumeric := regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`).MatchString(VarToString(m.value))
	if !is_alphanumeric {
		message = "this field should contain only numeric characters"
	}
	return
}
