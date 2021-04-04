package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Location struct {
	x int64
	y int64
	z int64
}

type Simple struct {
	deadTime float64
	location Location
	transportId int64
	launcherId  int64
	id int64
}

type Launcher struct {
	count     int64
	frequency int64
}

type Transport struct {
	size      int64
	tubeSize  int64
	processTime int64
}

type Receiver struct {
	tubeCount     int64
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// 按照 Simple.deadTime 从大到小排序
type SimpleSlice []*Simple

func (a SimpleSlice) Len() int {    // 重写 Len() 方法
	return len(a)
}
func (a SimpleSlice) Swap(i, j int){     // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a SimpleSlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[j].deadTime < a[i].deadTime
}

func makeSimpleList(size int64) (simpleList []*Simple) {
	var i int64 = 0
	for i <= size {
		simple := &Simple{
			id: i,
			deadTime: float64(RandInt64(2*60,100*60)),
		}
		i++
		simpleList = append(simpleList, simple)
	}

	return
}

func timeCount(size int64) (time float64,ok bool) {
	ok = true
	time = 0
	var i int64 = 0
	inputList := makeSimpleList(size)
	boxList := make([]*Simple,0)
	for i < size {
		time = time + 1.2
		boxList = append(boxList, inputList[i])

		if len(boxList) > 2500 {
			fmt.Println("failed")
			ok = false
			break
		}
		i++
		if i%4 == 0 {
			sort.Sort(sort.Reverse(SimpleSlice(boxList)))
			fmt.Printf("在时间为%f发射id为%d的样本，样本剩余时间%f\n",time,boxList[0].id,boxList[0].deadTime)
			boxList = boxList[1:]
			for _, i := range(boxList) {
				i.deadTime = i.deadTime - 3.6
				if i.deadTime <= 0 {
					fmt.Printf("failed,在时间为%f发射id为%d的样本，样本剩余时间%f\n", time, boxList[0].id, boxList[0].deadTime)
					break
				}
			}
		}
	}
	if len(boxList) > 0 {
		fmt.Printf("在时间为%f发射id为%d的样本，样本剩余时间%f\n",time,boxList[0].id,boxList[0].deadTime)
		boxList = boxList[1:]
		time = time + 3.6
		for _, i := range(boxList) {
			i.deadTime = i.deadTime - 3.6
			if i.deadTime <= 0 {
				fmt.Printf("failed,在时间为%f发射id为%d的样本，样本剩余时间%f\n", time, boxList[0].id, boxList[0].deadTime)
				break
			}
		}
	}
	fmt.Printf("total_time:%f\n", time)
	return
}

func main() {
	timeCount(3000)
}
