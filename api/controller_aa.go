package api

import (
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
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

func (controller *AccountAbstractController) GetAuthStatus(c *gin.Context) (any, error) {
	txHash := c.Param("txhash")
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
