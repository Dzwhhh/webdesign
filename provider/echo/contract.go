package echo

const Key = "wd:echo"

type Service interface {
	Echo() string
}
