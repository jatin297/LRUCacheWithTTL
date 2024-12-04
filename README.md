LRU Cache - A Simple and Efficient Caching With Time-To-Live expiry feature in Go

An LRU Cache storing frequently used data and evicting older or unused data automatically.

🚀 What It Does

Core Features:
- Stores data and removes the least recently used items when out of give caching capacity.
- Manages memory effectively by evicting old data.
- Handles simultaneous requests without issues like race condition.

Functionality:
- Retrieves and updates data in no time O(1) time-complexity for lookups and insertion in doubly linked list.
- Customize you cache size limit.
- Designed to perform well under many requests.

Easy to Use:
- Clean and organized code that’s easy to understand and modify.
- Comes with Doubly Linked List tests to ensure it works as expected.
