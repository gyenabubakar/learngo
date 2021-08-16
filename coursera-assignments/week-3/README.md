# Solution for *Peer-graded Assignment: Module 3 Activity*, on Coursera

### Instructions
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

### Solution
> A CLI program which accepts an input of 12 numbers separated with a comma, parses it into a multi-dimensional slice of 4 int slices (each containing 3 elements), concurrently sorts them all separately, and then together.
> 
You can find the coded solution to this problem written in the `main.go` file.
