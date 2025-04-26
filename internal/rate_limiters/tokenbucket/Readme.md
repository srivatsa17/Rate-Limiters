### Token Bucket

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
