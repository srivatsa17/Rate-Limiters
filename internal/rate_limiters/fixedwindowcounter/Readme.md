### Fixed Window Counter

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