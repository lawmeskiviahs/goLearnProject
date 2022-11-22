package slices

func main () {
/*
slices are parts of arrays, like in rust slices contain pointer to an array, length (current length) and capacity (maximum length) of the slice
slices do not copy the array data, rather it creates a slice value that points to the original array hence slices can be manipulated just like arrays
to grow slices we must create a new slice and then cpoly the old slice into thsi new and then assigning its value to the old slice
slices can also be grown bu using the AppendByte(slice, data1, .... datan) fn and also the append(slice, data1, ..... datan) fn. the append function is good for most cases
append() will also append one slice into another
if a and b are two slices, they can be appended by using append(a, b...) this means to use the following indices of the b slice to append them at the end of the slice a 
*/
}