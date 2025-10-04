package rwlock

type RWLock struct {
	readerCh    chan struct{}
	writerCh    chan struct{}
	readerCount int
}

func NewRWLock() *RWLock {
	rw := &RWLock{}
	rw.readerCh = make(chan struct{})
	rw.writerCh = make(chan struct{})
	return rw
}

func (rw *RWLock) RLock() {
	rw.readerCh <- struct{}{}
	if rw.readerCount == 0 {
		rw.writerCh <- struct{}{}
	}
	rw.readerCount++
	<-rw.readerCh
}

func (rw *RWLock) RUnlock() {
	rw.readerCh <- struct{}{}
	if rw.readerCount == 1 {
		<-rw.writerCh
	}
	rw.readerCount--
	<-rw.readerCh
}

func (rw *RWLock) Lock() {
	rw.writerCh <- struct{}{}
}

func (rw *RWLock) Unlock() {
	<-rw.writerCh
}
