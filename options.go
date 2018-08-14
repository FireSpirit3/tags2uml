
//   tags2uml
//   Copyright 2014 ruben2020 https://github.com/ruben2020/ 
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import "flag"

const ( 
    NONE           = 0
    ALL            = 1
    ONLYPUBLIC     = 2 // only public
    ONLYPUBLICPROT = 3 // only public and protected
)

// options and their default values
var opt_blackbox bool = false
var opt_inherit bool = true
var opt_relationship bool = true
var opt_methods int = ONLYPUBLIC
var opt_members int = ONLYPUBLIC

var input_file string = "tags"
var output_file string = "-"
var help bool = false
var ver bool = false

func checkRange() bool {
    retval := true
    if (opt_methods < 0)||(opt_methods > 3) {retval = false}
    if (opt_members < 0)||(opt_members > 3) {retval = false}
    return retval
}

func parse_options() {
    flag.BoolVar(&opt_blackbox, "blackbox", false, "true 是黑盒模型, false 是白盒模型 (默认=false)")
    flag.BoolVar(&opt_inherit, "inherit", true, "true 是展示继承信息, false 是不展示 (默认=true)")
    flag.BoolVar(&opt_relationship, "relations", true, "true 是展示关联信息, false 是不展示 (默认=true)")
    flag.IntVar(&opt_methods, "methods", 2, "0=方法不展示, 1=所有方法, 2=仅public (默认), 3=仅 public 和 protected")
    flag.IntVar(&opt_members, "members", 2, "0=成员不展示, 1=所有成员, 2=仅public (默认), 3=仅 public 和 protected")
    flag.StringVar(&input_file, "infile", "tags", "path to input file (default=\"tags\")")
    flag.StringVar(&output_file, "outfile", "-", "path to output file, use \"-\" for stdout (default=\"-\")")
    flag.BoolVar(&help, "help", false, "print help message")
    flag.BoolVar(&ver, "ver", false, "print version")
    flag.Parse()
}

