
request -> approve -> received

## User

field | data type  | options
------|------------|----------
name  | string     | not null
password_hash | string | not null
city  | string     |
state | varchar(2) |

## Book

field     | data type  | options
----------|------------|----------------------------------
id        | bigserial  | primary key, not null, unique
title     | string     | not null
isbn      | string
image_url | string
userId    | integer    | not null
status    | integer    | not null, (isOwned\|isRequested)

## Trade

field  | data type | options
-------|-----------|---------------------------------------------
userId | string    | not null
bookId | string    | not null
status | integer   | not null, (isPending\|isApproved\|isFailed)

# Tests

UPDATE trades
   SET status = CASE
                WHEN user_id = $2 THEN 'StatusAccepted'
                ELSE 'StatusCanceled' 
                END
   WHERE book_id = $1

