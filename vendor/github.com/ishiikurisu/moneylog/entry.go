package moneylog

import "fmt"
import "strings"

// Defines an entry, the fundamental element in the money log.
type Entry struct {
    // How the money was spent or earned.
    Description string

    // The money amount of this entry.
    Value float64
}

// Creates an entry from a description and a value.
func NewEntry(description string, value float64) Entry {
    return Entry{Description: description, Value: value}
}

// Creates an entry from a string in the format "Description: Value".
func EntryFromString(raw string) Entry {
    var contents []string = strings.Split(raw, ":")
    var description string = contents[0]
    var value float64
    fmt.Sscanf(contents[1], "%f", &value)
    return NewEntry(description, value)
}
