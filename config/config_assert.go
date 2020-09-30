package config

import (
	"fmt"
	"log"
	"reflect"
)

// Shamelessly writted from stretch/testify/assert so that I can use them as runtime/startup validation
// rather than compile time tests, now that the config file is going to be an env var.
func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

func assertFail(failureMessage string, msgAndArgs ...interface{}) {
	content := failureMessage
	message := messageFromMsgAndArgs(msgAndArgs...)
	if len(message) > 0 {
		content = fmt.Sprintf("%s -- %s", content, message)
	}

	log.Fatalf("\n%s", ""+content)
}

func assertFailf(failureMessage string, msg string, args ...interface{}) {
	assertFail(failureMessage, append([]interface{}{msg}, args...)...)
}

func assertNotEmptyf(object interface{}, msg string, args ...interface{}) {
	assertNotEmpty(object, append([]interface{}{msg}, args...)...)
}

func assertNotEmpty(object interface{}, msgAndArgs ...interface{}) {
	pass := !isEmpty(object)
	if !pass {
		assertFail(fmt.Sprintf("Should NOT be empty, but was %v", object), msgAndArgs...)
	}
}

func assertNilf(object interface{}, msg string, args ...interface{}) {
	assertNil(object, append([]interface{}{msg}, args...)...)
}

func assertNil(object interface{}, msgAndArgs ...interface{}) {
	if isNil(object) {
		return
	}
	assertFail(fmt.Sprintf("Expected nil, but got: %#v", object), msgAndArgs...)
}

func assertNotNilf(object interface{}, msg string, args ...interface{}) {
	assertNotNil(object, append([]interface{}{msg}, args...)...)
}

func assertNotNil(object interface{}, msgAndArgs ...interface{}) {
	if !isNil(object) {
		return
	}
	assertFail("Expected value not to be nil.", msgAndArgs...)
}

// isEmpty gets whether the specified object is considered empty or not.
func isEmpty(object interface{}) bool {
	// get nil case out of the way
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
		// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return isEmpty(deref)
		// for all other types, compare against the zero value
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

// isNil checks if a specified object is nil or not, without Failing.
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}
