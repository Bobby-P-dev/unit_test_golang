package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Testify menyediakan dua package untuk assertion, yaitu assert dan require
/* jika menggunakan assert maka akan me return error fail() dan akan melanjutkan kode program di bawah nya,
jika menggunakan require maka akan me return failNow() dan akan memberhentikan kode program di bawahnya.
*/

func TestMain(m *testing.M) {
	//before testing
	fmt.Println("testing di mulai")

	m.Run()

	//after testing
	fmt.Println("testing selesai")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bobby")
	// assert.Equal untuk mengecek data harus sama
	assert.Equal(t, "Hello Bobby", result)
	fmt.Println("ini kode di bawah fail")
}

func TestHelloDunia(t *testing.T) {
	result := HelloWorld("Jansen")
	// assert.NotEqual untuk mengecek data tidak boleh sama
	assert.NotEqual(t, "Hello jansen", result)
	fmt.Println("ini kode di bawah error")
}

func TestGrady(t *testing.T) {
	result := HelloWorld("Grady")
	// assert.Nil untuk mengecek data harus kosong
	assert.Nil(t, result)
	fmt.Println("ini kode di bawah gagal")
}

func TestFryan(t *testing.T) {
	result := HelloWorld("Bobby")
	// assert.NotNil untuk mengecek data tidak boleh
	assert.NotNil(t, result)
	fmt.Println("ini kode di bawah berhasil")
}

// ini adalah subtest ada func test di dalam func test tersebut yang bisa berisi banyak func atau testing
// penggunaan subtest dengan cara berikut

/*
Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
go test -run TestNamaFunction

Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
go test -run TestNamaFunction/NamaSubTest

Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
go test -run /NamaSubTest
*/

func TestSubTest(t *testing.T) {
	//sub test pertama
	t.Run("Bobby", func(t *testing.T) {
		result := HelloWorld("Bobby")
		require.Equal(t, "Hello Bobby", result, "Reslut must be Hello Bobby")
	})

	// sub test kedua
	t.Run("Grady", func(t *testing.T) {
		result := HelloWorld("Grady")
		require.NotEqual(t, "Hello Grady", result, "result must be Hello Grady")
	})
}

// ini adalah table test
func TestHelloWorldTable(t *testing.T) {
	// table test
	test := []struct {
		name    string
		request string
		expect  string
	}{
		{
			name:    "HelloWorld(Bobby)",
			request: "Bobby",
			expect:  "Hello Bobby",
		},
	}
	for _, tests := range test {
		t.Run(tests.name, func(t *testing.T) {
			result := HelloWorld(tests.request)
			assert.Equal(t, tests.expect, result, "result must be Hello Bobby")
		})
	}
}
