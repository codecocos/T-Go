package main

// type Player struct {
// }
// type Monster struct {
// }
// type chest struct {
// }
// func (p *Player) Attack(m *Monster) {
// }
// func (p *Player) AttackP(m *Monster) {
// }
// func (m *Monster) AttackP(p *Player) {
// }
// func (p *Monster) Attack(m *Monster) {}
// }

//인터페이스에 의존하는 경우

type Player struct {
}

func (p *Player)Attack(target *BeAttackable)  {
	// ...
	// 데미지를 계산해서
	//target.BeAttacked(dmg)
}

type Monster struct {
}

func (m *Monster)Attack(target *BeAttackable)  {
	// ...
	// 데미지를 계산해서
	//target.BeAttacked(dmg)
}


type chest struct {
}

type Attackable interface {
	Attack(BeAttackable)
}

type BeAttackable interface {
	BeAttacked(int)
}

func Attack(attacker *Attackable, defender *BeAttackable) {
	attacker.Attack(defender)
}
