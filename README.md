# cp-go
Basic structure for competitive programming in go

The idea of the library is to provide the minimal functionality needed in competitive programming without having any external dependencies.

**Note: The structures and methods written are specifically written to make it easy to compete in programming contests. Hence, error handling is not done! Panics are used in place of error at most of the places.**

No part of this code should EVER be used as a production code. 

## How-to

Just run the following command to generate main.go which would contain basic STL like structures that can be used in competitive programming.

```sh
$ cat $(ls *.go | grep -v _test.go$) | grep -v "package cpgo" > main.go

```

# Basic things I'm planning to have

- Objects
  [ ] String
  [ ] Integer
  [ ] Double
  [ ] Bool
  [ ] Pair
- Vector
  [ ] Get
  [ ] Set
  [ ] Len
  [ ] PushBack
  [ ] PushFront
  [ ] PopBack
  [ ] PopFront
  [ ] Sort
- PriorityQueue
- Queue (Linked list)
- Deque
