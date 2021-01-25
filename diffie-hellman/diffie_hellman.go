package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a random private key
func PrivateKey(p *big.Int) *big.Int {
	n, _ := rand.Int(rand.Reader, p)
	min := big.NewInt(2)
	if n.Cmp(min) < 0 {
		n = n.Add(n, min.Sub(min, n))
	}
	return n
}

// PublicKey calculates a public key based on the private key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair generates a new private/public keypair
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return private, public
}

// SecretKey calculates a secret key from a private and a public key
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
