package internal

import (
	"bytes"
	"io"
	"syscall"
	"unsafe"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	procOpenFileMapping = kernel32.NewProc("OpenFileMappingW")
)

const MmfSize uintptr = 32 * 1024

type Shm interface {
	Reader() io.Reader
	Close() error
}

type shm struct {
	handle syscall.Handle
	addr   uintptr
	data   []byte
}

func OpenShm(mmfName string) (Shm, error) {
	handle, err := openFileMappingW(syscall.FILE_MAP_READ, false, mmfName)
	if err != nil {
		return nil, err
	}

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_READ, 0, 0, MmfSize)
	if err != nil {
		_ = syscall.CloseHandle(handle)
		return nil, err
	}

	//goland:noinspection ALL
	data := unsafe.Slice((*byte)(unsafe.Pointer(addr)), MmfSize)

	return &shm{
		handle: handle,
		addr:   addr,
		data:   data,
	}, nil
}

func openFileMappingW(access int32, inheritHandle bool, name string) (handle syscall.Handle, err error) {
	var arg1 uint32
	if inheritHandle {
		arg1 = 1
	}
	arg2, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	r, _, e := procOpenFileMapping.Call(
		uintptr(access),
		uintptr(arg1),
		uintptr(unsafe.Pointer(arg2)),
	)
	handle = syscall.Handle(r)
	if r == 0 {
		err = e
	}
	return
}

func (s *shm) Reader() io.Reader {
	return bytes.NewReader(s.data)
}

func (s *shm) Close() error {
	if s.addr != 0 {
		err := syscall.UnmapViewOfFile(s.addr)
		if err != nil {
			return err
		}
		s.addr = 0
	}
	if s.handle != syscall.InvalidHandle {
		err := syscall.CloseHandle(s.handle)
		if err != nil {
			return err
		}
		s.handle = syscall.InvalidHandle
	}
	return nil
}
