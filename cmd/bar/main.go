package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "http://www.myurl.com?cid=__CID__&imei=__IMEI__&os=__OS__&timestamp=__TS__&plan =__PLAN__&unit=__UNIT__&callback_url=__CALLBACK_URL__"
	fmt.Printf("%v", strings.Contains(str, "__CALLBACK_URL__"))
}
