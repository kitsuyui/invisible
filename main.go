package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime/debug"
	"time"

	"github.com/google/subcommands"
	"github.com/kitsuyui/invisible/embedding"
	"github.com/kitsuyui/invisible/simplenoise"
)

type addNoise struct {
	frequency float64
	maxSize   int
	seed      int64
}

func (*addNoise) Name() string     { return "add-noise" }
func (*addNoise) Synopsis() string { return "read from stdin and write to stdout with noise." }
func (*addNoise) Usage() string {
	return `add-noise:
	Read from stdin and write to stdout with noise.
	Output varies on each run (non-deterministic) unless --seed is specified.
	Use --seed <N> for reproducible output. encode/decode are always deterministic.
`
}

func (p *addNoise) SetFlags(f *flag.FlagSet) {
	f.Float64Var(&p.frequency, "frequency", 0.5, "frequency for noise")
	f.Float64Var(&p.frequency, "f", 0.5, "frequency for noise")
	f.IntVar(&p.maxSize, "noise-size", 1, "max noise in once")
	f.IntVar(&p.maxSize, "s", 1, "max noise in once")
	f.Int64Var(&p.seed, "seed", 0, "random seed (0 = use time-based random seed, non-deterministic)")
}

func (p *addNoise) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.frequency < 0 {
		fmt.Fprintln(os.Stderr, "error: frequency must be non-negative")
		return subcommands.ExitUsageError
	}
	if p.maxSize < 0 {
		fmt.Fprintln(os.Stderr, "error: noise-size must be non-negative")
		return subcommands.ExitUsageError
	}
	seed := p.seed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rng := rand.New(rand.NewSource(seed))
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	if err := simplenoise.AddRandomNoise(rng, p.frequency, p.maxSize, reader, writer); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

type encode struct {
	message string
}

func (*encode) Name() string { return "encode" }
func (*encode) Synopsis() string {
	return "read from stdin and write to stdout with noise made from encoded message"
}
func (*encode) Usage() string {
	return `encode:
	Read from stdin and write to stdout with noise made from encoded message.
`
}

func (p *encode) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.message, "m", "Hello, World!", "message")
	f.StringVar(&p.message, "message", "Hello, World!", "message")
}

func (p *encode) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.message == "" {
		fmt.Fprintln(os.Stderr, "error: message must not be empty")
		return subcommands.ExitUsageError
	}
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	if err := embedding.Embed(p.message, reader, writer, true); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

type decode struct{}

func (*decode) Name() string { return "decode" }
func (*decode) Synopsis() string {
	return "read from stdin and write to stdout decoded message"
}
func (*decode) Usage() string {
	return `decode:
	Read from stdin and write to stdout decoded message.
`
}

func (p *decode) SetFlags(f *flag.FlagSet) {}

func (p *decode) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	decoded, err := embedding.Extract(reader, bufio.NewWriter(ioutil.Discard))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}
	if _, err := writer.WriteString(decoded); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}
	if err := writer.Flush(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

type versionCmd struct{}

func (*versionCmd) Name() string     { return "version" }
func (*versionCmd) Synopsis() string { return "print version information." }
func (*versionCmd) Usage() string {
	return `version:
	Print version information.
`
}
func (*versionCmd) SetFlags(f *flag.FlagSet) {}

func (*versionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	info, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Println(info.Main.Version)
	} else {
		fmt.Println("(unknown)")
	}
	return subcommands.ExitSuccess
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&addNoise{}, "")
	subcommands.Register(&encode{}, "")
	subcommands.Register(&decode{}, "")
	subcommands.Register(&versionCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
