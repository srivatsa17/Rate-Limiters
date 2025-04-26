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
