package findmyip

import (
	"context"
	"net"
	"time"
)

func FindMyIp() ([]string, error) {
	var r = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			var d = net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", "resolver1.opendns.com:53")
		},
	}

	return r.LookupHost(context.Background(), "myip.opendns.com")
}

func main() {
	ips, err := FindMyIp()
	if err != nil {
		panic(err.Error())
	}
	print(ips[0])
}

