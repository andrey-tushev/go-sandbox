package main

import "fmt"

// Пути находит, но я не уверен что пути оптимальные

func main() {
	s := Solution{}
	//d := s.Solve("hot", "dog", []string{"hot", "dog", "dot"}) // 3
	d := s.Solve("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}) // 5
	fmt.Println("Depth", d)
}

type Solution struct {
	words map[string]bool
	queue []string
}

func (s *Solution) Solve(startWord, finishWord string, list []string) int {
	// делаем из списка мапу
	s.words = make(map[string]bool, len(list))
	for _, w := range list {
		s.words[w] = true
	}

	// добавляем стартовое слово
	s.addWords([]string{startWord})
	delete(s.words, startWord)

	depth := 0
	for {
		currWord := s.getWord()
		if currWord == "" { // очередь опустела, больше вариантов нет
			return -1
		}
		if currWord == finishWord { // нашли
			return depth
		}

		// берем следующие шаги
		nextWords := s.nextFor(currWord)
		s.addWords(nextWords)
		fmt.Println("-", currWord, nextWords)

		depth++
	}
}

// вычисляет следующие возможные шаги
func (s *Solution) nextFor(curWord string) []string {
	nexts := make([]string, 0)
	for candidateWord := range s.words {
		for p := 0; p < len(candidateWord); p++ {
			buf := []byte(curWord)
			for c := 0; c < 26; c++ {
				buf[p] = byte('a' + c)

				if candidateWord == string(buf) && candidateWord != curWord {
					delete(s.words, candidateWord)
					nexts = append(nexts, candidateWord)
				}
			}
		}
	}
	return nexts
}

func (s *Solution) addWords(words []string) {
	s.queue = append(s.queue, words...)
}

func (s *Solution) getWord() string {
	if len(s.queue) == 0 {
		return ""
	}
	word := s.queue[0]
	s.queue = s.queue[1:len(s.queue)]
	return word
}
