// Copyright 2020 Alexander Fahlke.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
Package queue provides a simple implementation of a concurrency safe queue.
This is a FIFO (first-in, first-out) data structure.

Capacity

Functions to get information about various queue capacity information.
Time complexity is constant O(1).
    Empty(), Size()

Modifiers

Functions that modify the state of the queue.
    Push(), Pop()

Element access

Functions to peek into the queue without modifying it.
Time complexity is constant O(1).
    Front(), Back()
*/
package queue
