# Go Slice Modification in Functions
This example demonstrates a common pitfall in Go when working with slices: modifying a slice within a function can unexpectedly modify the original slice. 

## The Bug
The `myFunc` function increments each element of the input slice.  However, because slices are references, this modification is reflected in the original slice `x`.

## The Solution
To avoid this, you can create a copy of the slice within the function before modifying it.