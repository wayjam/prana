// This file was generated by counterfeiter
package fake

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/phogolabs/parcello"
)

type Buffer struct {
	WriteStub        func(p []byte) (n int, err error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		p []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	ReadStub        func(p []byte) (n int, err error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		p []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	ReaddirStub        func(count int) ([]os.FileInfo, error)
	readdirMutex       sync.RWMutex
	readdirArgsForCall []struct {
		count int
	}
	readdirReturns struct {
		result1 []os.FileInfo
		result2 error
	}
	StatStub        func() (os.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct{}
	statReturns     struct {
		result1 os.FileInfo
		result2 error
	}
	SeekStub        func(offset int64, whence int) (int64, error)
	seekMutex       sync.RWMutex
	seekArgsForCall []struct {
		offset int64
		whence int
	}
	seekReturns struct {
		result1 int64
		result2 error
	}
	OpenStub        func(name string) (parcello.File, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		name string
	}
	openReturns struct {
		result1 parcello.File
		result2 error
	}
	WalkStub        func(dir string, fn filepath.WalkFunc) error
	walkMutex       sync.RWMutex
	walkArgsForCall []struct {
		dir string
		fn  filepath.WalkFunc
	}
	walkReturns struct {
		result1 error
	}
	OpenFileStub        func(name string, flag int, perm os.FileMode) (parcello.File, error)
	openFileMutex       sync.RWMutex
	openFileArgsForCall []struct {
		name string
		flag int
		perm os.FileMode
	}
	openFileReturns struct {
		result1 parcello.File
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Buffer) Write(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.recordInvocation("Write", []interface{}{pCopy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(p)
	}
	return fake.writeReturns.result1, fake.writeReturns.result2
}

func (fake *Buffer) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *Buffer) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].p
}

func (fake *Buffer) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Buffer) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	return fake.closeReturns.result1
}

func (fake *Buffer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *Buffer) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Buffer) Read(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.recordInvocation("Read", []interface{}{pCopy})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(p)
	}
	return fake.readReturns.result1, fake.readReturns.result2
}

func (fake *Buffer) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *Buffer) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].p
}

func (fake *Buffer) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Buffer) Readdir(count int) ([]os.FileInfo, error) {
	fake.readdirMutex.Lock()
	fake.readdirArgsForCall = append(fake.readdirArgsForCall, struct {
		count int
	}{count})
	fake.recordInvocation("Readdir", []interface{}{count})
	fake.readdirMutex.Unlock()
	if fake.ReaddirStub != nil {
		return fake.ReaddirStub(count)
	}
	return fake.readdirReturns.result1, fake.readdirReturns.result2
}

func (fake *Buffer) ReaddirCallCount() int {
	fake.readdirMutex.RLock()
	defer fake.readdirMutex.RUnlock()
	return len(fake.readdirArgsForCall)
}

func (fake *Buffer) ReaddirArgsForCall(i int) int {
	fake.readdirMutex.RLock()
	defer fake.readdirMutex.RUnlock()
	return fake.readdirArgsForCall[i].count
}

func (fake *Buffer) ReaddirReturns(result1 []os.FileInfo, result2 error) {
	fake.ReaddirStub = nil
	fake.readdirReturns = struct {
		result1 []os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *Buffer) Stat() (os.FileInfo, error) {
	fake.statMutex.Lock()
	fake.statArgsForCall = append(fake.statArgsForCall, struct{}{})
	fake.recordInvocation("Stat", []interface{}{})
	fake.statMutex.Unlock()
	if fake.StatStub != nil {
		return fake.StatStub()
	}
	return fake.statReturns.result1, fake.statReturns.result2
}

func (fake *Buffer) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *Buffer) StatReturns(result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *Buffer) Seek(offset int64, whence int) (int64, error) {
	fake.seekMutex.Lock()
	fake.seekArgsForCall = append(fake.seekArgsForCall, struct {
		offset int64
		whence int
	}{offset, whence})
	fake.recordInvocation("Seek", []interface{}{offset, whence})
	fake.seekMutex.Unlock()
	if fake.SeekStub != nil {
		return fake.SeekStub(offset, whence)
	}
	return fake.seekReturns.result1, fake.seekReturns.result2
}

func (fake *Buffer) SeekCallCount() int {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return len(fake.seekArgsForCall)
}

func (fake *Buffer) SeekArgsForCall(i int) (int64, int) {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return fake.seekArgsForCall[i].offset, fake.seekArgsForCall[i].whence
}

func (fake *Buffer) SeekReturns(result1 int64, result2 error) {
	fake.SeekStub = nil
	fake.seekReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *Buffer) Open(name string) (parcello.File, error) {
	fake.openMutex.Lock()
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Open", []interface{}{name})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(name)
	}
	return fake.openReturns.result1, fake.openReturns.result2
}

func (fake *Buffer) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *Buffer) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].name
}

func (fake *Buffer) OpenReturns(result1 parcello.File, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 parcello.File
		result2 error
	}{result1, result2}
}

func (fake *Buffer) Walk(dir string, fn filepath.WalkFunc) error {
	fake.walkMutex.Lock()
	fake.walkArgsForCall = append(fake.walkArgsForCall, struct {
		dir string
		fn  filepath.WalkFunc
	}{dir, fn})
	fake.recordInvocation("Walk", []interface{}{dir, fn})
	fake.walkMutex.Unlock()
	if fake.WalkStub != nil {
		return fake.WalkStub(dir, fn)
	}
	return fake.walkReturns.result1
}

func (fake *Buffer) WalkCallCount() int {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return len(fake.walkArgsForCall)
}

func (fake *Buffer) WalkArgsForCall(i int) (string, filepath.WalkFunc) {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return fake.walkArgsForCall[i].dir, fake.walkArgsForCall[i].fn
}

func (fake *Buffer) WalkReturns(result1 error) {
	fake.WalkStub = nil
	fake.walkReturns = struct {
		result1 error
	}{result1}
}

func (fake *Buffer) OpenFile(name string, flag int, perm os.FileMode) (parcello.File, error) {
	fake.openFileMutex.Lock()
	fake.openFileArgsForCall = append(fake.openFileArgsForCall, struct {
		name string
		flag int
		perm os.FileMode
	}{name, flag, perm})
	fake.recordInvocation("OpenFile", []interface{}{name, flag, perm})
	fake.openFileMutex.Unlock()
	if fake.OpenFileStub != nil {
		return fake.OpenFileStub(name, flag, perm)
	}
	return fake.openFileReturns.result1, fake.openFileReturns.result2
}

func (fake *Buffer) OpenFileCallCount() int {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return len(fake.openFileArgsForCall)
}

func (fake *Buffer) OpenFileArgsForCall(i int) (string, int, os.FileMode) {
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return fake.openFileArgsForCall[i].name, fake.openFileArgsForCall[i].flag, fake.openFileArgsForCall[i].perm
}

func (fake *Buffer) OpenFileReturns(result1 parcello.File, result2 error) {
	fake.OpenFileStub = nil
	fake.openFileReturns = struct {
		result1 parcello.File
		result2 error
	}{result1, result2}
}

func (fake *Buffer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.readdirMutex.RLock()
	defer fake.readdirMutex.RUnlock()
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	fake.openFileMutex.RLock()
	defer fake.openFileMutex.RUnlock()
	return fake.invocations
}

func (fake *Buffer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}