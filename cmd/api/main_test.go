package main

import (
	"testing"
)

func Test_Main_Ok(t *testing.T) {
	go func() {
		main()
	}()
}
