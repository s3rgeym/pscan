package internal

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"pscan/pkg/utils"
	"regexp"
	"strconv"
	"sync"
)

var octetPattern = `\d{1,3}`
var ipPattern = octetPattern + `\.` + octetPattern + `\.` + octetPattern + `\.` + octetPattern

// x.x.x.x-x.x.x.x OR  x.x.x.x/x
var ipRangeRegex, _ = regexp.Compile(`^` + ipPattern + `(?:-` + ipPattern + `|/\d+)$`)

// Run fn
func Run(opts *Options, args []string) error {
	if opts.Input != "" {
		filename, err := utils.ExpandPath(string(opts.Input))
		if err != nil {
			return err
		}
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		lines, err := utils.ReadLines(file)
		if err != nil {
			return err
		}
		args = append(args, lines...)
	}
	var hosts []string
	for _, arg := range args {
		if arg == "" {
			continue
		}
		if ipRangeRegex.MatchString(arg) {
			ipRange, err := utils.GetIPRange(arg)
			if err != nil {
				return err
			}
			hosts = append(hosts, ipRange...)
		} else {
			hosts = append(hosts, arg)
		}
	}
	var out io.Writer
	if opts.Output == "-" {
		out = os.Stdout
	} else {
		filename, err := utils.ExpandPath(string(opts.Output))
		if err != nil {
			return err
		}
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		out = file
	}
	hostsLen := len(hosts)
	concurrency := utils.Min(opts.Concurrency, hostsLen)
	jobs := make(chan string, concurrency)
	go func() {
		for _, host := range hosts {
			jobs <- host
		}
		close(jobs)
	}()
	results := make(chan string)
	var wg sync.WaitGroup
	wg.Add(hostsLen)
	for i := 0; i < concurrency; i++ {
		go Worker(jobs, results, &wg, opts)
	}
	// из-за того что wg.Wait() не завернул в горутину работало неправильно
	go func() {
		wg.Wait()
		close(results)
	}()
	w := bufio.NewWriter(out)
	for result := range results {
		w.WriteString(fmt.Sprintf("%s\n", result))
		w.Flush()
	}
	return nil
}

// Worker fn
func Worker(
	jobs <-chan string,
	results chan<- string,
	wg *sync.WaitGroup,
	opts *Options,
) {
	for host := range jobs {
		for _, port := range opts.PortList {
			portStr := strconv.Itoa(port)
			hostname := net.JoinHostPort(host, portStr)
			opts.RateLimit.Limiter.Take()
			conn, err := net.DialTimeout("tcp", hostname, opts.Timeout)
			if err == nil {
				conn.Close()
				results <- hostname
			}
		}
		wg.Done()
	}
}
