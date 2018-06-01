package jtest1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNo1(t *testing.T) {
	ret := JenkinsTest1()
	assert.Contains(t, ret, "Hello from ")
}

func TestNo2(t *testing.T) {
	ret := JenkinsTest2()
	assert.Contains(t, ret, "test #2")
}
