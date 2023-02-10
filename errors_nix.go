// Unless explicitly stated otherwise all files in this repository are licensed
// under MIT License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

//go:build !windows
// +build !windows

package python3

/*
#include "Python.h"
*/
import "C"

//PySignal_SetWakeupFd : https://docs.python.org/3/c-api/exceptions.html#c.PySignal_SetWakeupFd
//func PySignal_SetWakeupFd(fd uintptr) uintptr {
//	return uintptr(C.PySignal_SetWakeupFd(C.int(fd)))
//}
