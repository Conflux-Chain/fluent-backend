package api

type SetCodeAuth struct {
	ChainId  uint64 `json:"chainId"`
	Contract string `json:"contract" binding:"required,len=42"`
	Nonce    uint64 `json:"nonce"`
	V        uint8  `json:"v"`
	R        string `json:"r" binding:"required,min=65,max=66"`
	S        string `json:"s" binding:"required,min=65,max=66"`
}

type SetCodeResult struct {
	Executed bool   `json:"executed"` // whether tx executed
	Success  bool   `json:"success"`  // whether set code successfully
	Error    string `json:"error"`    // failure reason
}
