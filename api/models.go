package api

type SetCodeAuth struct {
	// ChainID is the chain ID the authorization is bound to. 0 means the auth is valid on any chain.
	ChainId uint64 `json:"chainId"`
	// Contract is the 20-byte hex address (with 0x prefix) of the smart-contract to delegate the EOA to.
	// Set to "0x0000000000000000000000000000000000000000" to revoke an existing delegation.
	Contract string `json:"contract" binding:"required,len=42"`
	// Nonce is the current on-chain nonce of the signing authority (EOA). Must match exactly.
	Nonce uint64 `json:"nonce"`
	// V is the recovery identifier of the EIP-7702 authorization signature (0 or 1).
	V uint8 `json:"v"`
	// R is the R component of the EIP-7702 authorization signature, as a 0x-prefixed hex string.
	R string `json:"r" binding:"required"`
	// S is the S component of the EIP-7702 authorization signature, as a 0x-prefixed hex string.
	S string `json:"s" binding:"required"`
}

type SetCodeResult struct {
	// Executed indicates whether the transaction has been executed.
	Executed bool `json:"executed"`
	// Success indicates whether the set-code authorization succeeded. Only meaningful when Executed is true.
	Success bool `json:"success"`
	// Error contains the failure reason reported by the EVM when Success is false.
	Error string `json:"error"`
}
