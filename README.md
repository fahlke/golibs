# golibs

[![CircleCI](https://img.shields.io/circleci/build/github/fahlke/golibs?label=circleci)](https://circleci.com/gh/fahlke/golibs/tree/master)
[![Codecov](https://img.shields.io/codecov/c/github/fahlke/golibs?label=codecov)](https://codecov.io/gh/fahlke/golibs)
[![GoDoc](https://godoc.org/github.com/fahlke/golibs?status.svg)](https://godoc.org/github.com/fahlke/golibs)
[![Go Report Card](https://goreportcard.com/badge/github.com/fahlke/golibs)](https://goreportcard.com/report/github.com/fahlke/golibs)

A loose collection of handcrafted [Go][golang] implementations of common data structures and algorithms for my own reference and for demonstration purposes.

But most importantly this repository is created and maintained to have some ***fun*** developing stuff and to ***stay hungry, stay curious***!

[![gopher-logo]][gopherizeme]

## Implementations

The lists below are in no particular order and the various implementations are not guaranteed to be "perfect". Everything is implemented for mere self educational purposes and can very likely be improved on.

### Data structures

| Kind (Wikipedia link)                        | Package                              | Implemented |
|:---------------------------------------------|:-------------------------------------|:-----------:|
| [List][list]                                 | [list](list)                         |           - |
| [Sorted list][sorted-list]                   | [sortedlist](sortedlist)             |           - |
| [Linked list][linked-list]                   | [linkedlist](linkedlist)             |           ✓ |
| [Queue (FIFO)][queue-fifo]                   | [queue](queue)                       |           ✓ |
| [Stack (LIFO)][stack-lifo]                   | [stack](stack)                       |           ✓ |
| [Max heap (Binary heap)][max-heap]           | [heap/max](heap/max)                 |           - |
| [Min heap (Pairing heap)][min-heap]          | [heap/min](heap/min)                 |           - |
| [Hash table (hashmap)][hash-table]           | [hashmap](hashmap)                   |           ✓ |
| [Tree][tree]                                 | [tree](tree)                         |           - |
| [Trie][trie]                                 | [trie](trie)                         |           - |
| [Binary search tree][binary-tree]            | [binarytree](binarytree)             |           - |
| [B-tree][b-tree]                             | [btree](btree)                       |           - |
| [Undirected graph][undirected-graph]         | [graph/undirected](graph/undirected) |           - |
| [Directed graph][directed-graph]             | [graph/directed](graph/directed)     |           - |

### Sorting algorithms

| Name (Wikipedia link)                        | Package                                    | Implemented |
|:---------------------------------------------|:-------------------------------------------|:-----------:|
| [Quicksort][quick-sort]                      | [sort/quick](sort/quick)                   |           - |
| [Bubble sort][bubble-sort]                   | [sort/bubble](sort/bubble)                 |           - |
| [Merge sort][merge-sort]                     | [sort/merge](sort/merge)                   |           - |
| [Heapsort][heap-sort]                        | [sort/heap](sort/heap)                     |           - |
| [Selection sort][selection-sort]             | [sort/selection](sort/selection)           |           - |
| [Insertion sort][insertion-sort]             | [sort/insertion](sort/insertion)           |           - |
| [Cocktail shaker sort][cocktail-shaker-sort] | [sort/cocktailshaker](sort/cocktailshaker) |           - |
| [Shellsort][shell-sort]                      | [sort/shell](sort/shell)                   |           - |
| [Radix sort][radix-sort]                     | [sort/radix](sort/radix)                   |           - |

### Hashing algorithms

| Name (Wikipedia link)                        | Package                      | Implemented |
|:---------------------------------------------|:-----------------------------|:-----------:|
| [SHA-256][sha-256]                           | [hash/sha256](hash/sha256)   |           - |
| [MD5][md5]                                   | [hash/md5](hash/md5)         |           - |
| [Pearson hashing][pearson-hashing]           | [hash/pearson](hash/pearson) |           ✓ |
| [CRC-32][crc-32]                             | [hash/crc32](hash/crc32)     |           - |
| [MurmurHash][murmur-hash]                    | [hash/murmur](hash/murmur)   |           - |
| [SpookyHash][spooky-hash]                    | [hash/spooky](hash/spooky)   |           - |

### String metrics

| Name (Wikipedia link)                        | Package                                              | Implemented |
|:---------------------------------------------|:-----------------------------------------------------|:-----------:|
| [Levenshtein][levenshtein]                   | [stringmetric/levenshtein](stringmetric/levenshtein) |           - |
| [Hamming][hamming]                           | [stringmetric/hamming](stringmetric/hamming)         |           - |
| [N-Gram][n-gram]                             | [stringmetric/ngram](stringmetric/ngram)             |           - |

### Utilities

| Function                                     | Package                    | Implemented |
|:---------------------------------------------|:---------------------------|:-----------:|
| Mix                                          | [util](util)               |           ✓ |
| Max                                          | [util](util)               |           ✓ |

[golang]: https://golang.org/ref/spec
[gopher-logo]: assets/gopher-mini.png "Gopher logo"
[gopherizeme]: https://gopherize.me/gopher/692fbfd019724e297a7b0761bd3e9697ef8e6bee
[list]: https://en.wikipedia.org/wiki/List_(abstract_data_type)
[sorted-list]: https://en.wikipedia.org/wiki/Sorted_list
[linked-list]: https://en.wikipedia.org/wiki/Linked_list
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
[sha-256]: https://en.wikipedia.org/wiki/SHA-2
[md5]: https://en.wikipedia.org/wiki/MD5
[crc-32]: https://en.wikipedia.org/wiki/Cyclic_redundancy_check#CRC-32_algorithm
[murmur-hash]: https://en.wikipedia.org/wiki/MurmurHash
[spooky-hash]: https://en.wikipedia.org/wiki/Jenkins_hash_function#SpookyHash
[pearson-hashing]: https://en.wikipedia.org/wiki/Pearson_hashing
[undirected-graph]: https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)#Graph
[directed-graph]: https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)#Directed_graph
[tree]: https://en.wikipedia.org/wiki/Tree_(data_structure)
[trie]: https://en.wikipedia.org/wiki/Trie
[binary-tree]: https://en.wikipedia.org/wiki/Binary_tree
[b-tree]: https://en.wikipedia.org/wiki/B-tree
[levenshtein]: https://en.wikipedia.org/wiki/Levenshtein_distance
[hamming]: https://en.wikipedia.org/wiki/Hamming_distance
[n-gram]: https://en.wikipedia.org/wiki/N-gram
