package routers

import (
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"red_envelop_server/sql"
	"red_envelop_server/utils"
	"strconv"
)


func LoadSnatch(e *gin.Engine) {
	e.POST("/snatch", SnatchHandler)
}



func SnatchHandler(c *gin.Context) {
	//每个人能抢的最大红包数应该从配置文件读进来
	max_count := utils.MaxTimes

	uid, _ := c.GetPostForm("uid")
	logs.Printf("%s is snatching envelope", uid)
	int_uid, _ := strconv.ParseInt(uid, 10, 64)

	////查缓存表
	//func CheckUserInRedis() bool {
	//	return false/true
	//}

	//if(CheckUserInRedis() == false) {
		//在数据库中创建一条User记录，id为int_uid，count为0
		//sql.CreateUser(int_uid)--------------需补充
		//在redis中创建一条User记录，id为int_uid，count为1, 返回用户本次抢到的红包id
		//func CreateUserInRedis(int_uid) int64

	//}else {
		//func FindEnvelopeIdSnatchedByUserInRedis(int_uid) int64

	//}

	//if FindEnvelopeIdSnatchedByUserInRedis函数的返回值 == -1 {
		//		flag = false
		//		c.JSON(200, gin.H{
		//			"code": -1,
		//			"msg":  "fail! snatch too much",
		//			"data": gin.H{
		//				"max_count": max_count,
		//				"cur_count": user.Count,
		//			},
		//		})
		//		return
	//}

	////根据概率计算用户这次应不应该拿到红包，这里我想的是对所有的请求做统一的处理，直接放弃一部分请求不处理，
	////这样既满足了概率也减轻了后端的压力,只处理十分之一的请求
	//rand_num := rand.Intn(10)
	//if rand_num != 0 {
	//	flag = false
	//	c.JSON(200, gin.H{
	//		"code": -3,
	//		"msg":  "According to the probability, the red envelope can not be snatched this time",
	//	})
	//	return
	//}

	//如果上述的条件都满足了，并且概率正好也轮到了，为当前用户生成红包
	//if flag {
		//sql.CreateEnvelope(int_uid) int64----------需要修改原有代码
		//这个函数中需要根据int_uid在缓存中查找value, 并获取当前时间，设置opened = false在数据库插入一条红包表记录，
		//envelope_id自增长的，创建时不需要赋值,创建后要返回envelope_id回来

		//这个函数中还需要在redis中插入如下格式的数据
		//{key=红包id, value=[envelope_id1, envelope_id2, envelope_id3, ...]}
		//{key=红包id, value=[snatch_time1, snatch_time2, snatch_time3, ...]}
		//对应redis，提供一个函数，失效钱包列表缓存也需要函数中在这个执行
		//func AddEnvelopForUser(uid, envelope_id, snatch_time)

		//返回阶段
		//"envelope_id": envelope.ID,
		//"max_count":   max_count,
		//"cur_count":   redis提供一个返回用户当前次数的方法 func GetUserCurCount(uid int64) int64,

	//}

	//根据uid查询用户，没有的话就创建用户
	user := sql.GetUser(int_uid)
	flag := true

	//判断用户的count是否大于个人最多抢红包数
	if user.Count >= int64(max_count) {
		flag = false
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "fail! snatch too much",
			"data": gin.H{
				"max_count": max_count,
				"cur_count": user.Count,
			},
		})
		return
	}

	////根据概率计算用户这次应不应该拿到红包，这里我想的是对所有的请求做统一的处理，直接放弃一部分请求不处理，
	////这样既满足了概率也减轻了后端的压力,只处理十分之一的请求
	//rand_num := rand.Intn(10)
	//if rand_num != 0 {
	//	flag = false
	//	c.JSON(200, gin.H{
	//		"code": -3,
	//		"msg":  "According to the probability, the red envelope can not be snatched this time",
	//	})
	//	return
	//}

	if flag {
		//如果上述的条件都满足了，并且概率正好也轮到了，为当前用户生成红包
		envelope := sql.CreateEnvelope(user)

		//修改当前用户的count
		sql.UpdateCount(&user)

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"envelope_id": envelope.ID,
				"max_count":   max_count,
				"cur_count":   user.Count,
			},
		})
	}
}
