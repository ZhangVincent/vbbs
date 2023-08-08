package iplocator_test

import (
	"fmt"
	"testing"
	"vbbs/pkg/iplocator"
)

func TestSearch(t *testing.T) {
	iplocator.InitIpLocator("/data/ip2region.xdb")
	ip := "47.52.26.78"
	fmt.Println(iplocator.Search(ip))
	fmt.Println(iplocator.IpLocation(ip))
}
