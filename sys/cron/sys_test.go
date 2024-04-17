package cron_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zmloong/lotus/sys/cron"
)

// 1. cron表达式格式：
// {秒数} {分钟} {小时} {日期} {月份} {星期} {年份(可为空)}
// 2. cron表达式各占位符解释：
// {秒数} ==> 允许值范围: 0~59 ,不允许为空值，若值不合法，调度器将抛出SchedulerException异常
// "*" 代表每隔1秒钟触发；
// "," 代表在指定的秒数触发，比如"0,15,45"代表0秒、15秒和45秒时触发任务
// "-" 代表在指定的范围内触发，比如"25-45"代表从25秒开始触发到45秒结束触发，每隔1秒触发1次
// "/" 代表触发步进(step)，"/"前面的值代表初始值("*"等同"0")，后面的值代表偏移量，比如"0/20"或者"*/20"代表从0秒钟开始，每隔20秒钟触发1次，即0秒触发1次，20秒触发1次，40秒触发1次；"5/20"代表5秒触发1次，25秒触发1次，45秒触发1次；"10-45/20"代表在[10,45]内步进20秒命中的时间点触发，即10秒触发1次，30秒触发1次
// {分钟} ==> 允许值范围: 0~59 ,不允许为空值，若值不合法，调度器将抛出SchedulerException异常
// "*" 代表每隔1分钟触发；
// "," 代表在指定的分钟触发，比如"10,20,40"代表10分钟、20分钟和40分钟时触发任务
// "-" 代表在指定的范围内触发，比如"5-30"代表从5分钟开始触发到30分钟结束触 发，每隔1分钟触发
// "/" 代表触发步进(step)，"/"前面的值代表初始值("*"等同"0")，后面的值代表偏移量，比如"0/25"或者"*/25"代表从0分钟开始，每隔25分钟触发1次，即0分钟触发1次，第25分钟触发1次，第50分钟触发1次；"5/25"代表5分钟触发1次，30分钟触发1次，55分钟触发1次；"10-45/20"代表在[10,45]内步进20分钟命中的时间点触发，即10分钟触发1次，30分钟触发1次
// {小时} ==> 允许值范围: 0~23 ,不允许为空值，若值不合法，调度器将抛出SchedulerException异常
// "*" 代表每隔1小时触发；
// "," 代表在指定的时间点触发，比如"10,20,23"代表10点钟、20点钟和23点触发任务
// "-" 代表在指定的时间段内触发，比如"20-23"代表从20点开始触发到23点结束触发，每隔1小时触发
// "/" 代表触发步进(step)，"/"前面的值代表初始值("*"等同"0")，后面的值代表偏移量，比如"0/1"或者"*/1"代表从0点开始触发，每隔1小时触发1次；"1/2"代表从1点开始触发，以后每隔2小时触发一次；"19-20/2"表达式将只在19点触发
// {日期} ==> 允许值范围: 1~31 ,不允许为空值，若值不合法，调度器将抛出SchedulerException异常
// "*" 代表每天触发；
// "?" 与{星期}互斥，即意味着若明确指定{星期}触发，则表示{日期}无意义，以免引起 冲突和混乱
// "," 代表在指定的日期触发，比如"1,10,20"代表1号、10号和20号这3天触发
// "-" 代表在指定的日期范围内触发，比如"10-15"代表从10号开始触发到15号结束触发，每隔1天触发
// "/" 代表触发步进(step)，"/"前面的值代表初始值("*"等同"1")，后面的值代表偏移量，比如"1/5"或者"*/5"代表从1号开始触发，每隔5天触发1次；"10/5"代表从10号开始触发，以后每隔5天触发一次；"1-10/2"表达式意味着在[1,10]范围内，每隔2天触发，即1号，3号，5号，7号，9号触发
// "L" 如果{日期}占位符如果是"L"，即意味着当月的最后一天触发
// "W "意味着在本月内离当天最近的工作日触发，所谓最近工作日，即当天到工作日的前后最短距离，如果当天即为工作日，则距离为0；所谓本月内的说法，就是不能跨月取到最近工作日，即使前/后月份的最后一天/第一天确实满足最近工作日；因此，"LW"则意味着本月的最后一个工作日触发，"W"强烈依赖{月份}
// "C" 根据日历触发，由于使用较少，暂时不做解释
// {月份} ==> 允许值范围: 1~12 (JAN-DEC),不允许为空值，若值不合法，调度器将抛出SchedulerException异常
// "*" 代表每个月都触发；
// "," 代表在指定的月份触发，比如"1,6,12"代表1月份、6月份和12月份触发任务
// "-" 代表在指定的月份范围内触发，比如"1-6"代表从1月份开始触发到6月份结束触发，每隔1个月触发
// "/" 代表触发步进(step)，"/"前面的值代表初始值("*"等同"1")，后面的值代表偏移量，比如"1/2"或者"*/2"代表从1月份开始触发，每隔2个月触发1次；"6/6"代表从6月份开始触发，以后每隔6个月触发一次；"1-6/12"表达式意味着每年1月份触发
// {星期} ==> 允许值范围: 1~7 (SUN-SAT),1代表星期天(一星期的第一天)，以此类推，7代表星期六(一星期的最后一天)，不允许为空值，若值不合法，调度器将抛出SchedulerException异常
// "*" 代表每星期都触发；
// "?" 与{日期}互斥，即意味着若明确指定{日期}触发，则表示{星期}无意义，以免引起冲突和混乱
// "," 代表在指定的星期约定触发，比如"1,3,5"代表星期天、星期二和星期四触发
// "-" 代表在指定的星期范围内触发，比如"2-4"代表从星期一开始触发到星期三结束触发，每隔1天触发
// "/" 代表触发步进(step)，"/"前面的值代表初始值("*"等同"1")，后面的值代表偏移量，比如"1/3"或者"*/3"代表从星期天开始触发，每隔3天触发1次；"1-5/2"表达式意味着在[1,5]范围内，每隔2天触发，即星期天、星期二、星期四触发
// "L" 如果{星期}占位符如果是"L"，即意味着星期的的最后一天触发，即星期六触发，L= 7或者 L = SAT，因此，"5L"意味着一个月的最后一个星期四触发
// "#" 用来指定具体的周数，"#"前面代表星期，"#"后面代表本月第几周，比如"2#2"表示本月第二周的星期一，"5#3"表示本月第三周的星期四，因此，"5L"这种形式只不过是"#"的特殊形式而已
// "C" 根据日历触发，由于使用较少，暂时不做解释
// {年份} ==> 允许值范围: 1970~2099 ,允许为空，若值不合法，调度器将抛出SchedulerException异常
// "*"代表每年都触发；
// ","代表在指定的年份才触发，比如"2011,2012,2013"代表2011年、2012年和2013年触发任务
// "-"代表在指定的年份范围内触发，比如"2011-2020"代表从2011年开始触发到2020年结束触发，每隔1年触发
// "/"代表触发步进(step)，"/"前面的值代表初始值("*"等同"1970")，后面的值代表偏移量，比如"2011/2"或者"*/2"代表从2011年开始触发，每隔2年触发1次
// 注意：除了{日期}和{星期}可以使用"?"来实现互斥，表达无意义的信息之外，其他占位符都要具有具体的时间含义，且依赖关系为：年->月->日期(星期)->小时->分钟->秒数
// 3. cron表达式的强大魅力在于灵活的横向和纵向组合以及简单的语法，用cron表达式几乎可以写出任何你想要触发的时间点
// 经典案例：
// "30 * * * * ?" 每半分钟触发任务
// "30 10 * * * ?" 每小时的10分30秒触发任务
// "30 10 1 * * ?" 每天1点10分30秒触发任务
// "30 10 1 20 * ?" 每月20号1点10分30秒触发任务
// "30 10 1 20 10 ? *" 每年10月20号1点10分30秒触发任务
// "30 10 1 20 10 ? 2011" 2011年10月20号1点10分30秒触发任务
// "30 10 1 ? 10 * 2011" 2011年10月每天1点10分30秒触发任务
// "30 10 1 ? 10 SUN 2011" 2011年10月每周日1点10分30秒触发任务
// "15,30,45 * * * * ?" 每15秒，30秒，45秒时触发任务
// "15-45 * * * * ?" 15到45秒内，每秒都触发任务
// "15/5 * * * * ?" 每分钟的每15秒开始触发，每隔5秒触发一次
// "15-30/5 * * * * ?" 每分钟的15秒到30秒之间开始触发，每隔5秒触发一次
// "0 0/3 * * * ?" 每小时的第0分0秒开始，每三分钟触发一次
// "0 15 10 ? * MON-FRI" 星期一到星期五的10点15分0秒触发任务
// "0 15 10 L * ?" 每个月最后一天的10点15分0秒触发任务
// "0 15 10 LW * ?" 每个月最后一个工作日的10点15分0秒触发任务
// "0 15 10 ? * 5L" 每个月最后一个星期四的10点15分0秒触发任务
// "0 15 10 ? * 5#3" 每个月第三周的星期四的10点15分0秒触发任务

