package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"red_envelop_server/routers"
	"red_envelop_server/sql"
	"red_envelop_server/utils"
)



func main() {

	//创建数据库连接
	db, err := sql.InitDB()
	if err != nil {
		log.Println("database connection failure")
	}
	defer db.Close()

	//读取配置文件（zyz）
	//初始化六个变量
	//每个用户最多可抢到的次数MaxTimes、抢到的概率Probability、总金额TotalMoney、总个数TotalNum、每个红包的金额范围[MaxMoney, MinMoney]
	//除probability类型为（Float64）外，其余变量均为（Int64）,通过utils.MaxTimes调用这些变量
	utils.InitConfigs("./config.json")

	//算法生成红包的id和value的对应表
	//初始化redis中envelop_id 和 value的对应表
	//redis需要提供函数func InitEnvelopeValue(values []int)

	r := gin.Default()
	routers.LoadSnatch(r)
	routers.LoadOpen(r)
	routers.LoadWalletList(r)
	r.Run()
}
