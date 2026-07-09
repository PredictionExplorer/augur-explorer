// slog handler for the ETL binaries' legacy two-file layout: everything goes
// to the info log, records at Error and above are duplicated into the error
// log. This keeps the operational file conventions (ae_logs/<name>_info.log,
// <name>_error.log) while the engine logs structured records; §8.3 of the
// modernization roadmap replaces file logging wholesale.

package indexer

import (
	"context"
	"io"
	"log/slog"
)

type dualLogHandler struct {
	info slog.Handler
	err  slog.Handler
}

// NewDualLogHandler returns a text handler writing every record to infoW and
// duplicating records at Error and above to errW.
func NewDualLogHandler(infoW, errW io.Writer) slog.Handler {
	return &dualLogHandler{
		info: slog.NewTextHandler(infoW, nil),
		err:  slog.NewTextHandler(errW, nil),
	}
}

func (h *dualLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.info.Enabled(ctx, level)
}

func (h *dualLogHandler) Handle(ctx context.Context, r slog.Record) error {
	err := h.info.Handle(ctx, r.Clone())
	if r.Level >= slog.LevelError {
		if dupErr := h.err.Handle(ctx, r); err == nil {
			err = dupErr
		}
	}
	return err
}

func (h *dualLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &dualLogHandler{info: h.info.WithAttrs(attrs), err: h.err.WithAttrs(attrs)}
}

func (h *dualLogHandler) WithGroup(name string) slog.Handler {
	return &dualLogHandler{info: h.info.WithGroup(name), err: h.err.WithGroup(name)}
}
