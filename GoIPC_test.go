package GoIPC

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestRun(t *testing.T) {
	Create("GoTestMemory", 64)
	Open("GoTestMemory", 64)

	for i := 0; i < 100; i++ {
		test := "hello" + strconv.Itoa(i)
		WriteMemoryString(test)

		result := ReadMemoryString()
		if result != test {
			t.Errorf("Memory test failed wanted %s got %s", test, result)
		}
		buf := ReadMemory(int(unsafe.Sizeof(test)), 0)

		if buf[0] != test[0] {
			t.Errorf("Memory buf test failed wanted %d got %d", test[0], buf[0])

		}

		buf2 := ReadMemory(int(unsafe.Sizeof(test)), 1)

		if buf2[0] != test[1] {
			t.Errorf("Memory buf test 2 failed wanted %d got %d", test[1], buf2[0])

		}

	}
	byteArray := []byte{97, 98, 99, 100, 101, 102}
	WriteMemory(byteArray)

	bTest := ReadMemory(int(unsafe.Sizeof(byteArray)), 0)
	for y := range byteArray {
		if byteArray[y] != bTest[y] {
			t.Errorf("Memory r/w bytes test failed wanted %d got %d", byteArray[y], bTest[y])

		}
	}
}
