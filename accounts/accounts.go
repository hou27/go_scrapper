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

type Thing struct {
    Name  string
    Num   int
}

func NewThing(someParameter string) *Thing {
    p := new(Thing)
    p.Name = someParameter
    p.Num = 33 // <- a very sensible default value
    return p
}