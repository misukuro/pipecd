// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePipedKey(t *testing.T) {
	key, hash, err := GeneratePipedKey()
	assert.NoError(t, err)
	assert.True(t, len(key) > 0)
	assert.True(t, len(hash) > 0)

	p := &Piped{
		KeyHash: hash,
	}
	err = p.CompareKey(key)
	assert.NoError(t, err)

	err = p.CompareKey("invalid")
	assert.Error(t, err)
}

func TestGenerateRandomString(t *testing.T) {
	validator := func(s string) error {
		for _, c := range s {
			if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
				continue
			}
			return fmt.Errorf("invalid character: %s", c)
		}
		return nil
	}

	s1 := generateRandomString(10)
	assert.Equal(t, 10, len(s1))
	assert.NoError(t, validator(s1))

	s2 := generateRandomString(10)
	assert.Equal(t, 10, len(s2))
	assert.NoError(t, validator(s2))

	assert.NotEqual(t, s1, s2)
}
