package main

func main() {
	bc := NewBlockchain()
	bc.PrintLastBlock()

	bc.AddBlock("Send 1 BTC to Petya")
	bc.PrintLastBlock()

	bc.AddBlock("Send 2 BTC to Alex")
	bc.PrintLastBlock()

}
