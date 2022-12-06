package config

import (
	"time"

	"github.com/pkg/errors"
)

const (
	intType      = "int"
	uintType     = "uint"
	strType      = "string"
	boolType     = "bool"
	durationType = "duration"
)

// Entry represents some value that was passed in config
type Entry struct {
	Val interface{}
	T   string
}

// Validate casts interface to type and returns if type mismatches
func (e *Entry) Validate() error {
	switch e.T {
	case intType:
		if v, ok := e.Val.(float64); ok {
			e.Val = int(v)
		}
		if v, ok := e.Val.(int); ok {
			e.Val = v
			return nil
		}
		return errors.New("int type mismatch")
	case uintType:
		if v, ok := e.Val.(uint); ok {
			e.Val = v
			return nil
		}
		return errors.New("uint type mismatch")
	case strType:
		if v, ok := e.Val.(string); ok {
			e.Val = v
			return nil
		}
		return errors.New("string type mismatch")
	case boolType:
		if v, ok := e.Val.(bool); ok {
			e.Val = v
			return nil
		}
		return errors.New("bool type mismatch")
	case durationType:
		v, err := time.ParseDuration(e.Val.(string))
		if err != nil {
			return errors.Wrap(err, "duration type mismatch")
		}
		e.Val = v
	default:
		return errors.Errorf("unknown type %s", e.T)
	}
	return nil
}

// ToInt converts entry to int
func (e Entry) ToInt() int {
	if e.T != intType {
		return 0
	}

	return e.Val.(int)
}

// ToString converts entry to string
func (e Entry) ToString() string {
	if e.T != strType {
		return ""
	}

	return e.Val.(string)
}

// ToBool converts entry to bool
func (e Entry) ToBool() bool {
	if e.T != boolType {
		return false
	}

	return e.Val.(bool)
}

// ToDuration converts entry to Duration
func (e Entry) ToDuration() time.Duration {
	if e.T != durationType {
		return 0
	}

	return e.Val.(time.Duration)
}

// ToUInt converts entry to uint
func (e Entry) ToUInt() uint {
	if e.T != uintType {
		return 0
	}

	return e.Val.(uint)
}
