package general

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/util"
	"github.com/vitelabs/vite-explorer-server/type/response"
	serviceSnapshotChain "github.com/vitelabs/vite-explorer-server/service/snapshotchain"
	serviceGeneral "github.com/vitelabs/vite-explorer-server/service/general"
)

func Detail (c *gin.Context) {
	ffmCap, err := serviceGeneral.GetFfmCap()
	if err != nil {
		util.RespondFailed(c, 1, err, "")
		return
	}
	sysTps, err := serviceGeneral.GetSysTps()
	if err != nil {
		util.RespondFailed(c, 2, err, "")
		return
	}
	cirPrice, err := serviceGeneral.GetCirPrice()
	if err != nil {
		util.RespondFailed(c, 3, err, "")
		return
	}
	txTotalTAmount, err := serviceGeneral.GetTxTotalTAmount()
	if err != nil {
		util.RespondFailed(c, 4, err, "")
		return
	}
	chainHeight, err := serviceSnapshotChain.GetSnapshotChainHeight()
	if err != nil {
		util.RespondFailed(c, 5, err, "")
		return
	}
	txMonAmount, err := serviceGeneral.GetTxMonAmount()
	if err != nil {
		util.RespondFailed(c, 6, err, "")
		return
	}
	data := response.NewGeneral(ffmCap, sysTps, cirPrice, txTotalTAmount, chainHeight, txMonAmount)
	util.RespondSuccess(c, data,"")
}