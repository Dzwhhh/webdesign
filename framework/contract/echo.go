package contract

const EchoKey = "wd:echo"

type EchoIService interface {
	Echo() string
}
