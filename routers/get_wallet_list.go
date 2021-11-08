package routers

import (
	"red_envelop_server/sql"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func LoadWalletList(e *gin.Engine) {
	e.POST("/get_wallet_list", WalletListHandler)
}

func WalletListHandler(c *gin.Context) {

	uid, _ := c.GetPostForm("uid")
	logs.Printf("query %s 's wallets", uid)

	int_uid, _ := strconv.ParseInt(uid, 10, 64)

	//现在主要是针对redis查询
	//需要redis提供用户红包列表(id信息)func GetEnvelopeIdsByUid(uid int64) []int64
	//需要redis提供用户红包列表(time信息)func GetEnvelopeTimesByUid(uid int64) []int64
	//需要redis提供根据envelope_id查询红包钱数，func GetValueByEnvelopeId(envelope_id int64) int64
	//剩下的排序和构造json之前已经有了

	envelopes, _ := sql.GetAllEnvelopesByUID(int_uid)
	//先安装时间排序
	sort.SliceStable(envelopes, func(i, j int) bool {
		return envelopes[i].SnatchTime < envelopes[j].SnatchTime
	})

	var amount int64 = 0
	var myArray []map[string]interface{}

	for i := 0; i < len(envelopes); i++ {
		var curEnvelope map[string]interface{} = make(map[string]interface{})
		if envelopes[i].Opened == false {
			curEnvelope["envelope_id"] = envelopes[i].ID
			curEnvelope["opened"] = false
			curEnvelope["snatch_time"] = envelopes[i].SnatchTime

		} else {
			curEnvelope["envelope_id"] = envelopes[i].ID
			curEnvelope["value"] = envelopes[i].Value
			curEnvelope["opened"] = true
			curEnvelope["snatch_time"] = envelopes[i].SnatchTime
			amount = amount + envelopes[i].Value
		}
		myArray = append(myArray, curEnvelope)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"amount":        amount,
			"envelope_list": myArray,
		},
	})
}
