/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitKeyAndValue(t *testing.T) {
	key, value := SplitKeyAndValue("key=value", '=')
	require.Equal(t, "key", key)
	require.Equal(t, "value", value)

	key, value = SplitKeyAndValue("nosep", '=')
	require.Equal(t, "", key)
	require.Equal(t, "", value)
}

func TestStringRepeat(t *testing.T) {
	require.Equal(t, "", StringRepeat("%v", "-", 0))
	require.Equal(t, "%v", StringRepeat("%v", "-", 1))
	require.Equal(t, "%v-%v", StringRepeat("%v", "-", 2))
}
