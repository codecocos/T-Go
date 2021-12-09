package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//1.Account : lock
type Account struct {
	balance int
	mutex *sync.Mutex
}

func (a *Account) Withdraw(val int) {
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	//락을 했으면 언락을 해줘야 한다.
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()

	return balance
}

var accounts []*Account
//2. 출금 예금 : lock
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Withdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()
}

func GetTotalBalance() int {
	globalLock.Lock()
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	globalLock.Unlock()
	return total
}

func RandomTransfer() {
	var sender, balance int
	for {
		//0이상 n미만의 숫자를 랜덤으로 넣어준다.
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		} 
	}

	var receiver int
	for{
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender,receiver,money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n",GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000,mutex: &sync.Mutex{}})
	}
	
	globalLock = &sync.Mutex{}

	PrintTotalBalance()

	for i := 0; i < 10; i++ {
		//스레드가 10개가 생긴다
		go GoTransfer()
	}

	//이거 자체도 스레드
	for {
		PrintTotalBalance()
		time.Sleep(100 *time.Millisecond)
	}
}

// a.balance -= val
// a.balance = a.balance - val