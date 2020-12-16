package deps

import (
	"fmt"
	"io"
	"os"

	"github.com/wakatime/wakatime-cli/pkg/heartbeat"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	jww "github.com/spf13/jwalterweatherman"
)

// State is a token parsing state.
type State int

const (
	// StateUnknown represents a unknown token parsing state.
	StateUnknown State = iota
	// StateImport means we are in import section during token parsing.
	StateImport
)

// DependencyParser is a dependency parser for a programming language.
type DependencyParser interface {
	Parse(reader io.ReadCloser, lexer chroma.Lexer) ([]string, error)
}

// WithDetection initializes and returns a heartbeat handle option, which
// can be used in a heartbeat processing pipeline to detect dependencies
// inside the entity file of heartbeats of type FileType. Will prioritize
// local file if available.
func WithDetection() heartbeat.HandleOption {
	return func(next heartbeat.Handle) heartbeat.Handle {
		return func(hh []heartbeat.Heartbeat) ([]heartbeat.Result, error) {
			for n, h := range hh {
				if h.EntityType != heartbeat.FileType {
					continue
				}

				filepath := h.Entity

				if h.LocalFile != "" {
					filepath = h.LocalFile
				}

				dependencies, err := Detect(filepath, h.Language)
				if err != nil {
					jww.WARN.Printf("error detecting dependencies of heartbeat: %s", err)
					continue
				}

				hh[n].Dependencies = dependencies
			}

			return next(hh)
		}
	}
}

// Detect parses the dependencies from a heartbeat file of a specific language.
func Detect(filepath string, language heartbeat.Language) ([]string, error) {
	lexer := lexers.Get(language.StringChroma())
	if lexer == nil {
		return nil, fmt.Errorf("unable to detect lexer for language %q", language)
	}

	var parser DependencyParser

	switch language {
	case heartbeat.LanguageGo:
		parser = &ParserGo{}
	default:
		jww.DEBUG.Printf("parsing dependencies not supported for language %q", language)
		return nil, nil
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %s", filepath, err)
	}

	return parser.Parse(f, lexer)
}