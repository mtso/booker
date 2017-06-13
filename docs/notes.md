# Notes

## Questions
- How to store session cookie with Go?
- How to parse request body in Content-Type application/json in Go?
- Is it possible to write middleware that attaches arbitrary values to a specific http.Request?
- What is context.Context and how is it used?
- How to implement flash messages in Golang/React?

## API info

Google Books API

https://developers.google.com/books/

### Search by Title
https://www.googleapis.com/books/v1/volumes?q=search+terms


## Ways to render react markup

- single node instance for every request, JSON i/o, very slow
- protocol buffer stream?
- worker process with message queue
- embedded javascript interpreter like goja or otto
- embedded v8 engine with cgo (see `github.com/ry/v8worker`)


## Timeline

- [ ] Store user data and encrypt password (user accounts and auth).
- [ ] Render markup with worker process, and pass state with AMQP.
- [ ] Persist data with Postgres database connection.
- [ ] Authenticate users.
- [ ] Save session using cookies.
- [ ] Add books by connecting to Google Books API.
- [ ] Enable requesting and accepting of trades.
- [ ] Let users edit account info (name, location, city)
