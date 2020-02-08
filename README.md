# golibs

[![CircleCI](https://circleci.com/gh/fahlke/golibs/tree/master.svg?style=svg)](https://circleci.com/gh/fahlke/golibs/tree/master)
[![Codecov](https://codecov.io/gh/fahlke/golibs/branch/master/graph/badge.svg)](https://codecov.io/gh/fahlke/golibs)
[![GoDoc](https://godoc.org/github.com/fahlke/golibs?status.svg)](https://godoc.org/github.com/fahlke/golibs)
[![Go Report Card](https://goreportcard.com/badge/github.com/fahlke/golibs)](https://goreportcard.com/report/github.com/fahlke/golibs)

A loose collection of handcrafted [Go][golang] implementations of common data structures and algorithms for my own reference and for demonstration purposes.

[![gopher-logo]][gopherizeme]

## Implementations

The lists below are in no particular order and the various implementations are not guaranteed to be "perfect". Everything is implemented for mere self educational purposes and can very likely be improved on.

### Data structures

| Kind (Wikipedia link)                        | Package                  | Implemented |
|:---------------------------------------------|:-------------------------|:-----------:|
| [Queue (FIFO)][queue-fifo]                   | [queue](queue)           |           - |
| [Stack (LIFO)][stack-lifo]                   | [stack](stack)           |           - |
| [Max heap (Binary heap)][max-heap]           | [heap](heap)             |           - |
| [Min heap (Pairing heap)][min-heap]          | [heap](heap)             |           - |
| [Hash table (hashmap)][hash-table]           | [hashmap](hashmap)       |           - |
| [Binary tree][binary-tree]                   | [binarytree](binarytree) |           - |
| [Linked list][linked-list]                   | [linkedlist](linkedlist) |           - |
| [Undirected graph][undirected-graph]         | [graph](graph)           |           - |

### Sorting algorithms

| Kind (Wikipedia link)                        | Package                  | Implemented |
|:---------------------------------------------|:-------------------------|:-----------:|
| [Quicksort][quick-sort]                      | [sort](sort)             |           - |
| [Bubble sort][bubble-sort]                   | [sort](sort)             |           - |
| [Merge sort][merge-sort]                     | [sort](sort)             |           - |
| [Heapsort][heap-sort]                        | [sort](sort)             |           - |
| [Selection sort][selection-sort]             | [sort](sort)             |           - |
| [Insertion sort][insertion-sort]             | [sort](sort)             |           - |
| [Cocktail shaker sort][cocktail-shaker-sort] | [sort](sort)             |           - |
| [Shellsort][shell-sort]                      | [sort](sort)             |           - |
| [Radix sort][radix-sort]                     | [sort](sort)             |           - |

### Hashing algorithms

| Kind (Wikipedia link)                        | Package                  | Implemented |
|:---------------------------------------------|:-------------------------|:-----------:|
| [SHA-256][sha-256]                           | [hash](hash)             |           - |
| [MD5][md5]                                   | [hash](hash)             |           - |
| [Pearson hashing][pearson-hashing]           | [hash](hash)             |           - |
| [CRC-32][crc-32]                             | [hash](hash)             |           - |

### String metrics

| Kind (Wikipedia link)                        | Package                  | Implemented |
|:---------------------------------------------|:-------------------------|:-----------:|
| [Levenshtein distance][levenshtein]          | [string](string)         |           - |
| [Hamming distance][hamming]                  | [string](string)         |           - |

### Utilities

| Function                                     | Package                    | Implemented |
|:---------------------------------------------|:---------------------------|:-----------:|
| Mix                                          | [util/mix.go](util/mix.go) |           ✓ |
| Max                                          | [util/max.go](util/max.go) |           ✓ |

[golang]: https://golang.org/ref/spec
[gopher-logo]: assets/gopher-mini.png "Gopher logo"
[gopherizeme]: https://gopherize.me/gopher/692fbfd019724e297a7b0761bd3e9697ef8e6bee
[queue-fifo]: https://en.wikipedia.org/wiki/FIFO_(computing_and_electronics)
[stack-lifo]: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
[max-heap]: https://en.wikipedia.org/wiki/Binary_heap
[min-heap]: https://en.wikipedia.org/wiki/Pairing_heap
[quick-sort]: https://en.wikipedia.org/wiki/Quicksort
[bubble-sort]: https://en.wikipedia.org/wiki/Bubble_sort
[merge-sort]: https://en.wikipedia.org/wiki/Merge_sort
[heap-sort]: https://en.wikipedia.org/wiki/Heapsort
[selection-sort]: https://en.wikipedia.org/wiki/Selection_sort
[insertion-sort]: https://en.wikipedia.org/wiki/Insertion_sort
[cocktail-shaker-sort]: https://en.wikipedia.org/wiki/Cocktail_shaker_sort
[shell-sort]: https://en.wikipedia.org/wiki/Shellsort
[radix-sort]: https://en.wikipedia.org/wiki/Radix_sort
[hash-table]: https://en.wikipedia.org/wiki/Hash_table
[linked-list]: https://en.wikipedia.org/wiki/Linked_list
[sha-256]: https://en.wikipedia.org/wiki/SHA-2
[md5]: https://en.wikipedia.org/wiki/MD5
[crc-32]: https://en.wikipedia.org/wiki/Cyclic_redundancy_check#CRC-32_algorithm
[pearson-hashing]: https://en.wikipedia.org/wiki/Pearson_hashing
[undirected-graph]: https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)#Graph
[binary-tree]: https://en.wikipedia.org/wiki/Binary_tree
[levenshtein]: https://en.wikipedia.org/wiki/Levenshtein_distance
[hamming]: https://en.wikipedia.org/wiki/Hamming_distance
