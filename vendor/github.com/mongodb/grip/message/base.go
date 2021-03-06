package message

import (
	"fmt"
	"os"
	"time"

	"github.com/mongodb/grip/level"
)

// Base provides a simple embedable implementation of some common
// aspects of a message.Composer. Additionally the Collect() method
// collects some simple metadata, that may be useful for some more
// structured logging applications.
type Base struct {
	Level    level.Priority `bson:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	Hostname string         `bson:"hostname,omitempty" json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Time     time.Time      `bson:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
	Process  string         `bson:"process,omitempty" json:"process,omitempty" yaml:"process,omitempty"`
	Pid      int            `bson:"pid,omitempty" json:"pid,omitempty" yaml:"pid,omitempty"`
}

// Collect records the time, process name, and hostname. Useful in the
// context of a Raw() method.
func (b *Base) Collect() error {
	if !b.Time.IsZero() {
		return nil
	}

	var err error
	b.Hostname, err = os.Hostname()
	if err != nil {
		return err
	}

	b.Time = time.Now()
	b.Process = os.Args[0]
	b.Pid = os.Getpid()

	return nil
}

// Priority returns the configured priority of the message.
func (b *Base) Priority() level.Priority {
	return b.Level
}

// SetPriority allows you to configure the priority of the
// message. Returns an error if the priority is not valid.
func (b *Base) SetPriority(l level.Priority) error {
	if !level.IsValidPriority(l) {
		return fmt.Errorf("%s (%d) is not a valid priority", l, l)
	}

	b.Level = l

	return nil
}
