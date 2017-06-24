
```
JSON API
√ POST /auth/login
√ POST /auth/signup
√ POST /auth/logout
x GET  /api/search
x GET  /api/books/autocomplete
√ POST /api/book
√ GET  /api/books
√ GET  /api/books/mybooks
√ GET  /api/trades/incoming
√ GET  /api/trades/outgoing
x GET  /api/account
√ GET  /api/user/username
√ POST /api/user
√ POST /api/trade
√ PUT  /api/trade/:id
√ DEL  /api/trade/:id
√ GET  /api/book

Navigable react-router routes
/                 root
/login
/signup
/books/all
/books/mybooks
/book/:id         Modal-like UI? Need to save URL state/history leading up for this.
/trades           REDIRECTS to /trades/incoming
/trades/incoming
/trades/outgoing
/account
/new

reach:
/activity         show log of activity
GET /api/activity
```

