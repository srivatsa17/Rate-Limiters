### Sliding Window Log

##### How it works:

1.  Keep a log of request timestamps.
2.  When a new request comes in, remove all entries older than the window size.
3.  Count the remaining entries.
4.  If the count is less than the limit, allow the request and add its timestamp to the log.
5.  If the count exceeds the limit, request is denied.

##### Pros:

1.  Very accurate, no rough edges between windows.
2.  Works well for low-volume APIs.

##### Cons:

1.  Can be memory-intensive for high-volume APIs.
2.  Requires storing and searching through timestamps.
