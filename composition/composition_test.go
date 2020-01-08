/*
 * Copyright (c) 2013 Zhen, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package composition

import (
	"log"
	"testing"

	"github.com/Akron/encoding"
	"github.com/Akron/encoding/benchtools"
	"github.com/Akron/encoding/bp32"
	dbp32 "github.com/Akron/encoding/delta/bp32"
	dvb "github.com/Akron/encoding/delta/variablebyte"
	"github.com/Akron/encoding/generators"
	"github.com/Akron/encoding/variablebyte"
)

var (
	codec encoding.Integer
	data  []int32
	size  int = 10000000
)

func init() {
	log.Printf("composition_test/init: generating %d uint32s\n", size)
	data = generators.GenerateClustered(size, size*2)
	log.Printf("composition_test/init: generated %d integers for test", size)
}

func TestDeltaBP32andDeltaVariableByte(t *testing.T) {
	sizes := []int{100, 100 * 10, 100 * 100, 100 * 1000, 100 * 10000}
	benchtools.TestCodec(New(dbp32.New(), dvb.New()), data, sizes)
}

func TestBP32andVariableByte(t *testing.T) {
	sizes := []int{100, 100 * 10, 100 * 100, 100 * 1000, 100 * 10000}
	benchtools.TestCodec(New(bp32.New(), variablebyte.New()), data, sizes)
}
