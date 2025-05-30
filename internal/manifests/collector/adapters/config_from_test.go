// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-telemetry/opentelemetry-operator/internal/manifests/collector/adapters"
)

func TestInvalidYAML(t *testing.T) {
	// test
	config, err := adapters.ConfigFromString("🦄")

	// verify
	assert.Nil(t, config)
	assert.Equal(t, adapters.ErrInvalidYAML, err)
}

func TestEmptyString(t *testing.T) {
	// test and verify
	res, err := adapters.ConfigFromString("")
	assert.NoError(t, err)
	assert.Empty(t, res, 0)
}
