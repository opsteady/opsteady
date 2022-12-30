package utils

import (
	"reflect"
	"sort"
	"testing"
)

func TestUniqueNonEmptyElementsOf(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"foo", "bar", "baz", "foo"},
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    []string{"foo", "", "bar", "", "baz", "foo"},
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    []string{"", "", "", "", "", ""},
			expected: []string{},
		},
		{
			input:    []string{"foo"},
			expected: []string{"foo"},
		},
		{
			input:    []string{},
			expected: []string{},
		},
	}

	for _, test := range tests {
		result := UniqueNonEmptyElementsOf(test.input)
		sort.Strings(result)
		sort.Strings(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("UniqueNonEmptyElementsOf(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestReverseStringArray(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"foo", "bar", "baz"},
			expected: []string{"baz", "bar", "foo"},
		},
		{
			input:    []string{"foo", "bar"},
			expected: []string{"bar", "foo"},
		},
		{
			input:    []string{"foo"},
			expected: []string{"foo"},
		},
		{
			input:    []string{},
			expected: []string{},
		},
	}

	for _, test := range tests {
		result := ReverseStringArray(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ReverseStringArray(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
