// Package dbus provides a Context-enabled interface to
// the github.com/coreoes/go-systemd/dbus interface
package dbus

import (
	"log"
	"os"
	"time"

	"github.com/coreos/go-systemd/dbus"
	"golang.org/x/net/context"
)

const DefaultTimeout = time.Minute

var std = New(DefaultTimeout, log.New(os.Stderr, "", log.LstdFlags))

type Dbus struct {
	timeout time.Duration
	logger *log.Logger
}

// Creates a new Dbus
func New(timeout time.Duration, l *log.Logger) *Dbus {
	return &Dbus{
		timeout: timeout,
		logger: l,
	}
}

func (d *Dbus) Start(unit string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return d.StartCtx(ctx, unit)
}

func (d *Dbus) StartCtx(ctx context.Context, unit string) error {
	c, err := dbus.New()
	if err != nil {
		return err
	}
	defer c.Close()

	if _, err := c.StartUnit(unit, "replace", nil); err != nil {
		return err
	}

	return d.waitForStatus(ctx, c, unit, "active")
}

func (d *Dbus) Stop(unit string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return d.StopCtx(ctx, unit)
}

func (d *Dbus) StopCtx(ctx context.Context, unit string) error {
	c, err := dbus.New()
	if err != nil {
		return err
	}
	defer c.Close()

	if _, err := c.StopUnit(unit, "replace", nil); err != nil {
		return err
	}

	return d.waitForStatus(ctx, c, unit, "inactive")
}

func (d *Dbus) Reload(unit string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return d.ReloadCtx(ctx, unit)
}

func (d *Dbus) ReloadCtx(ctx context.Context, unit string) error {
	c, err := dbus.New()
	if err != nil {
		return err
	}
	defer c.Close()

	if _, err := c.ReloadUnit(unit, "replace", nil); err != nil {
		return err
	}

	return d.waitForStatus(ctx, c, unit, "active")
}

func (d *Dbus) Restart(unit string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return d.RestartCtx(ctx, unit)
}

func (d *Dbus) RestartCtx(ctx context.Context, unit string) error {
	c, err := dbus.New()
	if err != nil {
		return err
	}
	defer c.Close()

	if _, err := c.RestartUnit(unit, "replace", nil); err != nil {
		return err
	}

	return d.waitForStatus(ctx, c, unit, "active")
}

func (d *Dbus) waitForStatus(ctx context.Context, c *dbus.Conn, unit string, status string) error {
	cs := c.NewSubscriptionSet()
	cs.Add(unit)

	nctx, cancel := context.WithTimeout(ctx, d.timeout)
	defer cancel()

	statusCh, errorCh := cs.Subscribe()
	for {
		select {
		case smap := <-statusCh:
			if us, ok := smap[unit]; ok {
				if us.ActiveState == status {
					return nil
				}
			}
		case err := <-errorCh:
			if d.logger != nil {
				d.logger.Print(err)
			}
		case <-nctx.Done():
			return nctx.Err()
		}
	}
}
