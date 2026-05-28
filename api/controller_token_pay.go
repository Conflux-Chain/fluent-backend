package api

import (
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type TokenPayController struct {
	services service.Services
}

func NewTokenPayController(services service.Services) *TokenPayController {
	return &TokenPayController{services}
}

// Config returns the token-pay configuration exposed to clients.
//
// @ID				tokenPayConfig
// @Summary			Get token-pay configuration
// @Description		Returns the supported ERC20 token list and the configured recipient address for token-pay.
// @Tags			TokenPay
// @Accept			json
// @Produce			json
// @Success			200	{object}	api.BusinessError{data=TokenPayConfig}	"Token-pay configuration"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/tokenpay/config	[get]
func (controller *TokenPayController) Config(c *gin.Context) (any, error) {
	config := controller.services.TokenPay.Config()

	var tokens []common.Address
	for _, token := range config.Tokens {
		tokens = append(tokens, token)
	}

	return TokenPayConfig{
		Tokens:    tokens,
		Recipient: config.Recipient.Hex(),
	}, nil
}

// Submit accepts two user-signed raw transactions (transfer token + business), validates
// both transactions on the backend, then starts sponsor flow.
//
// Flow overview:
//  1. Validate request payload and decode raw tx hex.
//  2. Validate the two transactions (signature/nonce/order/gas/token constraints) in service layer.
//  3. Send a funding ETH transaction from sponsor account.
//  4. Asynchronously wait for funding confirmation, then broadcast transfer-token and business txs.
//
// @ID				tokenPaySubmit
// @Summary			Submit token-pay transactions
// @Description		Submits two raw transactions (transfer token and business) for sponsored execution.
// @Description		Backend validates both transactions first. If validation passes, it sends a funding ETH transaction.
// @Description		After funding is confirmed asynchronously, backend broadcasts transfer-token tx then business tx.
// @Description		Clients can poll on-chain status of both submitted transactions (transfer-token and business) by their tx hashes.
// @Description		The whole flow is considered completed only when both transactions are successfully executed.
// @Description		If transfer-token execution fails, the user may be blacklisted by risk control.
// @Description		If transactions remain pending for a long time, contact administrator for investigation.
// @Tags			TokenPay
// @Accept			json
// @Produce			json
// @Param			request	body	TokenPayRequest	true	"Raw transfer-token tx and raw business tx"
// @Success			200	{object}	api.BusinessError	"Submission accepted"
// @Failure			600	{object}	api.BusinessError{data=string}	"Internal server error"
// @Router			/tokenpay/submit	[post]
func (controller *TokenPayController) Submit(c *gin.Context) (any, error) {
	var input TokenPayRequest

	if err := c.ShouldBind(&input); err != nil {
		return nil, api.ErrValidation(err)
	}

	rawTransferTokenTx, err := hexutil.Decode(input.RawTransferTokenTx)
	if err != nil {
		return nil, api.ErrValidation(errors.WithMessage(err, "Failed to hex decode transfer token tx"))
	}

	rawBusinessTx, err := hexutil.Decode(input.RawBusinessTx)
	if err != nil {
		return nil, api.ErrValidation(errors.WithMessage(err, "Failed to hex decode business tx"))
	}

	return nil, controller.services.TokenPay.Sponsor(rawTransferTokenTx, rawBusinessTx)
}
