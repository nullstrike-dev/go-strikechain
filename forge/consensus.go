package forge

import "github.com/yourusername/go-strikechain/strike"

type ProofOfWork struct {
	Blockchain *strike.Blockchain
}

func NewProofOfWork(bc *strike.Blockchain) *ProofOfWork {
	return &ProofOfWork{
		Blockchain: bc,
	}
}
