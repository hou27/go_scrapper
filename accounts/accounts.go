package accounts

// add comment to export(optional)
// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

/**
 * method
 *
 * method has a struct behind it's name.
 * it is called receiver.
 *
 */

/**
 * Pointer Receiver
 * 
 * receiver의 값을 변경하고자 할 때(단순히 읽기가 아닌 쓰기 작업도 같이)
 * struct의 크기가 커서 deep copy 비용이 클 때 
 * 코드 일관성(option): 어떤 함수가 포인터 receiver를 사용하는 경우 일관성을 줄 때
 */
// Deposit x amount on your account
func (a *Account) Deposit(amount int) { // use pointer receiver
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}