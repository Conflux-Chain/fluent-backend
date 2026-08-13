package main

import "github.com/Conflux-Chain/fluent-backend/cmd"

// @title		Fluent Backend API
// @version		1.0
// @description	Fluent Backend provides account abstraction, gas tank paymaster, and ERC20 token gas payment services.
// @BasePath	/api
//
// @tag.name		AccountAbstract
// @tag.description	Provides account abstraction features, including free EOA to smart account upgrades and upgrade status queries.
//
// @tag.name		GasTank
// @tag.description	Provides an off-chain gas tank sponsorship service for smart accounts, including paymaster data preparation and user operation signing for on-chain paymaster contract validation.
//
// @tag.name		TokenPay
// @tag.description	Provides gas fee sponsorship so EOA accounts can pay gas fees with supported ERC20 tokens.

func main() {
	cmd.Execute()
}
