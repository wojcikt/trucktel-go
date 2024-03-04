package shm

import (
	"bytes"
	"syscall"
	"unsafe"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	procOpenFileMapping = kernel32.NewProc("OpenFileMappingW")
)

const (
	FileSize        uintptr = 32 * 1024
	DefaultFileName         = "Local\\SCSTelemetry"
)

type SharedMemory interface {
	Reader() *bytes.Reader
	Read(values *Values) error
	Close() error
}

type sharedMemory struct {
	handle syscall.Handle
	addr   uintptr
	data   []byte
}

func Open(mmfName string) (SharedMemory, error) {
	handle, err := openFileMappingW(syscall.FILE_MAP_READ, false, mmfName)
	if err != nil {
		return nil, err
	}

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_READ, 0, 0, FileSize)
	if err != nil {
		_ = syscall.CloseHandle(handle)
		return nil, err
	}

	//goland:noinspection ALL
	data := unsafe.Slice((*byte)(unsafe.Pointer(addr)), FileSize)

	return &sharedMemory{
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

func (mem *sharedMemory) Reader() *bytes.Reader {
	return bytes.NewReader(mem.data)
}

func (mem *sharedMemory) Read(values *Values) error {
	reader := bytes.NewReader(mem.data)
	return values.Read(reader)
}

func (mem *sharedMemory) Close() error {
	if mem.addr != 0 {
		err := syscall.UnmapViewOfFile(mem.addr)
		if err != nil {
			return err
		}
		mem.addr = 0
	}
	if mem.handle != syscall.InvalidHandle {
		err := syscall.CloseHandle(mem.handle)
		if err != nil {
			return err
		}
		mem.handle = syscall.InvalidHandle
	}
	return nil
}
