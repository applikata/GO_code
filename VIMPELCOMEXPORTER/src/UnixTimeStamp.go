package src

import "time"

func timestatpBuild(period int) int64 {

	unixTimeStamp := time.Now().Unix()

	unixTimeStamp = unixTimeStamp - (int64(period) * 60)

	return unixTimeStamp

}
