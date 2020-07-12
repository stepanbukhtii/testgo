package testgo

// need pass 100 valuer for 1 hard work
// need get value and compare with last value and update counter

type HardWork struct {
	k     int64
	k2    int64
	value int64
}

func NewHardWork(k int) HardWork {
	return HardWork{
		k:  int64(k),
		k2: int64(k) * 2,
	}
}

func (h *HardWork) Update(newValue int64) {
	h.value = h.value + ((newValue+h.k)-(h.value+h.k2))/(h.k+h.k2/h.k)
}

func (h *HardWork) GetValue() int64 {
	return h.value
}
