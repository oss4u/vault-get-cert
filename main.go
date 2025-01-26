/*
Copyright © 2025 Marc Ende <me@e-beyond.de>
*/
package main

import "vault-get-cert/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
