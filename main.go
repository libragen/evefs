// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package main

import (
	_ "github.com/dejavuzhou/spookyfs/initialization"
	"github.com/dejavuzhou/spookyfs/store"
	"github.com/sirupsen/logrus"
)

func main() {
	b := store.NewBarn("127.0.0.1:8787", "temp", 4)
	//b.Db.Put([]byte("key"), []byte("awesome is mime!"), nil)
	data, _ := b.Db.Get([]byte("key"), nil)
	logrus.Info(string(data))
	//cmd.Execute()

}