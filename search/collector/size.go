//  Copyright (c) 2020 Bluge Labs, LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"reflect"
)

func init() {
	var ptr *int
	sizeOfPtr = int(reflect.TypeOf(ptr).Size())
	var str string
	sizeOfString = int(reflect.TypeOf(str).Size())
	var coll TopNCollector
	reflectStaticSizeTopNCollector = int(reflect.TypeOf(coll).Size())
}

var sizeOfPtr int
var sizeOfString int
var reflectStaticSizeTopNCollector int
