package util

import (
	"os"
	"fmt"
	"runtime"
	 "encoding/json"
        "io/ioutil"
        "os/signal"
        "syscall"
	log "github.com/golang/glog"
)

var CVL_SCHEMA string = "schema/"
var CVL_CFG_FILE string = "./logconf.json"

//package init function 
func init() {
	if (os.Getenv("CVL_SCHEMA_PATH") != "") {
		CVL_SCHEMA = os.Getenv("CVL_SCHEMA_PATH") + "/"
	}

	if (os.Getenv("CVL_CFG_FILE") != "") {
		CVL_CFG_FILE = os.Getenv("CVL_CFG_FILE")
	}
}

var cvlCfgMap map[string]bool

/* Logging Level for CVL global logging. */
type CVLLogLevel uint8 
const (
        INFO  = 0 + iota
        WARNING
        ERROR
        FATAL
        INFO_API
	INFO_TRACE
	INFO_DEBUG
	INFO_DATA
	INFO_DETAIL
	INFO_ALL
)

/* Logging levels for CVL Tracing. */
type CVLTraceLevel uint8 
const (
        TRACE_CACHE  = 1 << 0
        TRACE_LIBYANG = 1 << 1
        TRACE_YPARSER = 1 << 2
        TRACE_CREATE = 1 << 3
        TRACE_UPDATE = 1 << 4
        TRACE_DELETE = 1 << 5
        TRACE_SEMANTIC = 1 << 6
        TRACE_SYNTAX = 1 << 7

)

const (
	TRACE_MIN = 0
	TRACE_MAX = 7
)


var traceLevelMap = map[int]string {
	/* Caching operation traces */
	TRACE_CACHE : "TRACE_CACHE",
	/* Libyang library traces. */
	TRACE_LIBYANG: "TRACE_LIBYANG",
	/* Yang Parser traces. */
	TRACE_YPARSER : "TRACE_YPARSER", 
	/* Create operation traces. */
	TRACE_CREATE : "TRACE_CREATE", 
	/* Update operation traces. */
	TRACE_UPDATE : "TRACE_UPDATE", 
	/* Delete operation traces. */
	TRACE_DELETE : "TRACE_DELETE", 
	/* Semantic Validation traces. */
	TRACE_SEMANTIC : "TRACE_SEMANTIC",
	/* Syntax Validation traces. */
	TRACE_SYNTAX : "TRACE_SYNTAX", 
}

var Tracing bool = false

var traceFlags uint16 = 0

func SetTrace(on bool) {
	if (on == true) {
		traceFlags = 1
	} else {
		traceFlags = 0
	}
}

func IsTraceSet() bool {
	if (traceFlags == 0) {
		return false
	} else {
		return true
	}
}

func IsLogTraceSet() bool {
	return true
}

func SetTraceLevel(level uint8) {
	traceFlags = traceFlags | (1 << level)
}


func ClearTraceLevel(level uint8) {
	traceFlags = traceFlags &^ (1 << level)
}

func TRACE_LOG(level log.Level, tracelevel CVLTraceLevel, fmtStr string, args ...interface{}) {
	if (IsTraceSet() == false) {
		return
	}

	traceEnabled := false
	var index uint8
	/* Check if incoming tracelevel is enabled. */
	for  index = TRACE_MIN ; index < TRACE_MAX ; index++  {
		if ((tracelevel &  (1 << index)) != 0) {
			if (cvlCfgMap[traceLevelMap[1 << index]] == true) {
				/* This log should be allowed as at least one flag is currently set */
				traceEnabled = true
				break
			}
		}
	}


	if IsTraceSet() == true && traceEnabled == true {
		pc := make([]uintptr, 10)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])

		fmt.Printf("%s:%d %s(): ", file, line, f.Name())
		fmt.Printf(fmtStr+"\n", args...)
	} else {
		log.V(level).Infof(fmtStr, args...)
	}
}

func CVL_LOG(level CVLLogLevel, format string, args ...interface{}) {

	switch level {
		case INFO:
		       log.Infof(format, args...)
		case  WARNING:
		       log.Warningf(format, args...)
		case  ERROR:
		       log.Errorf(format, args...)
		case  FATAL:
		       log.Fatalf(format, args...)
		case INFO_API:
			log.V(1).Infof(format, args...)
		case INFO_TRACE:
			log.V(2).Infof(format, args...)
		case INFO_DEBUG:
			log.V(3).Infof(format, args...)
		case INFO_DATA:
			log.V(4).Infof(format, args...)
		case INFO_DETAIL:
			log.V(5).Infof(format, args...)
		case INFO_ALL:
			log.V(6).Infof(format, args...)
	}	

}


func ConfigFileSyncHandler() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR2)
	go func() {
		for {
			<-sigs
			cvlCfgMap := ReadConfFile()
			CVL_LOG(INFO ,"Received SIGUSR2. Changed configuration values are %v", cvlCfgMap)
		}
	}()

}

func ReadConfFile()  map[string]bool{
        data, err := ioutil.ReadFile(CVL_CFG_FILE)

        err = json.Unmarshal(data, &cvlCfgMap)

        if err != nil {
		CVL_LOG(ERROR ,"Error in reading cvl configuration file %v", err)
        }

	CVL_LOG(INFO ,"Current Values of CVL Configuration File %v", cvlCfgMap)
	return cvlCfgMap
}


