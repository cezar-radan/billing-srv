package config

import (
	"fmt"

	"github.com/ledgertech/billing-srv/util/clog"
	"github.com/spf13/viper"
)

var App = AppConfig{
	ServiceName: "billing-srv",
	HostPort:    "localhost:17777",
	TLS:         tlsInfo{},
	Insecure:    false,
	WorkStyle: areaInfo{
		CustomColor: "#85C2EB",
		Lang:        "en-EN",
	},
	Log: logInfo{
		Console: true,
		File:    "",
	},
	Metrics: metricsInfo{
		Enabled:  true,
		Endpoint: "/metrics",
	},
}

// AppConfig represents the configuration file layout
// note the use of "mapstructure" since this is used with viper
type AppConfig struct {
	ServiceName string      `mapstructure:"servicename"`
	HostPort    string      `mapstructure:"hostport"`
	TLS         tlsInfo     `mapstructure:"tls,omitempty"`
	Insecure    bool        `mapstructure:"insecure"`
	WorkStyle   areaInfo    `mapstructure:"workstyle"`
	Log         logInfo     `mapstructure:"log,omitempty"`
	Metrics     metricsInfo `mapstructure:"metrics,omitempty"`
}

type tlsInfo struct {
	Enabled bool   `mapstructure:"enabled,omitempty"`
	Cert    string `mapstructure:"cert,omitempty"`
	Key     string `mapstructure:"key,omitempty"`
}

type areaInfo struct {
	OutputPath    string  `mapstructure:"outputpath"`
	CustomText    string  `mapstructure:"customtext"`
	CustomColor   string  `mapstructure:"customcolor"`
	LogoPath      string  `mapstructure:"logopath"`
	LogoURL       string  `mapstructure:"logourl"`
	WatermarkPath string  `mapstructure:"watermarkpath`
	Lang          string  `mapstructure:"lang"`
	VAT           float64 `mapstructure:"vat"`
}

type logInfo struct {
	Console    bool   `mapstructure:"console,omitempty"`
	File       string `mapstructure:"file,omitempty"`
	MaxSize    int    `mapstructure:"maxSize,omitempty"`
	MaxBackups int    `mapstructure:"maxBackups,omitempty"`
	MaxAge     int    `mapstructure:"maxAge,omitempty"`
	Compress   bool   `mapstructure:"compress,omitempty"`
}

type metricsInfo struct {
	Enabled  bool   `mapstructure:"enabled,omitempty"`
	Endpoint string `mapstructure:"endpoint,omitempty"`
}

func Load(name string, file string, paths ...string) error {
	v := viper.New()
	// Set the profile type
	v.SetConfigType("yaml")

	if file == "" {
		// Set the name of the configuration file to config
		v.SetConfigName(name)
	} else {
		v.SetConfigFile(file)
	}

	v.AutomaticEnv()

	if len(paths) == 0 {
		clog.Info("config paths not configured. assuming current folder")
		paths = append(paths, ".")
	}

	for _, path := range paths {
		v.AddConfigPath(path) // <- // path to look for the config file in
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	clog.Infof("using config file: %s\n", v.ConfigFileUsed())

	return v.Unmarshal(&App)
}
