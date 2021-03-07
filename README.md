```go
package main

import (
	"fmt"
	"time"

	"github.com/kpeu3i/go-tizen-tv"
)

func main() {
	manager := samsung.NewTVManager()
	
	tvs, err := manager.Discover()
	if err != nil {
		panic(err)
	}

	// Save configuration for the discovered TV-s
	err = manager.Store(tvs...) // or call "manager.Load" to create TV-s from the configuration
	if err != nil {
		panic(err)
	}

	for _, tv := range tvs {
		info, err := tv.Info()
		if err != nil {
			panic(err)
		}

		fmt.Println("********", info.Name, "********")
		fmt.Println("IMPORTANT! Confirm permissions on TV popup")
		fmt.Println("Sending keys...")

		keys := samsung.KeySequence{}
		keys.
			Click(samsung.KEY_VOLUP).
			Wait(500 * time.Millisecond).
			Repeat(5)

		err = tv.SendKeys(keys)
		if err != nil {
			panic(err)
		}

		fmt.Println("Opening Youtube...")

		err = tv.OpenApp("111299001912")
		if err != nil {
			panic(err)
		}
	}
}
```