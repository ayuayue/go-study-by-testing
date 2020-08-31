package point_err

import (
    "errors"
    "fmt"
    "testing"
)
type Wallet struct {
    balance int
}
type stringer interface {
    String()string
}
type Bitcoin int
func (b *Bitcoin) String() string {
    return fmt.Sprintf("%d BTC",b)
}
func (w *Wallet) Deposit (amount int){
    fmt.Printf("address of balance in Deposit is %v\n",&w.balance)
    w.balance += amount
}
func (w *Wallet) Balance() int {
    return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
    if amount > w.balance {
        return errors.New("oh no")
    }
    w.balance -= amount
    return nil
}
func TestWallet(t *testing.T){
    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
       assertBalance(t,wallet,Bitcoin(10))
    })
    t.Run("Withdraw with funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(10))
        assertBalance(t,wallet,Bitcoin(10))
        assertError(t,err)
    })
    t.Run("Withdraw insufficient funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(100))

        assertBalance(t,wallet,Bitcoin(20))
        assertError(t,err,ErrInsufficientFunds)
    })
}
func assertError(t *testing.T,got error,want error){
    t.Helper()
    if got == nil {
        t.Errorf("wanted an error but didn't get one")
    }
    if got != want {
        t.Errorf("got %s want %s",got,want)
    }
}
func assertBalance(t *testing.T,wallet Wallet,want Bitcoin){
    t.Helper()
    got := wallet.Balance()
    if got != want {
        t.Errorf("got %s want %s",got,want)
    }
})