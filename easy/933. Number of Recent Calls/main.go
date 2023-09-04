package easy

type RecentCounter struct {
	time []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		time: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.time = append(this.time, t)

	for _, pt := range this.time {
		if pt < t-3000 {
			this.time = this.time[1:]
		}
	}

	return len(this.time)
}
