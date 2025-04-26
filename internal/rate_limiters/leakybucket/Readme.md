### Leaky Bucket

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
