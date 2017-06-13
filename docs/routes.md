
```
JSON API
POST /api/auth/login
POST /api/auth/signup
POST /api/auth/logout
GET  /api/books/autocomplete
POST /api/book
POST /api/trade
GET  /api/books
GET  /api/books/mybooks
GET  /api/trades/incoming
GET  /api/trades/outgoing
GET  /api/account
POST /api/account


Navigable react-router routes
/                 root
/login  
/books/all
/books/mybooks
/book/:id         Modal-like UI? Need to save URL state/history leading up for this.
/trades           REDIRECTS to /trades/incoming
/trades/incoming
/trades/outgoing
/account


reach:
/activity         show log of activity
GET /api/activity
```

