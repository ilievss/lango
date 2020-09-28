package translate

import (
	"strings"
)

const writeQueueSize = 100

type HistoryEntry struct {
	Input       string
	Translation string
}

// Thread-safe implementation for storing and retrieving translation history
type TranslationHistory struct {
	entries          []HistoryEntry
	writeQueue       chan HistoryEntry
	readRequestQueue chan interface{}
	readQueue        chan []HistoryEntry
}

func NewTranslationHistory() *TranslationHistory {
	th := TranslationHistory{
		entries: make([]HistoryEntry, 0, 10),
		// put a generically sized buffer, so we seldom block callers
		writeQueue:       make(chan HistoryEntry, writeQueueSize),
		readRequestQueue: make(chan interface{}),
		readQueue:        make(chan []HistoryEntry),
	}
	go func() {
		for {
			select {
			case he := <-th.writeQueue:
				th.entries = append(th.entries, he)
			case <-th.readRequestQueue:
				entriesCopy := make([]HistoryEntry, len(th.entries))
				copy(entriesCopy, th.entries)
				th.readQueue <- entriesCopy
			}
		}
	}()
	return &th
}

func (th *TranslationHistory) Add(entry HistoryEntry) {
	th.writeQueue <- entry
}

func (th *TranslationHistory) Entries() []HistoryEntry {
	th.readRequestQueue <- 1
	return <-th.readQueue
}

func (th *TranslationHistory) String() string {
	builder := strings.Builder{}
	builder.Grow(512 * len(th.entries))
	builder.WriteString("[")
	for _, en := range th.entries {
		builder.WriteString("{Input: \"")
		builder.WriteString(en.Input)
		builder.WriteString("\", Translation: \"")
		builder.WriteString(en.Translation)
		builder.WriteString("\"}")
	}
	builder.WriteString("]")
	return builder.String()
}
