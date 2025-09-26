# Go `net` Package

The [`net`](https://pkg.go.dev/net) package provides portable networking primitives: **dialing** and **listening** for TCP/UDP, **name resolution**, **IP utilities**, and **network interfaces**.

## Common Types and Functions

- `net.Dial`, `net.DialTimeout`, `(&net.Dialer{}).DialContext`: Open client connections.
- `net.Listen`, `net.ListenPacket`: Create servers.
- `net.LookupHost`, `net.LookupIP`: Name resolution.
- `net.JoinHostPort`, `net.SplitHostPort`: Safe host:port formatting and parsing.
- `net.ParseIP`, `net.ParseCIDR`: IP parsing and CIDR network math.
- Deadlines: `SetDeadline`, `SetReadDeadline`, `SetWriteDeadline` on connections.

### `Dial`, `DialTimeout`, and `DialContext`

```go
func Dial(network, address string) (Conn, error)
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)
```

- `Dial` is the simplest form but blocks until connected. 
- `DialTimeout` is a legacy helper for adding a timeout.  
- `DialContext` is the most flexible and recommended, since it supports context cancellation and deadlines.

Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".

Some examples follow:

```go
Dial("tcp", "golang.org:http")
Dial("tcp", "192.0.2.1:http")
Dial("tcp", "198.51.100.1:80")
Dial("udp", "[2001:db8::1]:domain")
Dial("udp", "[fe80::1%lo0]:53")
Dial("tcp", ":80")
Dial("ip4:1", "192.0.2.1")
Dial("ip6:ipv6-icmp", "2001:db8::1")
Dial("ip6:58", "fe80::1%lo0")
```

### `Listen` and `ListenPacket`

```go
func Listen(network, address string) (Listener, error)
func ListenPacket(network, address string) (PacketConn, error)
```

`Listen` creates a server capable of listening for incoming connections. `Listen`'s network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket". `ListenPacket` is similar, but for "udp", "udp4", "udp6", "unixgram", or an IP transport.

See func `Dial` for a description of the network and address parameters.

If the port in the address parameter is empty or "0", as in "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.

### `LookupHost` and `LookupIP`

- `LookupHost` returns a slice of host addresses as strings.
- `LookupIP` returns parsed `net.IP` values.

`LookupHost` uses the local resolver. If you want to specify another one, use `Resolver.LookupHost` instead.

### Signatures of parsing functions

```go
func JoinHostPort(host, port string) string
func SplitHostPort(hostport string) (host, port string, err error)
func ParseIP(s string) IP
func ParseCIDR(s string) (IP, *IPNet, error)
```

---

[Go Back](../../README.md)