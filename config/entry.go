package config

import "time"

// Entry represents some value that was passed in config
type Entry interface{}

// ToInt converts entry to int
func ToInt(e Entry) int {
	return e.(int)
}

// ToString converts entry to string
func ToString(e Entry) string {
	return e.(string)
}

// ToBool converts entry to bool
func ToBool(e Entry) bool {
	return e.(bool)
}

// ToDuration converts entry to Duration
func ToDuration(e Entry) time.Duration {
	return e.(time.Duration)
}

// ToUInt converts entry to uint
func ToUInt(e Entry) uint {
	return e.(uint)
}
