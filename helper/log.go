package helper

import (
	"fmt"
	"time"
	"os"
)

// Log is the default Logger object.
var Log = Logger{
	LogPath:    "default.log",
	DebugLevel: DebugLevelLogEverything,
}

const (
	DebugLevelLogEverything = 10
	DebugLevelLogErrorsOnly = 5
	DebugLevelLogNothing    = 0
)


// Logger is a lightweight struct to provide logging due runtime.
type Logger struct {
	LogPath    string
	DebugLevel uint8
}

// LogInfo logs Debug Information on harddrive and stdout.
func LogInfo(msg string){

	if Log.DebugLevel < DebugLevelLogEverything {
		return
	}

	msg = time.Now().String() + " - " + "Info: " + msg
	fmt.Println(msg)
	writeToLogFile(msg)
}

// LogError logs Error Information on harddrive and stdout.
func LogError(msg string){

	if Log.DebugLevel < DebugLevelLogErrorsOnly {
		return
	}

	msg = time.Now().String() + " - " + "ERROR: " + msg
	fmt.Println(msg)
	writeToLogFile(msg)
}

// writeTiLogFile appends a given string as an newline to the logfile, which is stored in the Log object.
// if no logfile exists it will be created.
func writeToLogFile(msg string){
	logFile, err := os.OpenFile(Log.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {

		fmt.Println("FATAL ERROR: could't write to logfile at location '" + Log.LogPath + "' because " + err.Error())
		return
	}

	defer logFile.Close()

	if _, err = logFile.WriteString(msg + "\n"); err != nil {
		fmt.Println("FATAL ERROR: could't write to logfile at location '" + Log.LogPath + "' because" + err.Error())
	}
}