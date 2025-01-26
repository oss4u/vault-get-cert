/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "vault-get-cert/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
