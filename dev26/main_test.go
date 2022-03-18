package main

import (
	"strings"
	"testing"
)

type areCharactersUniqueTest struct {
	str      string
	expected bool
}

func createHugeString(n int) string {
	var sb strings.Builder
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(0)
	}
	return sb.String()
}

var areCharactersUniqueTests = []areCharactersUniqueTest{
	{"abcd", true},
	{"abCdefAaf", false},
	{"aabcd", false},
	{"", true},
	{"a", true},
	{"ascLBJLXRxnT", false},
}

func TestAreCharactersUniqueV1(t *testing.T) {
	for _, test := range areCharactersUniqueTests {
		if output := areCharactersUniqueV1(test.str); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

func TestAreCharactersUniqueV2(t *testing.T) {
	for _, test := range areCharactersUniqueTests {
		if output := areCharactersUniqueV2(test.str); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

func TestAreCharactersUniqueV3(t *testing.T) {
	for _, test := range areCharactersUniqueTests {
		if output := areCharactersUniqueV3(test.str); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

func BenchmarkAreCharactersUniqueV1(b *testing.B) {
	str := createHugeString(100000)
	for i := 0; i < b.N; i++ {
		areCharactersUniqueV1(str)
	}
}

func BenchmarkAreCharactersUniqueV2(b *testing.B) {
	str := createHugeString(100000)
	for i := 0; i < b.N; i++ {
		areCharactersUniqueV2(str)
	}
}

func BenchmarkAreCharactersUniqueV3(b *testing.B) {
	str := createHugeString(100000)
	for i := 0; i < b.N; i++ {
		areCharactersUniqueV3(str)
	}
}
