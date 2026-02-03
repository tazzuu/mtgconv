package core

import (
	"log/slog"
	"os"
)

// initialize internal logger
func ConfigureLogging(verbose bool) {

	// default log level is INFO
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	// create a TextHandler and set it globally
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:     level,
		AddSource: false, // adds the line number from source code
		// redact the user agent token string from output
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
        if a.Key == "UserAgent" {
            return slog.String(a.Key, "[REDACTED]")
        }
        if a.Key == "config" {
            if cfg, ok := a.Value.Any().(Config); ok {
                cfg.UserAgent = "[REDACTED]"
                return slog.Any(a.Key, cfg)
            }
        }
        return a
    },
	})
	slog.SetDefault(slog.New(handler))
	slog.Debug("logger initialized", "verbose", verbose)
}
