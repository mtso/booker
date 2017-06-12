
request -> approve -> received

## User

field | data type  | options
------|------------|----------
name  | string     | not null
city  | string     |
state | varchar(2) |

## Book

field   | data type  | options
--------|------------|----------------------------------
userId  | string     | not null
status  | integer    | not null, (isOwned\|isRequested)

## Trade

field  | data type | options
-------|-----------|---------------------------------------------
userId | string    | not null
bookId | string    | not null
status | integer   | not null, (isPending\|isApproved\|isFailed)


