package main

import "fmt"

func main() {
	task := Task{}
	task.set = []string{"A", "B", "C", "D"}
	task.onlyOnce = true
	task.search(0)
}

type Task struct {
	// true  - Только перестановки - N!
	// false - Без уникальности    - N^N
	onlyOnce bool

	set    []string
	subset []string
}

func (t *Task) search(level int) {
	if level == len(t.set) {
		t.process()
	} else {
		// для каждого возможного элемента
		for i := 0; i < len(t.set); i++ {
			if t.onlyOnce && t.isUsed(i) {
				continue
			}

			t.add(i)            // добавим элемент
			t.search(level + 1) // сходим глубже
			t.remove()          // удалим элемент
		}
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

func (t *Task) isUsed(k int) bool {
	value := t.set[k]
	for _, subsetValue := range t.subset {
		if subsetValue == value {
			return true
		}
	}
	return false
}
