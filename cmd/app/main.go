package main

import (
	"log"

	v1 "github.com/rhnvrm/test-submodule-versioning/providers/foo/v1"
	v2 "github.com/rhnvrm/test-submodule-versioning/providers/foo/v2"
	v3 "github.com/rhnvrm/test-submodule-versioning/providers/foo/v3"
)

func main() {
	one := v1.Foo{Bar: "bar"}
	two := v2.Foo{Bar2: "bar2"}
	three := v3.Foo{Bar3: "bar3"}

	log.Println(one)
	log.Println(two)
	log.Println(three)
}
