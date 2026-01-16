# Server Monitor - Refactored Architecture

## Overview

This is a professional, maintainable server monitoring application built with Go and termbox for terminal UI. It monitors RPC nodes, databases, web APIs, disk usage, and image servers.

## Architecture

The application follows a clean, modular architecture with clear separation of concerns:

```
srvmonitor/
├── main.go                     # Application entry point
├── config/                     # Configuration management
│   └── config.go              # Loads from environment variables
├── display/                    # Display abstraction
│   ├── interface.go           # Display interface definition
│   └── termbox_display.go     # Termbox implementation
├── monitor/                    # Monitoring modules
│   ├── interface.go           # Monitor interface
│   ├── manager.go             # Monitor coordinator
│   ├── rpc_monitor.go         # RPC node monitoring
│   ├── database_monitor.go    # Layer 1 database monitoring
│   ├── application_monitor.go # Application layer monitoring
│   ├── webapi_monitor.go      # Web API monitoring
│   ├── disk_monitor.go        # Disk usage monitoring
│   └── image_monitor.go       # Image availability monitoring
├── types/                      # Shared types
│   └── types.go               # Common data structures
└── utils/                      # Utility functions
    └── db.go                  # Database utilities
```

## Key Design Principles

### 1. **Configuration from Environment**
All configuration is loaded from environment variables (same as before):
```bash
export RPC0_NAME="..."
export RPC0_URL="..."
export DB_L1_HOST_SRV1="..."
# etc.
```

No config files needed - the config layer just centralizes env var reading and provides validation.

### 2. **Monitor Interface Pattern**
Each monitor implements a simple interface:
```go
type Monitor interface {
    Name() string
    Start(ctx context.Context, display Display, errorChan chan<- string)
    GetDisplayPosition() Position
}
```

This makes it easy to:
- Add new monitors without changing existing code
- Test monitors independently
- Run specific monitors selectively

### 3. **Display Abstraction**
The display layer is abstracted behind an interface:
```go
type Display interface {
    Init() error
    Close() error
    Clear()
    DrawLine(line DisplayLine)
    DrawText(pos Position, text string, fg, bg Color)
    Flush()
    Size() (width, height int)
}
```

Benefits:
- Can swap termbox for another UI (web, GUI, etc.)
- Thread-safe with built-in mutex
- Testable with mock implementations

### 4. **Manager Pattern**
The `MonitorManager` coordinates all monitors:
- Registers monitors dynamically
- Starts them in separate goroutines
- Collects errors from all monitors
- Displays errors centrally

### 5. **Context-Based Cancellation**
Uses Go's context for clean shutdown:
- Ctrl+C or 'q' key triggers cancellation
- All monitors respect context cancellation
- Graceful cleanup on exit

## Adding a New Monitor

To add a new monitoring component:

1. **Create a new file** in `monitor/` directory
2. **Implement the Monitor interface:**
```go
type MyNewMonitor struct {
    config   MyConfig
    position types.Position
}

func NewMyNewMonitor(cfg MyConfig) *MyNewMonitor {
    return &MyNewMonitor{
        config:   cfg,
        position: types.Position{X: 0, Y: 50}, // Choose Y position
    }
}

func (m *MyNewMonitor) Name() string {
    return "My New Monitor"
}

func (m *MyNewMonitor) GetDisplayPosition() types.Position {
    return m.position
}

func (m *MyNewMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Do monitoring work
            m.check(disp, errorChan)
            time.Sleep(60 * time.Second)
        }
    }
}

func (m *MyNewMonitor) check(disp display.Display, errorChan chan<- string) {
    // Perform checks
    // Report errors: errorChan <- "error message"
    // Display results: disp.DrawText(...)
    disp.Flush()
}
```

3. **Register it in main.go:**
```go
myMon := monitor.NewMyNewMonitor(myConfig)
mgr.Register(myMon)
```

That's it! No need to modify existing monitors.

## Configuration Validation

The config layer validates required settings at startup:
```go
cfg, err := config.LoadFromEnv()
if err != nil {
    // Clear error message about what's missing
    log.Fatal(err)
}
```

If a required environment variable is missing, you get a clear error message immediately instead of a panic deep in the code.

## Thread Safety

All display operations are protected by a mutex in the `TermboxDisplay` implementation. Multiple monitors can safely draw to the screen concurrently.

## Error Handling

Errors flow through a central channel:
- Each monitor sends errors to `errorChan`
- The manager collects and displays them
- Errors are also logged to `/tmp/srvmonitor.log`

## Testing

The architecture is designed for testability:

```go
// Mock display for testing
type MockDisplay struct {
    drawn []string
}

func (m *MockDisplay) DrawText(pos types.Position, text string, fg, bg types.Color) {
    m.drawn = append(m.drawn, text)
}

// Test a monitor
func TestMyMonitor(t *testing.T) {
    mock := &MockDisplay{}
    mon := NewMyMonitor(testConfig)
    mon.check(mock, make(chan string, 10))
    // Assert on mock.drawn
}
```

## Migration from Old Version

The old implementation is preserved in `old_version/` directory. The new version:
- Uses the same environment variables
- Provides the same functionality
- Has the same display layout
- Is more maintainable and extensible

## Running

```bash
# Set environment variables (same as before)
export RPC0_NAME="..."
export RPC0_URL="..."
# ... all other env vars ...

# Run
./srvmonitor

# Exit
Press 'q' or Ctrl+C
```

## Logging

Logs are written to `/tmp/srvmonitor.log`:
- Application startup
- Monitor registration
- Errors from monitors
- Shutdown events

Old log is moved to `/tmp/srvmonitor-old.log` on restart.

## Benefits of Refactoring

### Before (Old Architecture)
- ❌ Global variables everywhere
- ❌ Hard-coded array sizes
- ❌ Scattered environment variable reads
- ❌ Tight coupling between monitors
- ❌ Difficult to test
- ❌ Hard to add new monitors
- ❌ No clear module boundaries

### After (New Architecture)
- ✅ Clean module boundaries
- ✅ Interface-based design
- ✅ Centralized configuration
- ✅ Easy to add new monitors
- ✅ Testable components
- ✅ Thread-safe by design
- ✅ Context-based cancellation
- ✅ Clear error handling
- ✅ Flexible for future changes

## Future Enhancements

The architecture makes these enhancements trivial:

1. **Web UI**: Implement a new `WebDisplay` and swap it in
2. **Metrics Export**: Add a Prometheus exporter monitor
3. **Alerting**: Add monitors that send alerts to Slack/Email
4. **Configuration File**: Add YAML/TOML loader alongside env vars
5. **Remote Monitoring**: Add monitors for remote services
6. **Historical Data**: Add time-series storage monitors
7. **Custom Dashboards**: Mix and match monitors easily

## Code Quality

- ✅ No linter errors
- ✅ Clear naming conventions
- ✅ Proper error handling
- ✅ Thread-safe operations
- ✅ Context-aware goroutines
- ✅ Clean separation of concerns
- ✅ Single Responsibility Principle
- ✅ Open/Closed Principle (easy to extend, no need to modify)