// 一些cron表达式案例
// */5 * * * * ? 每隔5秒执行一次
// 0 */1 * * * ? 每隔1分钟执行一次
// 0 0 5-15 * * ? 每天5-15点整点触发
// 0 0/3 * * * ? 每三分钟触发一次
// 0 0-5 14 * * ? 在每天下午2点到下午2:05期间的每1分钟触发
// 0 0/5 14 * * ? 在每天下午2点到下午2:55期间的每5分钟触发
// 0 0/5 14,18 * * ? 在每天下午2点到2:55期间和下午6点到6:55期间的每5分钟触发
// 0 0/30 9-17 * * ? 朝九晚五工作时间内每半小时
// 0 0 10,14,16 * * ? 每天上午10点，下午2点，4点

// 0 0 12 ? * WED 表示每个星期三中午12点
// 0 0 17 ? * TUES,THUR,SAT 每周二、四、六下午五点
// 0 10,44 14 ? 3 WED 每年三月的星期三的下午2:10和2:44触发
// 0 15 10 ? * MON-FRI 周一至周五的上午10:15触发
// 0 0 23 L * ? 每月最后一天23点执行一次
// 0 15 10 L * ? 每月最后一日的上午10:15触发
// 0 15 10 ? * 6L 每月的最后一个星期五上午10:15触发
// 0 15 10 * * ? 2005 2005年的每天上午10:15触发
// 0 15 10 ? * 6L 2002-2005 2002年至2005年的每月的最后一个星期五上午10:15触发
// 0 15 10 ? * 6#3 每月的第三个星期五上午10:15触发

func Test_sys(t *testing.T) {
	if err := cron.OnInit(nil); err == nil {
		// AddFunc("@every 1s", func() { //每一秒
		// 	fmt.Printf("@every 1s")
		// })
		// AddFunc("@every 1m", func() { //每一分
		// 	fmt.Printf("@every 1m")
		// })
		// AddFunc("@hourly", func() { //每一小时
		// 	fmt.Printf("@hourly")
		// })
		// AddFunc("@daily", func() { //每天凌晨
		// 	fmt.Printf("@daily")
		// })
		// AddFunc("@daily", func() { //每天凌晨
		// 	fmt.Printf("@daily")
		// })
		// AddFunc("*/5 * * * * ?", func() { //每天凌晨
		// 	fmt.Printf("*/5 * * * * ?")
		// })
		cron.AddFunc("30/59 0/2 * * * ?", func() { //每隔90秒
			fmt.Printf("*/5 * * * * ?")
		})

	}
	time.Sleep(time.Minute * 5)
}
