package main

import (
	"os"
	"log"
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/gargath/prom-slo-recorder/pkg/config"

)

func main() {

	log.Printf("prom-slo-recorder %s\n", version())

//	viper.AutomaticEnv()
//	flag.String("DBAddr", "", "Backend Database host:port; overrides DBADDR")
//	flag.String("DBUser", "", "Backend Database username; overrides DBUSER")
//	flag.String("DBPass", "", "Backend Database password; overrides DBPASS")
//	flag.String("serviceFile", "services.toml", "TOML file to read service definitions from; overrides SERVICEFILE")
//	flag.Bool("help", false, "print this help and exit")
//	flag.Parse()
//	viper.BindPFlags(flag.CommandLine)

	config.ParseFlags()

	if viper.GetBool("help") {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, flag.CommandLine.FlagUsages())
		os.Exit(0)
	}

	config, err := config.LoadConfig(viper.GetString("serviceFile"))
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}
	for _, s := range config.Services {
		fmt.Printf("  Service: %s\n", s.Name)
		fmt.Printf("  Prometheus: %s\n", s.PromAddr)
		fmt.Printf("  Query: %s\n\n", s.Query)
	}

	fmt.Printf("DB: %s:%s@%s\n", viper.GetString("DBUser"), viper.GetString("DBPass"), viper.GetString("DBAddr"))

	os.Exit(0)
}
