package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"time"
)

/*
demoTCP: echo server + client on 127.0.0.1:0 (ephemeral port).
An ephemeral port is a temporary, automatically assigned network port on a client device used for
the duration of a specific communication session.
*/
func demoTCP() error {
	// Create the TCP server.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer func() {
		if err := ln.Close(); err != nil {
			fmt.Println("listener close error:", err)
		}
	}()

	// Get the assigned address (with ephemeral port).
	addr := ln.Addr().String()
	fmt.Println("listening on", addr)

	// Server: accept a single conn and echo once.
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			return
		}
		defer func() {
			if err := conn.Close(); err != nil {
				fmt.Println("server conn close error:", err)
			}
		}()
		_ = conn.SetDeadline(time.Now().Add(2 * time.Second))

		// Simple echo.
		if _, err := io.Copy(conn, conn); err != nil {
			fmt.Println("echo error:", err)
			return
		}
	}()

	// Client connects.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("client conn close error:", err)
		}
	}()
	_ = conn.SetDeadline(time.Now().Add(2 * time.Second))

	// Client sends message.
	msg := "hello over tcp"
	if _, err := io.WriteString(conn, msg); err != nil {
		return err
	}

	// Client reads echoed message.
	buf := make([]byte, 64)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("client got: %q\n", string(buf[:n]))

	// Allow server goroutine to finish.
	<-done
	return nil
}

/*
demoUDP: echo using ListenPacket + Dial("udp", ...).
*/
func demoUDP() error {
	// Server uses ListenPacket to get a PacketConn.
	pc, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer func() {
		if err := pc.Close(); err != nil {
			fmt.Println("packet conn close error:", err)
		}
	}()

	// Get the assigned address (with ephemeral port).
	addr := pc.LocalAddr().String()
	fmt.Println("listening (udp) on", addr)

	// Server: read a single packet, convert to upper case, send back.
	done := make(chan struct{})
	go func() {
		defer close(done)
		_ = pc.SetDeadline(time.Now().Add(2 * time.Second))
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			fmt.Println("udp read error:", err)
			return
		}

		// Echo back in upper case.
		resp := bytes.ToUpper(buf[:n])
		if _, err := pc.WriteTo(resp, addr); err != nil {
			fmt.Println("udp write error:", err)
		}
	}()

	// Client uses net.Dial with UDP to get a connected UDPConn-like API.
	c, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer func() {
		if err := c.Close(); err != nil {
			fmt.Println("client conn close error:", err)
		}
	}()
	_ = c.SetDeadline(time.Now().Add(2 * time.Second))

	// Client sends message.
	if _, err := io.WriteString(c, "ping"); err != nil {
		return err
	}

	// Client reads echoed message.
	buf := make([]byte, 16)
	n, err := c.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("client got: %q\n", string(buf[:n]))

	// Allow server goroutine to finish.
	<-done
	return nil
}

/*
demoDNS: DNS over localhost, i.e. works offline.
*/
func demoDNS() error {
	// LookupHost returns both IPv4 and IPv6 addresses.
	hosts, err := net.LookupHost("localhost")
	if err != nil {
		return err
	}
	fmt.Println("LookupHost localhost ->", hosts)

	// LookupIP returns a slice of net.IP, which can be filtered for v4/v6.
	ips, err := net.LookupIP("localhost")
	if err != nil {
		return err
	}
	var v4, v6 []string
	for _, ip := range ips {
		if ip.To4() != nil {
			v4 = append(v4, ip.String())
		} else {
			v6 = append(v6, ip.String())
		}
	}
	fmt.Println("LookupIP localhost -> IPv4:", v4, "IPv6:", v6)
	return nil
}

/*
demoIPandPorts shows the usage of IP utilities + Join/Split for host:port.
*/
func demoIPandPorts() error {
	// ParseIP returns nil if the string is not a valid textual representation of an IP address.
	ip := net.ParseIP("192.0.2.1")
	fmt.Println("ParseIP 192.0.2.1 ->", ip)

	// ParseCIDR parses a CIDR notation IP address and prefix length.
	_, cidr, err := net.ParseCIDR("2001:db8::/32")
	if err != nil {
		return err
	}
	fmt.Println("ParseCIDR 2001:db8::/32 -> network:", cidr)

	// Contains reports whether the network contains the given IP address.
	fmt.Println("Contains 2001:db8::1 ?", cidr.Contains(net.ParseIP("2001:db8::1")))

	// JoinHostPort combines a host and port into a network address of the form "host:port".
	host, port := "127.0.0.1", "8080"
	hp := net.JoinHostPort(host, port)
	fmt.Println("JoinHostPort ->", hp)

	// SplitHostPort splits a network address of the form "host:port".
	h2, p2, err := net.SplitHostPort(hp)
	if err != nil {
		return err
	}
	fmt.Printf("SplitHostPort %q -> host=%q port=%q\n", hp, h2, p2)
	return nil
}

func main() {
	fmt.Println("== net.Dial / net.Listen (TCP echo on loopback) ==")
	if err := demoTCP(); err != nil {
		fmt.Println("TCP demo error:", err)
	}
	fmt.Println()

	fmt.Println("== net.ListenPacket / net.Dial (UDP echo on loopback) ==")
	if err := demoUDP(); err != nil {
		fmt.Println("UDP demo error:", err)
	}
	fmt.Println()

	fmt.Println("== Name resolution (LookupHost / LookupIP) ==")
	if err := demoDNS(); err != nil {
		fmt.Println("DNS demo error:", err)
	}
	fmt.Println()

	fmt.Println("== IP utilities (ParseIP / ParseCIDR / Join+Split host:port) ==")
	if err := demoIPandPorts(); err != nil {
		fmt.Println("IP/CIDR demo error:", err)
	}
	fmt.Println()
}
