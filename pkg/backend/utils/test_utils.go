package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AbortIfNotTesting(t *testing.T) {
	if !IsTestEnv() {
		assert.Fail(t, "should only be calling this method while testing")
	}
}
