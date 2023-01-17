package main

// examples at https://github.com/flant/libjq-go

import (
	"fmt"
	"sync"

	. "github.com/flant/libjq-go"
)

func main() {
	var res string
	var err error
	var inputJsons []string

	// 1. Run one program with one input.
	res, err = Jq().Program(".foo").Run(`{"foo":"bar"}`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("1. %s\n", res)
	// Should print
	// 1. "bar"

	// 3. Use program text as a key for a cache.
	inputJsons = []string{
		`{ "foo":"bar-quux" }`,
		`{ "foo":"baz-baz" }`,
		// ...
	}
	for _, data := range inputJsons {
		res, err = Jq().Program(".foo").Cached().Run(data)
		if err != nil {
			panic(err)
		}
		// Now do something with filter result ...
		fmt.Printf("2. %s\n", res)
	}
	// Should print
	// 2. "bar-quux"
	// 2. "baz-baz"

	// 3. Explicitly precompile jq expression.
	inputJsons = []string{
		`{ "bar":"Foo quux" }`,
		`{ "bar":"Foo baz" }`,
		// ...
	}
	prg, err := Jq().Program(".bar").Precompile()
	if err != nil {
		panic(err)
	}
	for _, data := range inputJsons {
		res, err = prg.Run(data)
		if err != nil {
			panic(err)
		}
		// Now do something with filter result ...
		fmt.Printf("3. %s\n", res)
	}
	// Should print
	// 3. "Foo quux"
	// 3. "Foo baz"

	// 4. It is safe to use Jq() from multiple go-routines.
	//    Note however that programs are executed synchronously.
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		res, err = Jq().Program(".foo").Run(`{"foo":"bar"}`)
		if err != nil {
			panic(err)
		}
		fmt.Printf("4. %s\n", res)
		wg.Done()
	}()
	go func() {
		res, err = Jq().Program(".foo").Cached().Run(`{"foo":"bar"}`)
		if err != nil {
			panic(err)
		}
		fmt.Printf("4. %s\n", res)
		wg.Done()
	}()
	go func() {
		res, err = Jq().Program(".foo").Cached().Run(`{"foo":"bar"}`)
		if err != nil {
			panic(err)
		}
		fmt.Printf("4. %s\n", res)
		wg.Done()
	}()
	wg.Wait()
	// Should print
	// 4. "bar"
	// 4. "bar"
	// 4. "bar"
}
