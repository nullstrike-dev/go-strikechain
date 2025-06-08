package strike

type Block struct {
	Index        uint64
	Timestamp    int64
	Transactions []Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        uint64
	Difficulty   uint8
}

type Blockchain struct {
	Chain      []Block
	Difficulty uint8
}

func NewBlockchain(difficulty uint8) *Blockchain {
	return &Blockchain{
		Difficulty: difficulty,
	}
}
