Logey
=====

Money logging for the web.

About
-----

Logey is open source software distributed under [Unlicense](http://unlicense.org).

Its color palette was generated with [coolors.co](https://coolors.co/5bc0eb-fde74c-9bc53d-c3423f-211a1e).

Building
--------

Just run

```
go get github.com/ishiikurisu/logey
go get github.com/ishiikurisu/logeyweb
go build github.com/ishiikurisu/logeyweb
```

API
---

This server also provides an API to use Logey in your application by storing the log in a string and providing some tools to interact with it. The actions to post are the following:

### `/api/new`

- Creates a new log string to be used.
- Returns an empty log in the standard string format.

### `/api/add`

- Adds an entry to the log.
- Parameter `log`: the log that will have the entry added.
- Parameter `description`: the entry's description.
- Parameter `value`: the entry's value.
- Returns the log in standard format with the added entry.

### `/api/get/balance`

- Calculates the balance of the log.
- Parameter `log`: the log in standard format.
- Returns the balance in string format.

### `/api/get/descriptions`

- Gets the description of each entry.
- Parameter `log`: the log in standard format.
- Returns a list of descriptions separated by commas.

### `/api/get/values`

- Gets the value of each entry.
- Parameter `log`: the log in standard format.
- Returns a list of values separated by commas.
