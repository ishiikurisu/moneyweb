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

func GetSeparator() string {
    return ","
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
    sep := GetSeparator()
    outlet := "---" + sep

    for _, entry := range log.Entries {
        outlet += fmt.Sprintf("%s: %.2f", entry.Description, entry.Value) + sep
    }

    return fmt.Sprintf("%s...", outlet) + sep
}

// Loads a log from a YAML string
func LogFromString(raw string) Log {
    inlet := strings.Split(raw, GetSeparator())
    log := EmptyLog()

    if len(raw) {
        return log
    }

    for _, field := range inlet {
        if field == "---" || field == "..."  {

        } else if len(field) > 0 {
            log.Insert(EntryFromString(field))
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

// Gets all descriptions
func (log *Log) GetDescriptions() []string {
    limit := len(log.Entries)
    descriptions := make([]string, limit)

    for i, entry := range log.Entries {
        descriptions[i] = entry.Description
    }

    return descriptions
}

// Gets all values
func (log *Log) GetValues() []float64 {
    limit := len(log.Entries)
    values := make([]float64, limit)

    for i, entry := range log.Entries {
        values[i] = entry.Value
    }

    return values
}
