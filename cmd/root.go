package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ledgertech/billing-srv/config"
	"github.com/ledgertech/billing-srv/util/clog"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/cobra"
)

var Version = "development"
var AppName = "billing-srv"

var cfgFile *string
var logLevel *string

var rootCmd = &cobra.Command{
	Use:     AppName,
	Version: Version,
	Short:   AppName + " is an utility app",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cfgFile = rootCmd.PersistentFlags().String("config", "", "config file (default is searched in . and $HOME/.conf/"+AppName+".yaml)")
	logLevel = rootCmd.PersistentFlags().String("loglevel", "WARN", "log levels NONE, ERROR, INFO, WARN, DEBUG, VERBOSE (default is WARN)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	clog.Infof("%s: %s\n", AppName, Version)

	home, err := os.UserHomeDir()
	if err != nil {
		clog.Error(err)
		return
	}

	confPath := filepath.Join(home, ".conf")

	err = config.Load(AppName, *cfgFile, ".", confPath)
	if err != nil {
		clog.Errorf("could not load config: %s", err)
		return
	}

	clog.Infof("active configuration: %+v\n", config.App)

	level, err := clog.FindLevel(strings.ToUpper(*logLevel))
	if err != nil {
		level = clog.WarnLevel
	}
	clog.SetLevel(level)

	logFile := false
	if config.App.Log.File != "" {
		fileLogger := &lumberjack.Logger{
			Filename:   config.App.Log.File,
			MaxSize:    config.App.Log.MaxSize,    // megabytes
			MaxBackups: config.App.Log.MaxBackups, //number of files to keep
			MaxAge:     config.App.Log.MaxAge,     //days
			Compress:   config.App.Log.Compress,   // disabled by default
		}
		if config.App.Log.Console {
			mw := io.MultiWriter(os.Stdout, fileLogger)
			log.SetOutput(mw)
		} else {
			clog.Infof("switching log to file: %s\n", config.App.Log.File)
			log.SetOutput(fileLogger)
		}
		logFile = true
	}
	if !logFile {
		log.SetOutput(ioutil.Discard)
	}

	if config.App.Insecure {
		clog.Info("disabling certificate verification.")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
}
