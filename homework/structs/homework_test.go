package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		for i := range person.name {
			person.name[i] = 0
		}
		copy(person.name[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x, person.y, person.z = int32(x), int32(y), int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = int32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		if mana < 0 {
			mana = 0
		}
		if mana > 1000 {
			mana = 1000
		}
		person.setBits(offMana, widthMana, uint64(mana))
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		if health < 0 {
			health = 0
		}
		if health > 1000 {
			health = 1000
		}
		person.setBits(offHealth, widthHealth, uint64(health))
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		if respect < 0 {
			respect = 0
		}
		if respect > 10 {
			respect = 10
		}
		person.setBits(offRespect, widthSmall, uint64(respect))
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		if strength < 0 {
			strength = 0
		}
		if strength > 10 {
			strength = 10
		}
		person.setBits(offStrength, widthSmall, uint64(strength))
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		if experience < 0 {
			experience = 0
		}
		if experience > 10 {
			experience = 10
		}
		person.setBits(offExperience, widthSmall, uint64(experience))
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		if level < 0 {
			level = 0
		}
		if level > 10 {
			level = 10
		}
		person.setBits(offLevel, widthSmall, uint64(level))
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.setBits(offHouse, 1, 1)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.setBits(offGun, 1, 1)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.setBits(offFamily, 1, 1)
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		if personType < 0 {
			personType = 0
		}
		if personType > 3 {
			personType = 3
		}
		person.setBits(offType, widthType, uint64(personType))
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
	widthMana   = 10
	widthHealth = 10
	widthSmall  = 4
	widthType   = 2

	offMana       = 0
	offHealth     = offMana + widthMana
	offRespect    = offHealth + widthHealth
	offStrength   = offRespect + widthSmall
	offExperience = offStrength + widthSmall
	offLevel      = offExperience + widthSmall
	offHouse      = offLevel + widthSmall
	offGun        = offHouse + 1
	offFamily     = offGun + 1
	offType       = offFamily + 1
)

func mask(width uint) uint64 {
	return (uint64(1) << width) - 1
}

func (p *GamePerson) getPacked() uint64 {
	return uint64(p.packed[0]) |
		uint64(p.packed[1])<<8 |
		uint64(p.packed[2])<<16 |
		uint64(p.packed[3])<<24 |
		uint64(p.packed[4])<<32 |
		uint64(p.packed[5])<<40
}

func (p *GamePerson) setPacked(v uint64) {
	p.packed[0] = byte(v)
	p.packed[1] = byte(v >> 8)
	p.packed[2] = byte(v >> 16)
	p.packed[3] = byte(v >> 24)
	p.packed[4] = byte(v >> 32)
	p.packed[5] = byte(v >> 40)
}

func (p *GamePerson) setBits(off, width uint, val uint64) {
	all := p.getPacked()
	m := mask(width) << off
	all = (all &^ m) | ((val & mask(width)) << off)
	p.setPacked(all)
}

func (p *GamePerson) getBits(off, width uint) uint64 {
	return (p.getPacked() >> off) & mask(width)
}

type GamePerson struct {
	name   [42]byte
	packed [6]byte

	x    int32
	y    int32
	z    int32
	gold int32
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}
	for _, option := range options {
		option(&person)
	}
	return person
}

func (p *GamePerson) Name() string {
	i := 0
	for ; i < len(p.name); i++ {
		if p.name[i] == 0 {
			break
		}
	}
	return string(p.name[:i])
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
	return int(p.getBits(offMana, widthMana))
}

func (p *GamePerson) Health() int {
	return int(p.getBits(offHealth, widthHealth))
}

func (p *GamePerson) Respect() int {
	return int(p.getBits(offRespect, widthSmall))
}

func (p *GamePerson) Strength() int {
	return int(p.getBits(offStrength, widthSmall))
}

func (p *GamePerson) Experience() int {
	return int(p.getBits(offExperience, widthSmall))
}

func (p *GamePerson) Level() int {
	return int(p.getBits(offLevel, widthSmall))
}

func (p *GamePerson) HasHouse() bool {
	return p.getBits(offHouse, 1) != 0
}

func (p *GamePerson) HasGun() bool {
	return p.getBits(offGun, 1) != 0
}

func (p *GamePerson) HasFamilty() bool {
	return p.getBits(offFamily, 1) != 0
}

func (p *GamePerson) Type() int {
	return int(p.getBits(offType, widthType))
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
