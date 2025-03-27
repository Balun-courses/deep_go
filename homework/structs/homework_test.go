package main

import (
	"fmt"
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = name
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = int32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mana = uint16(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.health = uint16(health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respect = uint8(respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strength = uint8(strength)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experience = uint8(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.level = uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasHouse = true
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasGun = true
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasFamilty = true
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.personType = uint8(personType)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

// GamePerson выравнивание 8 байт, потому что у строка это заголовок вида (ptr, len), то есть 8 + 8 байт
type GamePerson struct {
	hasHouse, hasGun, hasFamilty                     bool  // выравнивание на 8
	personType, respect, strength, experience, level uint8 // пойдет вместе булами выше в одно выравнивание, видимо потому что все однобайтовое

	mana, health  uint16 // новые выравнивание на 8, offset 8 и 10 соотвественно
	x, y, z, gold int32  // займет 2 раза по 8 байт, offsets 12, 16, 20, 24 соответственно
	name          string // займет 2 раза по 8 байт, offset 32

	// итоговый размер будет 48
}

func NewGamePerson(options ...Option) GamePerson {
	p := GamePerson{}

	for _, opt := range options {
		opt(&p)
	}

	return p
}

func (p *GamePerson) Name() string {
	return p.name
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.mana)
}

func (p *GamePerson) Health() int {
	return int(p.health)
}

func (p *GamePerson) Respect() int {
	return int(p.respect)
}

func (p *GamePerson) Strength() int {
	return int(p.strength)
}

func (p *GamePerson) Experience() int {
	return int(p.experience)
}

func (p *GamePerson) Level() int {
	return int(p.level)
}

func (p *GamePerson) HasHouse() bool {
	return p.hasHouse
}

func (p *GamePerson) HasGun() bool {
	return p.hasGun
}

func (p *GamePerson) HasFamilty() bool {
	return p.hasFamilty
}

func (p *GamePerson) Type() int {
	return int(p.personType)
}

func TestGamePerson(t *testing.T) {

	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	fmt.Println("Sizeof:", unsafe.Sizeof(GamePerson{}))
	fmt.Println("Alignof:", unsafe.Alignof(GamePerson{}))
	fmt.Println()

	gp := GamePerson{}
	fmt.Println("Offsetof(hasHouse):", unsafe.Offsetof(gp.hasHouse))
	fmt.Println("Offsetof(hasGun):", unsafe.Offsetof(gp.hasGun))
	fmt.Println("Offsetof(hasFamilty):", unsafe.Offsetof(gp.hasFamilty))
	fmt.Println()
	fmt.Println("Offsetof(personType):", unsafe.Offsetof(gp.personType))
	fmt.Println("Offsetof(respect):", unsafe.Offsetof(gp.respect))
	fmt.Println("Offsetof(strength):", unsafe.Offsetof(gp.strength))
	fmt.Println("Offsetof(experience):", unsafe.Offsetof(gp.experience))
	fmt.Println("Offsetof(level):", unsafe.Offsetof(gp.level))
	fmt.Println()
	fmt.Println("Offsetof(mana):", unsafe.Offsetof(gp.mana))
	fmt.Println("Offsetof(health):", unsafe.Offsetof(gp.health))
	fmt.Println()
	fmt.Println("Offsetof(x):", unsafe.Offsetof(gp.x))
	fmt.Println("Offsetof(y):", unsafe.Offsetof(gp.y))
	fmt.Println("Offsetof(z):", unsafe.Offsetof(gp.z))
	fmt.Println("Offsetof(gold):", unsafe.Offsetof(gp.gold))
	fmt.Println()
	fmt.Println("Offsetof(name):", unsafe.Offsetof(gp.name))

	const x, y, z = math.MinInt32, math.MaxInt32, 0
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 1000
	const health = 1000
	const respect = 10
	const strength = 10
	const experience = 10
	const level = 10

	options := []Option{
		WithName(name),
		WithCoordinates(x, y, z),
		WithGold(gold),
		WithMana(mana),
		WithHealth(health),
		WithRespect(respect),
		WithStrength(strength),
		WithExperience(experience),
		WithLevel(level),
		WithHouse(),
		WithFamily(),
		WithType(personType),
	}

	person := NewGamePerson(options...)
	assert.Equal(t, name, person.Name())
	assert.Equal(t, x, person.X())
	assert.Equal(t, y, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, gold, person.Gold())
	assert.Equal(t, mana, person.Mana())
	assert.Equal(t, health, person.Health())
	assert.Equal(t, respect, person.Respect())
	assert.Equal(t, strength, person.Strength())
	assert.Equal(t, experience, person.Experience())
	assert.Equal(t, level, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamilty())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
