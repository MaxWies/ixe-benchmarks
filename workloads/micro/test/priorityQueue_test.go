package test

import (
	"container/heap"
	"faas-micro/utils"
	"math"
	"testing"
)

func TestPriorityQueueMax(t *testing.T) {
	items := []int64{
		7,
		math.MaxInt64,
		10000,
		0,
	}
	pq := utils.PriorityQueueMax{
		Limit: 10,
		Items: make([]int64, len(items)),
	}
	for i, p := range items {
		pq.Items[i] = p
	}
	heap.Init(&pq)
	heap.Push(&pq, int64(16))
	sortedItems := []int64{
		0,
		7,
		16,
		10000,
		math.MaxInt64,
	}
	for i := range sortedItems {
		item := heap.Pop(&pq).(int64)
		if item != sortedItems[i] {
			t.Error("Wrong priority")
		}
	}
}

func TestPriorityQueueMin(t *testing.T) {
	items := []int64{
		7,
		math.MaxInt64,
		10000,
		0,
	}
	pq := utils.PriorityQueueMin{
		Limit: 10,
		Items: make([]int64, len(items)),
	}
	for i, p := range items {
		pq.Items[i] = p
	}
	heap.Init(&pq)
	heap.Push(&pq, int64(16))
	sortedItems := []int64{
		math.MaxInt64,
		10000,
		16,
		7,
		0,
	}
	for i := range sortedItems {
		item := heap.Pop(&pq).(int64)
		if item != sortedItems[i] {
			t.Error("Wrong priority")
		}
	}
}

func TestPriorityQueueMax_Shrinking(t *testing.T) {
	items := []int64{
		7,
		10000,
		0,
		5,
		99,
		1,
		1,
	}
	pq := utils.PriorityQueueMax{
		Limit: 1,
		Items: make([]int64, len(items)),
	}
	for i, p := range items {
		pq.Items[i] = p
	}
	heap.Init(&pq)
	//heap.Fix(&pq)
	pq.Shrink()
	item := pq.Pop().(int64)
	if item != 10000 {
		t.Error("Wrong priority")
	}
}

func TestPriorityQueueMax_Add(t *testing.T) {
	items := []int64{
		1000,
		10000,
	}
	pq := utils.PriorityQueueMax{
		Limit: 3,
		Items: make([]int64, len(items)),
	}
	for i, p := range items {
		pq.Items[i] = p
	}
	heap.Init(&pq)
	pq.Add(5)
	if item, _ := pq.Peek(); item != 5 {
		t.Error("Wrong priority")
	}
	pq.Add(5000)
	if item, _ := pq.Peek(); item != 1000 {
		t.Error("Wrong priority")
	}
	pq.Limit = 2
	pq.Shrink()
	v := heap.Pop(&pq).(int64)
	if v != 5000 {
		t.Error("Wrong priority")
	}
	v = heap.Pop(&pq).(int64)
	if v != 10000 {
		t.Error("Wrong priority")
	}
}
