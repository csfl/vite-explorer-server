package response

import (
	"math/big"
	"github.com/gin-gonic/gin"
)

type General struct {
	FfmCap string
	SysTps string
	CirPrice int
	TxTotalTAmount *big.Int
	ChainHeight *big.Int
	TxMonAmount *big.Int
}

func NewGeneral (ffmCap string, sysTps string, cirPrice int,
	txTotalTAmount *big.Int, chainHeight *big.Int, txMonAmount *big.Int) *General {
	var general = &General{
		FfmCap: ffmCap,
		SysTps: sysTps,
		CirPrice: cirPrice,
		TxTotalTAmount: txTotalTAmount,
		ChainHeight: chainHeight,
		TxMonAmount: txMonAmount,
	}

	return general
}

func (gl *General) ToResponse() gin.H {
	return gin.H{
		"ffmCap": gl.FfmCap,
		"sysTps": gl.SysTps,
		"cirPrice": gl.CirPrice,
		"txTotalTAmount": gl.TxTotalTAmount.String(),
		"chainHeight": gl.ChainHeight.String(),
		"txMonAmount": gl.TxMonAmount.String(),
	}
}