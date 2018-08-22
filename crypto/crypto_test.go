package crypto

import (
	"encoding/hex"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Errorf(Random(32))
}

func TestGCM(t *testing.T) {
	key := []byte("2pdMiskZXe928M6T7TqmjwoQBhhVcJiv")
	plainText := "this is test"
	cipherText, err := EncryptByGCM(key, plainText)
	if err != nil {
		t.Fatal("failed to encrypt:", err)
	}
	res, err := DecryptByGCM(key, cipherText)
	if err != nil {
		t.Fatal("failed to decrypt:", err)
	}
	if res != plainText {
		t.Fatal("plain text was not match:", res, plainText)
	}
}

func BenchmarkEncryptGCM(b *testing.B) {
	key := []byte("2pdMiskZXe928M6T7TqmjwoQBhhVcJiv")
	plainText := "this is test"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := EncryptByGCM(key, plainText); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecryptGCM(b *testing.B) {
	key := []byte("2pdMiskZXe928M6T7TqmjwoQBhhVcJiv")
	cipherText, _ := hex.DecodeString("52091fc0409a889735e8860450b116442839617ed993b6f072df29470958c8023547ec48169faf20")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := DecryptByGCM(key, cipherText); err != nil {
			b.Fatal(err)
		}
	}
}
