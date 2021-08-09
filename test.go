import (
  "fmt"
  "github.com/kataras/golog"
  "github.com/go-resty/resty/v2"
)

const (
  root = "vxrail"
)

var (
  logger = golog.Child(root).setPrefix(fmt.Sprintf("[api:%s]", root))
  restClient = resty.New()
)
