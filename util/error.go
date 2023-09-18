package util

import (
	"errors"
)

var (
	ErrBPFProgramNotLoaded      = errors.New("the bpf program was not loaded")
	ErrBPFProgramNotPinned      = errors.New("the bpf program was not pinned")
	ErrBPFMapFailedToCheck      = errors.New("failed to check if the bpf map exists")
	ErrBPFMapNotExist           = errors.New("the bpf map does not exist")
	ErrBPFMapNotLoaded          = errors.New("the bpf map was not loaded")
	ErrBPFMapNotPinned          = errors.New("the bpf map was not pinned")
	ErrBPFMapNotUpdated         = errors.New("the bpf map was not updated")
	ErrBPFMapNotRemoved         = errors.New("the bpf map was not removed")
	ErrBPFCustomCallMapNotExist = errors.New("the bpf map for custom call does not exist")
)
