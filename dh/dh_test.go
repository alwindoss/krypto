package dh

import (
	"math/big"
	"testing"
)

func Test(t *testing.T) {
	AlicePrivKey, _ := new(big.Int).SetString("10000000", 16)
	BobPrivKey, _ := new(big.Int).SetString("89898989", 16)
	// AlicePrivKey, _ := new(big.Int).SetString("87e0beefd8122561e9c329d764c6e3b3dafe538a", 16)
	// BobPrivKey, _ := new(big.Int).SetString("4fc9904887ac7fabff87f054003547c2d9458c1f6f584c140d7271f8b266bb390af7e3f625a629bec9c6a057a4cbe1a556d5e3eb2ff1c6ff677a08b0c7c509110b9e7c6dbc961ca4360362d3dbcffc5bf2bb7207e0a5922f77cf5464b316aa49fb62b338ebcdb30bf573d07b663bb7777b69d6317df0a4f636ba3d9acbf9e8ac", 16)
	alice, err := New(GroupFourteen, AlicePrivKey)
	if err != nil {
		t.FailNow()
	}
	alice.Generate()
	alicePubKey := alice.PublicKey()
	bob, err := New(GroupFourteen, BobPrivKey)
	if err != nil {
		t.FailNow()
	}
	bob.Generate()
	bobPubKey := bob.PublicKey()
	bobSecretKey, _ := bob.SecretKey(alicePubKey)
	aliceSecretKey, _ := alice.SecretKey(bobPubKey)
	if bobSecretKey.String() != aliceSecretKey.String() {
		t.FailNow()
	}
	// fmt.Printf("Alice Secret Key: %s\n", aliceSecretKey.String())
	// fmt.Printf("Bob Secret Key: %s\n", bobSecretKey.String())
	// t.FailNow()
}
