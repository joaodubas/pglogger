package main

import (
	"time"

	"code.google.com/p/gcfg"
	"github.com/Sirupsen/logrus"
)

const (
	sleepTime  = 2 * time.Second
	configFile = "/etc/pglogger/pglogger.conf"
)

var (
	log    *logrus.Logger
	config Config
)

func init() {
	loadLogging()
	loadConfig()
}

func main() {
	log.Info("Firing up...")
	for {
		work()
		sleep()
	}
}

// loadLogging -- Create a log instance.
func loadLogging() {
	// XXX (mmr) : find a better way to do this
	log = logrus.New()
	fmt := new(logrus.TextFormatter)
	fmt.FullTimestamp = true
	log.Formatter = fmt
	log.Level = logrus.DebugLevel
}

type Config struct {
	AwsCredentials struct {
		Key       string
		SecretKey string
	}
	Graphite struct {
		Host         string
		Port         int
		MetricPrefix string
	}
}

// loadConfig -- Load application conf.
func loadConfig() {
	gcfg.ReadFileInto(&config, configFile)
}

// work -- Start postgres log processing, this pipe contains:
// 1. Fetch logs from rds
// 2. Anaylize log lines
// 3. Send results to graphite
func work() {
	SendLogs(AnalyzeLogs(FetchLogs()))
}

// sleep -- Pause to be used between executions.
func sleep() {
	time.Sleep(sleepTime)
}
