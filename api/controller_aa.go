package api

import (
	"math/big"

	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/holiman/uint256"
	"github.com/openweb3/web3go/types/enums"
)

type AccountAbstractController struct {
	services service.Services
}

func NewAccountAbstractController(services service.Services) *AccountAbstractController {
	return &AccountAbstractController{services}
}

// SendAuth validates and broadcasts an EIP-7702 set-code transaction on behalf of the authority
// identified by the provided signed authorization message.
//
// The server's configured fee-payer account will cover the gas cost. So, the wallet user only needs to
// sign the EIP-7702 auth message with their EOA private key (no ETH balance required on the EOA).
//
// @ID				aaAuthSend
// @Summary			Submit EIP-7702 authorization
// @Description		Validates the signed EIP-7702 authorization and broadcasts a set-code transaction.
// @Description		The service will sponsor the set-code transaction gas fee, and user needs to poll GET /aa/auth/{txHash} for the result.
// @Tags			AccountAbstract
// @Accept			json
// @Produce			json
// @Param			auth	body	SetCodeAuth	true	"Signed EIP-7702 authorization message"
// @Success			200	{object}	api.BusinessError{data=string}	"Transaction hash (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/auth	[post]
func (controller *AccountAbstractController) SendAuth(c *gin.Context) (any, error) {
	var input SetCodeAuth

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	r, err := uint256.FromHex(input.R)
	if err != nil {
		return nil, api.ErrValidationStrf("Invalid signature R: %v", err.Error())
	}

	s, err := uint256.FromHex(input.S)
	if err != nil {
		return nil, api.ErrValidationStrf("Invalid signature S: %v", err.Error())
	}

	return controller.services.AccountAbstract.SendSetCodeTransaction(gethTypes.SetCodeAuthorization{
		ChainID: *uint256.NewInt(input.ChainId),
		Address: common.HexToAddress(input.Contract),
		Nonce:   input.Nonce,
		V:       input.V,
		R:       *r,
		S:       *s,
	})
}

// GetAuthStatus returns the on-chain result of an EIP-7702 set-code transaction.
//
// Polling pattern: call this endpoint repeatedly after submitting via POST /aa/auth.
//   - If the response contains code 1001 (transaction not found), the transaction may have been
//     dropped. In this case, the wallet should re-submit the authorization.
//   - If executed=false, the transaction has not yet been executed, user should continue polling.
//   - If executed=true and success=true, the EOA is now delegated to the target contract.
//   - If executed=true and success=false, check the error field for the EVM failure reason.
//
// @ID				aaAuthStatus
// @Summary			Query EIP-7702 set-code result
// @Description		Returns the executed result of an EIP-7702 set-code transaction. Returns executed=false while the transaction is still pending.
// @Description		If the transaction is not found (business error code 1001), re-submit via POST /aa/auth.
// @Tags			AccountAbstract
// @Accept			json
// @Produce			json
// @Param			txHash	path	string	true	"0x-prefixed 32-byte transaction hash (66 characters)"
// @Success			200	{object}	api.BusinessError{data=SetCodeResult}	"Set-code result (check executed and success fields)"
// @Failure			600	{object}	api.BusinessError{data=string}			"Internal server error"
// @Router			/aa/auth/{txHash}	[get]
func (controller *AccountAbstractController) GetAuthStatus(c *gin.Context) (any, error) {
	txHash := c.Param("txHash")
	if len(txHash) == 0 {
		return nil, api.ErrValidationStr("Tx hash not specified")
	}

	if len(txHash) != 66 {
		return nil, api.ErrValidationStr("Invalid tx hash")
	}

	result, ok, err := controller.services.AccountAbstract.GetSetCodeResult(common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}

	if !ok {
		return SetCodeResult{}, nil
	}

	if result.Result == enums.SetAuthOutcomeSuccess {
		return SetCodeResult{
			Executed: true,
			Success:  true,
		}, nil
	}

	return SetCodeResult{
		Executed: true,
		Error:    string(result.Result),
	}, nil
}

// GasTankPrepareDeposit builds paymasterData for CREDIT mode deposit estimation.
//
// The returned paymasterData should be attached to the UserOperation when wallet clients estimate
// gas in CREDIT mode before submitting a signed operation.
//
// @ID				aaGasTankPrepareDeposit
// @Summary			Prepare paymasterData for CREDIT mode deposit estimation
// @Description		Returns paymasterData used to estimate UserOperation gas in CREDIT mode for token deposit flow.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			request		body	GasTankPrepareDepositRequest	true	"Paymaster data request"
// @Success			200	{object}	api.BusinessError{data=string}	"Paymaster and data (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/gastank/prepare/deposit	[get]
func (controller *AccountAbstractController) GasTankPrepareDeposit(c *gin.Context) (any, error) {
	var input GasTankPrepareDepositRequest

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	token := common.HexToAddress(input.Token)
	amount, ok := new(big.Int).SetString(input.Amount, 10)
	if !ok {
		return nil, api.ErrValidationStr("Invalid amount")
	}

	paymasterData := controller.services.GasTank.PrepareDeposit(token, amount)

	return hexutil.Encode(paymasterData), nil
}

// GasTankPrepare builds paymasterData for normal UserOperation gas estimation in REFUND mode.
//
// The returned paymasterData should be attached to the UserOperation when wallet clients estimate
// gas in REFUND mode before signing and sending the operation.
//
// @ID				aaGasTankPrepare
// @Summary			Prepare paymasterData for REFUND mode estimation
// @Description		Returns paymasterData used to estimate gas for a normal UserOperation in REFUND mode.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			request		body	GasTankPrepareRequest	true	"Paymaster data request"
// @Success			200	{object}	api.BusinessError{data=string}	"Paymaster and data (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/gastank/prepare	[get]
func (controller *AccountAbstractController) GasTankPrepare(c *gin.Context) (any, error) {
	var input GasTankPrepareRequest

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	sender := common.HexToAddress(input.Sender)
	token := common.HexToAddress(input.Token)

	paymasterData, err := controller.services.GasTank.Prepare(sender, token)
	if err != nil {
		return nil, err
	}

	return hexutil.Encode(paymasterData), nil
}

// GasTankSign calculates maxTokenCost for a UserOperation, signs it with paymaster key, and
// returns the reassembled paymasterData.
//
// The wallet should call this endpoint after gas estimation, then place the returned paymasterData
// into UserOperation.paymasterData before final submission.
//
// @ID				aaGasTankSign
// @Summary			Sign paymasterData for UserOperation
// @Description		Calculates maxTokenCost for the given UserOperation, adds paymaster signature, and returns reassembled paymasterData.
// @Tags			GasTank
// @Accept			json
// @Produce			json
// @Param			userOp	body	UserOperation	true	"UserOperation for maxTokenCost estimation and paymaster signing"
// @Success			200	{object}	api.BusinessError{data=string}	"Signed and reassembled paymasterData (0x-prefixed hex)"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/aa/gastank/sign	[post]
func (controller *AccountAbstractController) GasTankSign(c *gin.Context) (any, error) {
	var input UserOperation

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	userOp := input.ToPackedUserOperation()

	paymasterData, err := controller.services.GasTank.Sign(userOp)
	if err != nil {
		return nil, err
	}

	return hexutil.Encode(paymasterData), nil
}
