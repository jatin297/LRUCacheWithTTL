LRU Cache - A Simple and Efficient Caching With Time-To-Live Expiry in Go

An LRU Cache that stores frequently used data and automatically evicts older or unused data, with an added Time-To-Live (TTL) expiration feature for precise control over cached items.

**🚀 What It Does**

**Core Features**
- Smart Caching: Stores data and removes the least recently used items when the cache reaches its capacity.
- Time-To-Live Expiry: Automatically removes keys that have expired based on their TTL, ensuring fresh data.
- Memory Efficiency: Evicts old data to optimize memory usage.
- Thread-Safe Design: Handles simultaneous requests without race conditions.


**Functionality**
- Fast Operations: Retrieves and updates data with O(1) time complexity using a combination of a hash map and a doubly linked list.
- Active Expiration: Uses a Go ticker to periodically remove expired keys, running cleanup tasks concurrently for better performance.
- Dynamic Expiration Capacity: Adjusts the frequency of expiry checks based on workload, ensuring scalability.
- Customizable Cache: Easily configure the cache size and expiration behavior to suit application needs.
- High Performance Under Load: Efficiently handles large numbers of read and write operations.
