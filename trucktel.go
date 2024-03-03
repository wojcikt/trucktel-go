package trucktel_go

import (
	"errors"
	"github.com/wojcikt/trucktel-go/internal"
)

const (
	DefaultMmfName = "Local\\SCSTelemetry"
)

type Telemetry interface {
	Read(values *Values) error
	Close() error
}

type telemetry struct {
	shm internal.Shm
}

func Open() (Telemetry, error) {
	return OpenFile(DefaultMmfName)
}

func OpenFile(name string) (Telemetry, error) {
	shm, err := internal.OpenShm(name)
	if err != nil {
		return nil, err
	}
	return &telemetry{shm}, nil
}

func (t *telemetry) Read(v *Values) error {
	if t.shm == nil {
		return errors.New("closed")
	}
	r := t.shm.Reader()
	return v.read(r)
}

func (t *telemetry) Close() error {
	err := t.shm.Close()
	if err != nil {
		return err
	}
	t.shm = nil
	return nil
}
