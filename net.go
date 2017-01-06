package stdext

import (
	"fmt"
	"net"
)

// JoinHostPort is wrapper for net.JoinHostPort with the port parameter having
// type int rather than string.
func JoinHostPort(host string, port int) string {
	return net.JoinHostPort(host, fmt.Sprint(port))
}
