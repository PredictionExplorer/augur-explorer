package srvmonitor

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestDiskMonitorRendersDfOutput(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewDiskMonitor([]DiskConfig{
		{Title: "srv1", User: "ops", IP: "10.0.0.1", DeviceList: "/dev/sda1"},
	}, logger, testIntervals())

	var gotArgs []string
	m.run = func(_ context.Context, name string, args ...string) ([]byte, error) {
		gotArgs = append([]string{name}, args...)
		return []byte("Mounted on Use%\n/          42%\n"), nil
	}

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	if m.statuses[0].ErrStr != "" {
		t.Fatalf("ErrStr = %q", m.statuses[0].ErrStr)
	}

	want := []string{
		"/usr/bin/ssh",
		"-o", "StrictHostKeyChecking=accept-new",
		"-l", "ops",
		"10.0.0.1",
		"df --output=target,pcent", "/dev/sda1",
	}
	if len(gotArgs) != len(want) {
		t.Fatalf("ssh args = %v, want %v", gotArgs, want)
	}
	for i := range want {
		if gotArgs[i] != want[i] {
			t.Fatalf("ssh arg %d = %q, want %q", i, gotArgs[i], want[i])
		}
	}

	// The first df line is a header and the last is empty; only the middle
	// line renders.
	if row := disp.Row(20); !strings.Contains(row, "42%") {
		t.Fatalf("row 20 = %q, want df content line", row)
	}
	if title := disp.Row(19); !strings.Contains(title, "srv1") {
		t.Fatalf("title row = %q", title)
	}
	if header := disp.Row(17); !strings.Contains(header, "Disk Usage (df)") {
		t.Fatalf("header = %q", header)
	}
}

func TestDiskMonitorColumnLayout(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	disks := []DiskConfig{{Title: "a"}, {Title: "b"}, {Title: "c"}, {Title: "d"}}
	m := NewDiskMonitor(disks, logger, testIntervals())

	wantX := []int{1, 25, 50, 1} // fourth disk falls back to column 1
	for i, want := range wantX {
		if m.statuses[i].X != want {
			t.Fatalf("disk %d X = %d, want %d", i, m.statuses[i].X, want)
		}
	}
}

func TestDiskMonitorCommandFailure(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewDiskMonitor([]DiskConfig{{Title: "srv1", User: "ops", IP: "10.0.0.1"}}, logger, testIntervals())
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		return nil, errors.New("exit status 255")
	}

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	if m.statuses[0].ErrStr != "exit status 255" {
		t.Fatalf("ErrStr = %q", m.statuses[0].ErrStr)
	}
	if msgs := drain(errCh); len(msgs) != 1 {
		t.Fatalf("errors = %v", msgs)
	}
	if row := disp.Row(20); !strings.Contains(row, "ERROR: exit status 255") {
		t.Fatalf("row = %q", row)
	}
}

func TestDiskMonitorPendingShowsChecking(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewDiskMonitor([]DiskConfig{{Title: "srv1"}}, logger, testIntervals())

	disp := newFakeDisplay()
	m.display(disp) // no check ran yet: no output, no error

	if row := disp.Row(20); !strings.Contains(row, "Checking...") {
		t.Fatalf("row = %q, want Checking placeholder", row)
	}
}

func TestDiskMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewDiskMonitor([]DiskConfig{{Title: "srv1"}}, logger, testIntervals())
	cycles := make(chan struct{}, 100)
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		select {
		case cycles <- struct{}{}:
		default:
		}
		return []byte("h\nc\n"), nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		m.Start(ctx, newFakeDisplay(), make(chan string, 100))
		close(done)
	}()

	<-cycles
	<-cycles
	cancel()
	waitFor(t, "loop exit", func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	})
}
