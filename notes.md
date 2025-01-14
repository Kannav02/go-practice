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


the difference between nil and empty data structures is that, nil haven't been initialised at all, while empty ones have been


var sliceNil []int

sliceEmpty := []int{}

sliceEmpty2 := make([]int,2,3)


can't add elements to nil data structures but can add elements to empty data structures


maps are used to map a value to another value


sampleMap[string][]int

or sampleMap := map[string][]int {}


when accessing a value from the map, you can even see if the value existed in the map before


v,ok = sampleMap[someValue]

fmt.Println(v,ok)

it would print the value and ok would contain either true or false based on wether the value exists in the map

structs are implemented as follows

type person struct{
    varName varType
    varName varType
    varName varType
}


there can be anonymous structs as well

they can be used when transforming data from json to be used in go or for test cases

anonStruct := struct{
    name string
    age int
}{
    name: "kannav"
    age: 21
}