package log

import (
	"fmt"
	"path"

	"github.com/cihub/seelog"
)

func InitLog(loglevel string, logfile string) int {
	logConf := `
<seelog minlevel="` + loglevel + `">
    <outputs formatid="main">
        <filter levels="critical">
            <file path="` + path.Dir(logfile) + `/stack.log"/>
        </filter>
        <filter levels="trace, debug, info, warn, error">
`
	logConf = logConf + `
            <rollingfile type="date" filename="` + logfile + `" datepattern="02.01.2006" maxrolls="3"/>
        </filter>
    </outputs>
    <formats>
        <format id="main" format="%Date(2006-01-02T15:04:05.999999999Z07:00) [@%File.%Line][%LEV] %Msg%n"/>
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(logConf))
	if err != nil {
		fmt.Println("err parsing config:", err)
		return -1
	}

	seelog.ReplaceLogger(logger)
	return 0
}
