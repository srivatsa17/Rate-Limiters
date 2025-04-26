## Rate Limiters

### 1. Token Bucket

##### How It Works:

1.  Imagine a bucket that holds tokens.
2.  The bucket has a maximum capacity of tokens.
3.  Tokens are added to the bucket at a fixed rate (e.g., 10 tokens per second).
4.  When a request arrives, it must obtain a token from the bucket to proceed.
5.  If there are enough tokens, the request is allowed and tokens are removed.
6.  If there aren't enough tokens, the request is dropped.

##### Pros:
1.  Relatively straightforward to implement and understand.
2.  Allows bursts of requests up to the bucket's capacity, accommodating short-term spikes.

##### Cons:
1.  The memory usage scales with the number of users if implemented per-user.
2.  It doesn’t guarantee a perfectly smooth rate of requests.

### 2. Leaky Bucket

##### How it works:

1.  Imagine a bucket with a small hole in the bottom.
2.  Requests enter the bucket from the top.
3.  The bucket processes ("leaks") requests at a constant rate through the hole.
4.  If the bucket is full, new requests are discarded.

##### Pros:

1.  Processes requests at a steady rate, preventing sudden bursts from overwhelming the system.
2.  Provides a consistent and predictable rate of processing requests.

##### Cons:

1.  Does not handle sudden bursts of requests well; excess requests are immediately dropped.
2.  Slightly more complex to implement compared to Token Bucket.

### 3. Fixed Window Counter

##### How it works:

1.  Time is divided into fixed windows (e.g., 1-minute intervals).
2.  Each window has a counter that starts at zero.
3.  New requests increment the counter for the current window.
4.  If the counter exceeds the limit, requests are denied until the next window.

##### Pros:
1.  Easy to implement and understand.
2.  Provides clear and easy-to-understand rate limits for each time window.

##### Cons:
1.  Does not handle bursts of requests at the boundary of windows well.
2.  Can allow twice the rate of requests at the edges of windows.

### 4. Sliding Window Log

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

### 5. Sliding Window Counter

This algorithm combines the Fixed Window Counter and Sliding Window Log approaches for a more accurate and efficient solution.

Instead of keeping track of every single request’s timestamp as the sliding log does, it focus on the number of requests from the last window.

So, if you are in 75% of the current window, 25% of the weight would come from the previous window, and the rest from the current one:

`weight = (100 - 75)% * lastWindowRequests + currentWindowRequests`

Now, when a new request comes, you add one to that weight (weight + 1). If this new total crosses our set limit, we have to reject the request.

##### How it works:

1.  Keep track of request count for the current and previous window.
2.  Calculate the weighted sum of requests based on the overlap with the sliding window.
3.  If the weighted sum is less than the limit, allow the request.

##### Pros:

1.  More accurate than Fixed Window Counter.
2.  More memory-efficient than Sliding Window Log.
3.  Smooths out edges between windows.

##### Cons:

1.  Slightly more complex to implement.

### How to run ?

On the terminal, run the `make run` binary and enter the possible options of the algorithms