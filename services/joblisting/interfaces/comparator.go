package interfaces

// Comparator compares whether two models are equal
type Comparator interface {
	IsEqual(m2 interface{}) bool
}
