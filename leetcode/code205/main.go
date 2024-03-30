package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("foo", "ewg"))
}

type OneWayMap struct {
	data map[byte]byte
}

// NewOneWayMap initializes a new OneWayMap
func NewOneWayMap() *OneWayMap {
	return &OneWayMap{
		data: make(map[byte]byte),
	}
}

// Set sets the value associated with a key
// This operation is not allowed in a one-directional map
func (m *OneWayMap) Set(key byte, value byte) {
	// Do nothing
}

// Get retrieves the value associated with a key
func (m *OneWayMap) Get(key byte) (byte, bool) {
	value, ok := m.data[key]
	return value, ok
}

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sToT := NewOneWayMap()
	tToS := NewOneWayMap()

	for i := 0; i < len(s); i++ {
		if val, ok := sToT.Get(s[i]); ok && val != t[i] {
			fmt.Println(string(val), string(t[i]))
			return false
		}
		if val, ok := tToS.Get(t[i]); ok && val != s[i] {
			return false
		}
		sToT.data[s[i]] = t[i]
		tToS.data[t[i]] = s[i]
	}

	return true
}
