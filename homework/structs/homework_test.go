package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const (
	Type      = 7
	HasFamily = 9
	HasGun    = 10
	Mana      = 22
	Health    = 12
	HasHouse  = 11
)
const (
	Experience = 4
	Strength   = 8
	Respect    = 12
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > 42 {
			name = name[:42]
		}
		copy(person.name[:], name)
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
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.attributs |= uint32(mana) << Mana
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.attributs |= uint32(health) << Health
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.properties |= uint16(respect) << Respect
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.properties |= uint16(strength) << Strength
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.properties |= uint16(experience) << Experience
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.properties |= uint16(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.attributs |= 1 << HasHouse
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.attributs |= 1 << HasGun
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.attributs |= 1 << HasFamily
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.attributs |= uint32(personType) << Type
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	x          int32
	y          int32
	z          int32
	gold       uint32
	attributs  uint32
	properties uint16
	name       [42]byte
}

func NewGamePerson(options ...Option) GamePerson {
	p := GamePerson{}
	for _, option := range options {
		option(&p)
	}
	return p
}

func (p *GamePerson) Name() string {
	return unsafe.String(&p.name[0], len(p.name))
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
	return int(p.attributs >> Mana)
}

func (p *GamePerson) Health() int {
	return int((p.attributs >> Health) & 0x3FF)
}

func (p *GamePerson) Respect() int {
	return int(p.properties >> Respect)
}

func (p *GamePerson) Strength() int {
	return int((p.properties >> Strength) & 0xF)
}

func (p *GamePerson) Experience() int {
	return int((p.properties >> Experience) & 0xF)
}

func (p *GamePerson) Level() int {
	return int(p.properties & 0xF)
}

func (p *GamePerson) HasHouse() bool {
	return (p.attributs>>HasHouse)&0b1 == 1
}

func (p *GamePerson) HasGun() bool {
	return (p.attributs>>HasGun)&0b1 == 1
}

func (p *GamePerson) HasFamilty() bool {
	return (p.attributs>>HasFamily)&0b1 == 1
}

func (p *GamePerson) Type() int {
	return int((p.attributs >> Type) & 0b11)
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

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
