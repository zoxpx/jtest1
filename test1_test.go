package jtest1

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNo1(t *testing.T) {
    ret := JenkinsTest1()
    assert.Contains(t, ret, "Hello from ")
}

