package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

// 定义一个节点： 通过这个全局的 node，就可以用于制造 ID了。
var node *snowflake.Node

// 源码里有很多默认的基本设定，比如开始时间等等， 我们可以改变这些，自己初始化一个 node节点
func Init(startTime string, machineID int64) {
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}
	// 设置时间
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		panic(err)
	}
	zap.L().Info("snowflake node init success")
}

// 返回int64位的 id值
func GenID() int64 {
	return node.Generate().Int64()
}

func main() {
	Init("2023-01-01", 1)
	println(GenID())
}
