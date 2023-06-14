package main

import "fmt"

func main() {
	task := Task{}
	task.set = []string{"A", "B", "C", "D"}
	task.search(0)

	task2 := Task2{}
	task2.set = []string{"A", "B", "C", "D"}
	task2.subset = make([]string, len(task2.set))
	task2.search(0)
}

// Вариант со стеком

type Task struct {
	set    []string
	subset []string
}

func (t *Task) search(k int) {
	// if k > len(t.set)-1
	if k == len(t.set) { // самый нижний уровень, уже перебрали все возможное
		t.process() // их будет 2^N
	} else {
		t.add(k) // Добавим элемент и перепробуем все с ним
		t.search(k + 1)

		t.remove() // Уберем, и перепробуем все без него
		t.search(k + 1)
	}
}

func (t *Task) process() {
	fmt.Println(t.subset)
}

func (t *Task) add(k int) {
	v := t.set[k]
	t.subset = append(t.subset, v)
}

func (t *Task) remove() {
	n := len(t.subset) - 1
	t.subset = t.subset[0:n]
}

// Вариант с пустыми значениями в слайсе

type Task2 struct {
	set    []string
	subset []string
}

func (t *Task2) search(k int) {
	if k == len(t.set) {
		t.process()
	} else {
		// Кстати, здесь можно и местами поменять добавление и удаление (будет инверсно)
		t.add(k)
		t.search(k + 1)

		t.remove(k)
		t.search(k + 1)
	}
}

func (t *Task2) process() {
	fmt.Println(t.subset)
}

func (t *Task2) add(k int) {
	v := t.set[k]
	t.subset[k] = v
}

func (t *Task2) remove(k int) {
	t.subset[k] = "."
}
