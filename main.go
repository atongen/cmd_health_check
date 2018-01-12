package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

// build flags
var (
	Version   string = "development"
	BuildTime string = "unset"
	BuildHash string = "unset"
	GoVersion string = "unset"
)

// cli flags
var (
	cmdFlag     = flag.String("cmd", "", "Health check bash command")
	portFlag    = flag.Int("port", 0, "Port to listen on")
	verboseFlag = flag.Bool("verbose", false, "Print verbose output")
	versionFlag = flag.Bool("version", false, "Print version information and exit")
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func checkCmd() (stdout, stderr string, exitCode int) {
	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command("bash", "-c", *cmdFlag)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err == nil {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
		return
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		ws := exitError.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	} else {
		exitCode = 1
		if stderr == "" {
			stderr = err.Error()
		}
	}

	return
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	stdout, stderr, exitCode := checkCmd()
	if exitCode == 0 {
		// success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	} else {
		// failure
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oh no!"))
		if *verboseFlag {
			logger.Printf("exit code: %d, stdout: %s, stderr: %s\n", exitCode, stdout, stderr)
		}
	}
}

func maintPingHandler(w http.ResponseWriter, r *http.Request) {
	// always success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func versionStr() string {
	return fmt.Sprintf("%s %s %s %s %s", path.Base(os.Args[0]), Version, BuildTime, BuildHash, GoVersion)
}

func main() {
	flag.Parse()

	logger.Println(versionStr())
	if *versionFlag {
		os.Exit(0)
	}

	if cmdFlag == nil || *cmdFlag == "" {
		logger.Println("cmd is required")
		os.Exit(1)
	}

	if *verboseFlag {
		logger.Printf("cmd: %s\n", *cmdFlag)
	}

	if portFlag == nil || *portFlag <= 0 {
		stdout, stderr, exitCode := checkCmd()
		if *verboseFlag {
			logger.Printf("exit code: %d, stdout: %s, stderr: %s\n", exitCode, stdout, stderr)
		}
		os.Exit(exitCode)
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/maint-ping", maintPingHandler)
	http.ListenAndServe(":"+strconv.Itoa(*portFlag), nil)
}
