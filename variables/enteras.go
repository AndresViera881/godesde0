package variables

import (
	"fmt"
)

func MuestraEnteros() {
	var intComun int
	intde32 := int32(2147483647)
	intde64 := int64(9223372036854775807)
	fmt.Println("intComun:", intComun)
	fmt.Println("intde32:", intde32)
	fmt.Println("intde64:", intde64)
}
