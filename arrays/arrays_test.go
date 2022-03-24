package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	arrInt := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 15, Sum(arrInt))

	arrInt8 := []int8{1, 2, 3, 4, 5}
	assert.Equal(t, int8(15), Sum(arrInt8))

	arrInt16 := []int16{1, 2, 3, 4, 5}
	assert.Equal(t, int16(15), Sum(arrInt16))

	arrInt32 := []int32{1, 2, 3, 4, 5}
	assert.Equal(t, int32(15), Sum(arrInt32))

	arrInt64 := []int64{1, 2, 3, 4, 5}
	assert.Equal(t, int64(15), Sum(arrInt64))

	arrUint := []uint{1, 2, 3, 4, 5}
	assert.Equal(t, uint(15), Sum(arrUint))

	arrUint8 := []uint8{1, 2, 3, 4, 5}
	assert.Equal(t, uint8(15), Sum(arrUint8))

	arrUint16 := []uint16{1, 2, 3, 4, 5}
	assert.Equal(t, uint16(15), Sum(arrUint16))

	arrUint32 := []uint32{1, 2, 3, 4, 5}
	assert.Equal(t, uint32(15), Sum(arrUint32))

	arrUint64 := []uint64{1, 2, 3, 4, 5}
	assert.Equal(t, uint64(15), Sum(arrUint64))

	arrFloat32 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	assert.Equal(t, float32(15.0), Sum(arrFloat32))

	arrFloat64 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	assert.Equal(t, 15.0, Sum(arrFloat64))
}

func TestContains(t *testing.T) {
	arrString := []string{"Lahmacun", "Köfte", "Turşu"}
	assert.Equal(t, true, Contains(arrString, "Köfte"))
	assert.Equal(t, false, Contains(arrString, "Adana Dürüm"))

	arrInt := []int{2, 4, 6, 8, 10}
	assert.Equal(t, true, Contains(arrInt, 2))
	assert.Equal(t, false, Contains(arrInt, 3))
}

func TestRemove(t *testing.T) {
	arrString := []string{"Lahmacun", "Köfte", "Turşu"}
	resultArrayString := []string{"Lahmacun", "Köfte"}
	assert.Equal(t, resultArrayString, Remove(arrString, "Turşu"))

	arrInt := []int{2, 4, 6, 8, 10}
	resultArrayInt := []int{2, 4, 6, 10}
	assert.Equal(t, resultArrayInt, Remove(arrInt, 8))
}

func TestFindIndex(t *testing.T) {
	arrString := []string{"Lahmacun", "Köfte", "Turşu"}
	assert.Equal(t, 2, FindIndex(arrString, "Turşu"))

	arrInt := []int{2, 4, 6, 8, 10}
	assert.Equal(t, 3, FindIndex(arrInt, 8))
	assert.Equal(t, -1, FindIndex(arrInt, 12))
}

func TestEquals(t *testing.T) {
	arrString1 := []string{"Lahmacun", "Köfte", "Turşu"}
	arrString2 := []string{"Lahmacun", "Köfte", "Turşu"}
	arrString3 := []string{"Lahmacun", "Köfte"}

	assert.Equal(t, true, Equals(arrString1, arrString2))
	assert.Equal(t, false, Equals(arrString1, arrString3))

	arrInt1 := []int{2, 4, 6, 8, 10}
	arrInt2 := []int{2, 4, 6, 8, 10}
	arrInt3 := []int{2, 6, 8, 10}
	assert.Equal(t, true, Equals(arrInt1, arrInt2))
	assert.Equal(t, false, Equals(arrInt2, arrInt3))
}

func TestEqualsItems(t *testing.T) {
	arrString1 := []string{"Lahmacun", "Köfte", "Turşu"}
	arrString2 := []string{"Turşu", "Köfte", "Lahmacun"}
	arrString3 := []string{"Lahmacun", "Köfte"}

	assert.Equal(t, true, EqualsItems(arrString1, arrString2))
	assert.Equal(t, false, EqualsItems(arrString1, arrString3))

	arrInt1 := []int{2, 4, 6, 8, 10}
	arrInt2 := []int{10, 8, 6, 4, 2}
	arrInt3 := []int{2, 6, 8, 10}
	assert.Equal(t, true, EqualsItems(arrInt1, arrInt2))
	assert.Equal(t, false, EqualsItems(arrInt2, arrInt3))
}
