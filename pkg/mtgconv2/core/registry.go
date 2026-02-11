package core

import "sync"

// maintains a registry of all the input and output handlers
// handlers should add themselves to the registry with their 'init' methods

var (
	sourceMu sync.RWMutex
	outputMu sync.RWMutex
	sources  = map[InputSource]SourceHandler{}
	outputs  = map[OutputFormat]OutputHandler{}
)

func RegisterSource(h SourceHandler) {
	sourceMu.Lock()
	defer sourceMu.Unlock()
	sources[h.Source()] = h
}

func HandlerForSource(src InputSource) (SourceHandler, error) {
	sourceMu.RLock()
	defer sourceMu.RUnlock()
	h, ok := sources[src]
	if !ok {
		return nil, &UnrecognizedDomain{Message: string(src)}
	}
	return h, nil
}

func RegisterOutput(h OutputHandler) {
	outputMu.Lock()
	defer outputMu.Unlock()
	outputs[h.Format()] = h
}

func HandlerForOutput(fmt OutputFormat) (OutputHandler, error) {
	outputMu.RLock()
	defer outputMu.RUnlock()
	h, ok := outputs[fmt]
	if !ok {
		return nil, &UnknownOutputFormat{Format: fmt}
	}
	return h, nil
}
