# Notes

## Questions
- 

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
