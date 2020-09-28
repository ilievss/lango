package translate

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTranslationHistoryConcurrency(t *testing.T) {
	numEntries := 50
	th := NewTranslationHistory()

	for i := 0; i < numEntries; i++ {
		i := i
		go func() {
			th.Add(HistoryEntry{
				Input:       strconv.Itoa(i),
				Translation: strconv.Itoa(i),
			})
		}()
	}

	// wait until the goroutines finish pushing entries
	time.Sleep(100 * time.Millisecond)

	numbers := make([]bool, numEntries)
	entries := th.Entries()
	for i := 0; i < numEntries; i++ {
		num, err := strconv.Atoi(entries[i].Input)
		if err != nil {
			t.Fatal(fmt.Sprintf("Failed to convert string '%s' to number", entries[i].Input))
		}
		if numbers[num] {
			t.Fatal(fmt.Sprintf("Number %d found more than once: %s", i, th))
		}
		numbers[num] = true
	}
}
