package initialization

import (
	"gin-IM/internal/conf"
	"gin-IM/pkg/tools/snowflake"
)

func init() {
	// 初始化雪花算法工具
	snowflake.SnowflakeInit(conf.Conf.MachineID)
}
