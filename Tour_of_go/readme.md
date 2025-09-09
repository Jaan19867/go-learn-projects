



# Interface 
Value receiver (func (T) Method) → both T and *T implement the interface.

Pointer receiver (func (*T) Method) → only *T implements the interface, not T.

Type assertion is a way to retrieve the underlying concrete value stored inside an interface.

# Stringers 

fmt.Sprintf does not print anything it just format a string as per defination 
String() function need to return string (related to tour of go )