package main

import "fmt"

const (
	MAIN_PORT = "8001"
)

func main() {
	const a = 10
	fmt.Println(a)

	const (
		ipAddress = "127.127.00.03"
		port      = "8000"
		dbName    = "postgres"
	)
	fmt.Println("ipAddress value:", ipAddress)

	fmt.Println(checkPortIsValid())

	const ADMIN_EMAIL string = "admin@admin.com"

	const NUMERIC = 100
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	var numFloat32 float32 = NUMERIC
	var numComplex64 complex64 = NUMERIC

	fmt.Printf("numInt8 value %v type %T\n", numInt8, numInt8)
	fmt.Printf("numInt32 value %v type %T\n", numInt32, numInt32)
	fmt.Printf("numInt64 value %v type %T\n", numInt64, numInt64)
	fmt.Printf("numFloat32 value %v type %T\n", numFloat32, numFloat32)
	fmt.Printf("numComplex64 value %v type %T\n", numComplex64, numComplex64)
	fmt.Printf("%v + %v is %v\n", numInt8, NUMERIC, numInt8+NUMERIC)
}

func checkPortIsValid() bool {
	if MAIN_PORT == "8001" {
		return true
	}
	return false
}
