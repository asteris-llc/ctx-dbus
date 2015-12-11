package dbus

import (
	"golang.org/x/net/context"
)

// Start a Systemd unit
func Start(unit string) error {
	return std.Start(unit)
}

func StartCtx(ctx context.Context, unit string) error {
	return std.StartCtx(ctx, unit)
}

// Stop a Systemd unit
func Stop(unit string) error {
	return std.Stop(unit)
}

func StopCtx(ctx context.Context, unit string) error {
	return std.StopCtx(ctx, unit)
}

// Restart a Systemd unit
func Restart(unit string) error {
	return std.Restart(unit)
}

func RestartCtx(ctx context.Context, unit string) error {
	return std.RestartCtx(ctx, unit)
}

// Reload a Systemd unit
func Reload(unit string) error {
	return std.Reload(unit)
}

func ReloadCtx(ctx context.Context, unit string) error {
	return std.ReloadCtx(ctx, unit)
}
