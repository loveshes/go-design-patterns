package main

import (
	"fmt"
	"github.com/loveshes/go-design-patterns/pattern/proxy-pattern/proxy"
)

func main() {
	girl := proxy.Girl{"小红"}
	trueLove := proxy.TrueLove{"小明"}
	trueLove.SendMsg(girl)
	fmt.Println()

	proxy := proxy.Proxy{Name: "工具人小军"}
	proxy.Wrap(trueLove)
	proxy.SendMsg(girl)
}
