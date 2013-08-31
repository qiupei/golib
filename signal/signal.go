package signal

import (
    "os"
    S "os/signal"
)

// return true if need to break
type Callback func() bool

type Handler struct {
    schan chan os.Signal
    cb map[os.Signal]Callback
}

func NewHandler() (sh *Handler) {
    sh = &Handler{
        schan: make(chan os.Signal),
        cb: make(map[os.Signal]Callback, 5),
    }
    S.Notify(sh.schan, os.Interrupt, os.Kill)
    return
}

func (sh *Handler)Bind(s os.Signal, cb Callback) {
    sh.cb[s] = cb
}

func (sh *Handler)Loop() os.Signal {
    for s := range sh.schan {
        if f, ok := sh.cb[s]; ok {
            if f() {
	            return s
			}
        }
    }
    return nil
}

func (sh *Handler) Close() {
    close(sh.schan)
}


var (
    DefaultHandler = NewHandler()
)

func Bind(s os.Signal, cb Callback) {
    DefaultHandler.Bind(s, cb)
}

func Loop() os.Signal {
    return DefaultHandler.Loop()
}

func Send(pid int, signal os.Signal) error {
    proc, err := os.FindProcess(pid)
    if err != nil {
        return err
    }
    if err := proc.Signal(signal); err != nil {
        return err
    }
    return nil
}

func Close() {
	DefaultHandler.Close()
}
