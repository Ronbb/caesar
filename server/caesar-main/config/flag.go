package config

import "flag"

var (
  debug bool
)

func init()  {
  flag.BoolVar(&debug, "debug", true, "Set debug mode for more runtime information.")
}
