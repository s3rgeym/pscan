package internal

import (
	"strconv"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
	"go.uber.org/ratelimit"
)

// Options type
type Options struct {
	Concurrency int            `short:"c" long:"concurrency" default:"10" description:"Number of concurrent jobs"`
	Input       flags.Filename `short:"i" long:"input" description:"Input filename"`
	Output      flags.Filename `short:"o" long:"output" default:"-" description:"Output filename"`
	PortList    portList       `short:"p" long:"ports" default:"21-23,25,53,80,110-111,135,139,143,443,445,993,995,1723,3306,3389,5900,8080" description:"Ports for scanning"`
	RateLimit   rateLimit      `short:"l" long:"rate-limit" default:"100" description:"Requests per second"`
	Timeout     time.Duration  `short:"t" long:"timeout" default:"5s" description:"Connect timeout"`
}

// invalid receiver type *rateLimit
type rateLimit struct {
	Value int
	ratelimit.Limiter
}

func (m *rateLimit) UnmarshalFlag(s string) error {
	value, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	m.Value = value
	m.Limiter = ratelimit.New(value)
	return nil
}

type portList []int

func (l *portList) UnmarshalFlag(s string) error {
	items := strings.Split(s, ",")
	for _, item := range items {
		slice := strings.SplitN(item, "-", 2)
		if len(slice) == 2 {
			start, err := strconv.Atoi(slice[0])
			if err != nil {
				return err
			}
			end, err := strconv.Atoi(slice[0])
			if err != nil {
				return err
			}
			for i := start; i <= end; i++ {
				*l = append(*l, i)
			}
		} else {
			port, err := strconv.Atoi(slice[0])
			if err != nil {
				return err
			}
			*l = append(*l, port)
		}
	}
	return nil
}
