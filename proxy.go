package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Proxy struct contains proxy parameters and traffic statistics.
type Proxy struct {
	listen  string
	forward string
	stats   *Stats
}

// NewProxy constructs new Proxy object.
func NewProxy(from, to string) *Proxy {
	return &Proxy{
		listen:  from,
		forward: to,
		stats:   NewStats(),
	}
}

// Run proxy loop and process incoming system signals.
func (p *Proxy) Run() error {
	listener, err := net.Listen("tcp", p.listen)
	if err != nil {
		return err
	}

	go p.run(listener)

	p.processSignals()
	return nil
}

// run proxy loop waiting for incoming connections.
func (p *Proxy) run(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		} else {
			go p.handle(conn)
		}
	}
}

// processSignals runs infinite loop checking for system signals.
func (p *Proxy) processSignals() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGUSR2)
	for {
		signalType := <-signals
		if signalType == syscall.SIGUSR2 {
			p.stats.Print()
		}
	}
}

// handle single incoming connection.
func (p *Proxy) handle(connection net.Conn) {
	defer connection.Close()
	remote, err := net.Dial("tcp", p.forward)
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()

	wait := &sync.WaitGroup{}
	wait.Add(2)
	go p.copy(remote, connection, wait)
	go p.copy(connection, remote, wait)
	wait.Wait()
}

// copy message through proxy.
func (p *Proxy) copy(listen, forward net.Conn, wait *sync.WaitGroup) {
	defer wait.Done()

	scanner := bufio.NewScanner(listen)
	for scanner.Scan() {
		message := scanner.Text()
		p.stats.Update(message)
		strings.NewReader(message + "\n").WriteTo(forward)
	}
}
