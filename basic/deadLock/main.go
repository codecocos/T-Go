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
	//2. lock 작게 잡기
	//a.mutex.Lock()
	a.balance -= val
	//a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	//a.mutex.Lock()
	a.balance += val
	//a.mutex.Unlock()
}

func (a *Account) Balance() int {
	//락을 했으면 언락을 해줘야 한다.
	balance := a.balance

	return balance
}

var accounts []*Account
//3. lock 크게 잡기 
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {

	// // deadlock 확인하기 위한 코드
	// accounts[sender].mutex.Lock()
	// fmt.Println("Lock", sender)
	// time.Sleep(1000 * time.Millisecond)
	// accounts[receiver].mutex.Lock()
	// fmt.Println("Lock", receiver)

	globalLock.Lock()
	//출금과 송금 사이에 다른 스레드가 끼어들 수 있음으로 송금인과 수신인에 미리 락을 건다.
	accounts[sender].Withdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()

	// // deadlock 확인하기 위한 코드
	// accounts[sender].mutex.Unlock()
	// accounts[receiver].mutex.Unlock()
	
	fmt.Println("Transfer", sender, receiver, money)
}

func GetTotalBalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
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
	
	go func ()  {
		for {
			//0번에서 1번으로 100원 송금
			Transfer(0,1,100)
		}
	}()

	go func ()  {
		for {
			//1번에서 0번으로 100원 송금
			Transfer(1,0,100)
		}
	}()

	//이거 자체도 스레드
	for {
		time.Sleep(100 *time.Millisecond)
	}
}

/*
	Transfer 0 1 100
	Transfer 1 0 100

	cpu 1 | Lock 0 -> Lock 2 : 대기중
	cpu 2 | Lock 1 -> Lcok 1 : 대기중

	서로 계속 대기중인 상태가 됨 -> Dead Lock 상태


*/
