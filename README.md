Logey
=====

Money logging for the web.

Building
--------

Execute:

```
go get github.com/ishiikurisu/logey
go get github.com/ishiikurisu/logeyweb
go build github.com/ishiikurisu/logeyweb
```

API
---

This server also provides an API to use Logey in your application by storing the log in a string and providing some tools to interact with it. The actions to post are the following:

### `POST /api/user/new`

Creates a new user.

Parameters:

- `username`
- `password`
- `log`: The log to use when creating the user. Optional.

Returns:

- `authentication`: it's a secret.
- `log`: the user's log. If not provided on creation, will create a dummy log for him.

### `POST /api/user/login`

Logs the user in.

Parameters:

- `username`
- `password`

Returns:

- `authentication`
- `log`
- `error`: a message detailing what went wrong. `null` otherwise.

### `POST /api/log`

Adds a new

Parameters:

- `authentication`
- `log`

Returns:

- `error`

### `POST /api/log/entries`

Gets the user's log.

Parameters:

- `authentication`

Returns:

- `log`
