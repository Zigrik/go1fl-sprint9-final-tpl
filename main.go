package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var (
	randSource = rand.NewSource(time.Now().UnixNano())
	randRange  = rand.New(randSource)
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size < 1 {
		err := fmt.Errorf("the number of items to generate cannot be less than 1")
		log.Println("Error:", err)
		return []int{}
	}

	maxNum := math.MaxInt - 1
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = randRange.Intn(maxNum) + 1
	}

	return arr
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		err := fmt.Errorf("the number of elements to search for the maximum must be greater than 0")
		log.Println("Error:", err)
		return 0
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		err := fmt.Errorf("the number of elements to search for the maximum must be greater than 0")
		log.Println("Error:", err)
		return 0
	}
	results := make([]int, CHUNKS)
	n := min(CHUNKS, len(data))

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func() {
			if i < n { //cutting off goroutines that have nothing to do
				chunk := (data[i*len(data)/n : ((i + 1) * len(data) / n)])
				results[i] = slices.Max(chunk)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return slices.Max(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	nums := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(nums)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(nums)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
