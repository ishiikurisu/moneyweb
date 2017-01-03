/*
This package implements the fundamental
*/

package moneylog

import "fmt"
import "strings"

// Defines the log data structure
type Log struct {
    // Is an array of entry objects
    Entries []Entry
}

// Creates an empty log without entries
func EmptyLog() Log {
    return Log{Entries: make([]Entry, 0)}
}

// Starts the log with a first entry
func NewLog(first Entry) Log {
    log := Log{Entries: make([]Entry, 1)}
    log.Entries[0] = first
    return log
}

// Starts the log with a first entry, describing its description and its value
func StartLog(description string, value float64) Log {
    return NewLog(NewEntry(description, value))
}

// Adds an entry to a log
func (log *Log) Insert(entry Entry) {
    log.Entries = append(log.Entries, entry)
}

// Adds an entry to a log
func (log *Log) Add(description string, value float64) {
    log.Insert(NewEntry(description, value))
}

// TODO: Change string standard for money log. I can't store a cookie like that.
// Turns a log into a YAML string
func (log *Log) ToString() string {
    outlet := "---\n"

    for _, entry := range log.Entries {
        outlet += fmt.Sprintf("%s: %.2f\n", entry.Description, entry.Value)
    }

    return fmt.Sprintf("%s...\n", outlet)
}

// Loads a log from a YAML string
func LogFromString(raw string) Log {
    inlet := strings.Split(raw, "\n")
    log := EmptyLog()

    for _, line := range inlet {
        if line == "---" || line == "..."  {

        } else if len(line) > 0 {
            log.Insert(EntryFromString(line))
        }
    }

    return log
}

// Calculates the current log balance
func (log *Log) CalculateBalance() float64 {
    var outlet float64 = 0

    for _, entry := range log.Entries {
        outlet += entry.Value
    }

    return outlet
}
