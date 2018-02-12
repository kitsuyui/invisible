package main

import (
	"bufio"
	"context"
	"flag"
	"io/ioutil"
	"os"

	"github.com/google/subcommands"
	"github.com/kitsuyui/invisible/embedding"
	"github.com/kitsuyui/invisible/simplenoise"
)

type addNoise struct {
	frequency float64
	maxSize   int
}

func (*addNoise) Name() string     { return "add-noise" }
func (*addNoise) Synopsis() string { return "read from stdin and write to stdout with noise." }
func (*addNoise) Usage() string {
	return `add-noise:
	Read from stdin and write to stdout with noise.
`
}

func (p *addNoise) SetFlags(f *flag.FlagSet) {
	f.Float64Var(&p.frequency, "frequency", 0.5, "frequency for noise")
	f.Float64Var(&p.frequency, "f", 0.5, "frequency for noise")
	f.IntVar(&p.maxSize, "noise-size", 1, "max noise in once")
	f.IntVar(&p.maxSize, "s", 1, "max noise in once")
}

func (p *addNoise) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	simplenoise.AddRandomNoise(p.frequency, p.maxSize, reader, writer)
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
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	embedding.Embed(p.message, reader, writer, true)
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
	decoded := embedding.Extract(reader, bufio.NewWriter(ioutil.Discard))
	writer.WriteString(decoded)
	writer.Flush()
	return subcommands.ExitSuccess
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&addNoise{}, "")
	subcommands.Register(&encode{}, "")
	subcommands.Register(&decode{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
