package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// build flags
var (
	Version   string = "development"
	BuildTime string = "unset"
	BuildHash string = "unset"
	GoVersion string = "unset"
)

const cmd string = "ps aux | egrep 'sidekiq [0-9]+\\.[0-9]+\\.[0-9]+ .+busy]' | wc -l"

// cli flags
var (
	portFlag    = flag.Int("port", 0, "Port to listen on")
	versionFlag = flag.Bool("v", false, "Print version information and exit")
	numFlag     = flag.Int("num", runtime.NumCPU(), "Number of sidekiq processes needed to be health")
)

func checkCmd() bool {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute command: %s\n", err)
		return false
	}

	i, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed count sidekiq processes: %s\n", err)
		return false
	}

	return i == *numFlag
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if checkCmd() {
		// success
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	} else {
		// failure
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oh no!"))
	}
}

func maintPingHandler(w http.ResponseWriter, r *http.Request) {
	// always success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s %s %s %s %s\n", path.Base(os.Args[0]), Version, BuildTime, BuildHash, GoVersion)
		os.Exit(0)
	}

	if *portFlag <= 0 {
		result := checkCmd()
		fmt.Printf("Sidekiq healthy?: %t\n", result)
		os.Exit(0)
	}

	if *numFlag <= 0 {
		fmt.Fprintf(os.Stderr, "Invalid check number: %d\n", *numFlag)
		os.Exit(1)
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/maint-ping", maintPingHandler)
	http.ListenAndServe(":"+strconv.Itoa(*portFlag), nil)
}
