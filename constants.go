package main

import "github.com/charmbracelet/lipgloss"

var errorCodesDir = "errors"
var docStyle = lipgloss.NewStyle().Margin(1, 2)
var errorCodesSource = map[string]string{
	"anchor":              "https://github.com/coral-xyz/anchor/blob/master/lang/src/error.rs",
	"spl-token":           "https://github.com/solana-labs/solana-program-library/blob/master/token/program/src/error.rs",
	"spl-token-2022":      "https://github.com/solana-labs/solana-program-library/blob/master/token/program-2022/src/error.rs",
	"account-compression": "https://github.com/solana-labs/solana-program-library/blob/master/account-compression/programs/account-compression/src/error.rs",
}
