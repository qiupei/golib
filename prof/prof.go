package prof

import (
    "os"
    "runtime/pprof"
)

var (
    cpuFile, heapFile *os.File
)

func Start(fn string) (err error) {
    cpuFile , err = os.Create(fn)
    if err != nil {
        return
    }
    pprof.StartCPUProfile(cpuFile)
    return
}

func Stop() {
    pprof.StopCPUProfile()
    cpuFile.Close()
}

func NewDump(fn string) (err error) {
    heapFile, err = os.Create(fn)
    if err != nil {
        return
    }
    return
}

func Dump() {
    pprof.WriteHeapProfile(heapFile)
}

func CloseDump() {
    heapFile.Close()
}
