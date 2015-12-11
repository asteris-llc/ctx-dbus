# ctx-dbus
[![Build Status](https://travis-ci.org/asteris-llc/ctx-dbus.svg)](https://travis-ci.org/asteris-llc/ctx-dbus)

Context-enabled interface to the github.com/coreos/go-systemd/dbus package

```
PACKAGE DOCUMENTATION

package dbus
    import "github.com/asteris-llc/ctx-dbus"

    Package dbus provides a Context-enabled interface to the
    github.com/coreoes/go-systemd/dbus interface

CONSTANTS

const DefaultTimeout = time.Minute

FUNCTIONS

func Reload(unit string) error
    Reload a Systemd unit

func ReloadCtx(ctx context.Context, unit string) error

func Restart(unit string) error
    Restart a Systemd unit

func RestartCtx(ctx context.Context, unit string) error

func Start(unit string) error
    Start a Systemd unit

func StartCtx(ctx context.Context, unit string) error

func Stop(unit string) error
    Stop a Systemd unit

func StopCtx(ctx context.Context, unit string) error

TYPES

type Dbus struct {
    // contains filtered or unexported fields
}

func New(timeout time.Duration, l *log.Logger) *Dbus
    Creates a new Dbus

func (d *Dbus) Reload(unit string) error
func (d *Dbus) ReloadCtx(ctx context.Context, unit string) error
func (d *Dbus) Restart(unit string) error
func (d *Dbus) RestartCtx(ctx context.Context, unit string) error
func (d *Dbus) Start(unit string) error
func (d *Dbus) StartCtx(ctx context.Context, unit string) error
func (d *Dbus) Stop(unit string) error
func (d *Dbus) StopCtx(ctx context.Context, unit string) error

```
