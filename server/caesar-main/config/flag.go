package config

import "flag"

var (
	// Debug If this program runs in debug mode.
	Debug = flag.Bool("debug", false, `Set debug mode for more runtime information.
This option may be overwritten by config.`)
	// ConfigFile Config File Path.
  ConfigFile = flag.String("ConfigFile", "caesar-main.yaml", "Load config from this file.")
  help = flag.Bool("help", false, "Show help.")
)

func init() {
  if *help {
    flag.Usage()
  }
}
