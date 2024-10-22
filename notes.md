# GoLang Tutorial


used for multi-threaded applications capable of running on multiple cores

compiled language

scalable and fast systems


go needs 3 workspaces kind of like folders
- src - stores source files
- bin - stores executables
- pkg - stores compiled packages

the length of an array is part of the array's type, so you can't assign a variable to determine the size of the array

that is why we use slices

using [...] makes an array and using [] makes a slice

slices can't be compared with each other using operator

go is a call by value language, so each in-built function returns a copy that must be assigned otherwise it is a compile time error


make is an inbuilt function that specifies the type of the variable to be made , the size and the capacity for the same

clear() function is used to set all the values in an array to zero values

when you use the slice expression and not the data type slice, you are not making a copy rather you are sharing the same memory point

copy function to create a copy of the slice you want to use


Arrays can be converted into a slice by the expression of slicing

xArrays = [4]int{1,2,3,4}
xSlices = xArrays[:]

