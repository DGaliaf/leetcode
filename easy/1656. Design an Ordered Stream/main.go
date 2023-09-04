package easy

type OrderedStream struct {
	ptr    int
	orders []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr:    0,
		orders: make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.orders[idKey-1] = value

	chank := make([]string, 0)

	for i := this.ptr; i < len(this.orders); i++ {
		if this.orders[i] == "" {
			this.ptr = i
			break
		}

		chank = append(chank, this.orders[i])
	}

	return chank
}
