package vendingMachine

type VendingMachine struct {
	insertedMoney int
	coins 		map[string]int
}

func NewVendingMachine(coins, items map[string]int) *VendingMachine {
	return &VendingMachine{coins: coins, items: items}
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine) SelectSD() string {
	m.insertedMoney = 0
	return "SD"
}

func (m *VendingMachine) SelectCC() string {
	m.insertedMoney = 0
	return "CC"
}

func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "0": 1}
	vm := VendingMachine{coins: coins}
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 17
	can := vm.SelectSD()
	fmt.Println(can) //SD
	vm.InsertCoin("T")
	vm.InsertCoin("TW")
	fmt.Println("Insert Money:", vm.InsertedMoney())
	// Inserted Money: 12
	can = vm.SelectCC()
	fmt.Println(can) // CC
}