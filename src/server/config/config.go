package config

const (
  EnvDev = iota
  EnvStage
  EnvProd
)

type ConfigObject struct {
  Host string
  Port int
  Environment int
}
