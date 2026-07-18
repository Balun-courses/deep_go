package main

import (
	"bytes"
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		copy(person.name[:], []byte(name))
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
		person.all = person.all | uint64(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.all = person.all | (uint64(mana) << 32)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.all = person.all | (uint64(health) << 42)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.all = person.all | (uint64(respect) << 52)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.all = person.all | (uint64(strength) << 56)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.all = person.all | (uint64(experience) << 60)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.combo = person.combo | uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.all = person.all | (uint64(1) << 31)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.combo = person.combo | (uint8(1) << 4)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.combo = person.combo | (uint8(1) << 5)
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.combo = person.combo | (uint8(personType) << 6)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	// low -> high
	// gold = 31 bit
	// hasHouse = 1 bit
	// mana = 10 bit
	// health = 10 bit
	// respect 4 bit
	// stength 4 bit
	// exp 4 bit
	all uint64
	x   int32
	y   int32
	z   int32

	combo uint8 // lo->high level(4 bits) weapon(v) family(v) type(v) (2 bits and 1 bit remaining fields)
	name  [43]byte
}

func NewGamePerson(options ...Option) GamePerson {
	p := GamePerson{}
	for _, op := range options {
		op(&p)
	}
	return p
}

func (p *GamePerson) Name() string {
	cleanBytes := bytes.TrimRight(p.name[:], "\x00")
	return string(cleanBytes)
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
	return int(p.all & 0x000000007FFFFFFF)
}

func (p *GamePerson) Mana() int {
	return int((p.all & 0x000003FF00000000) >> 32)
}

func (p *GamePerson) Health() int {
	return int((p.all & 0x000FFC0000000000) >> 42)
}

func (p *GamePerson) Respect() int {
	return int((p.all & 0x00F0000000000000) >> 52)
}

func (p *GamePerson) Strength() int {
	return int((p.all & 0x0F00000000000000) >> 56)
}

func (p *GamePerson) Experience() int {
	return int((p.all & 0xF000000000000000) >> 60)
}

func (p *GamePerson) Level() int {
	return int(p.combo & 0x0F)
}

func (p *GamePerson) HasHouse() bool {
	return ((p.all & 0x0000000080000000) >> 31) == 1
}

func (p *GamePerson) HasGun() bool {
	return p.combo&0x10>>4 == 1
}

func (p *GamePerson) HasFamilty() bool {
	return p.combo&0x20>>5 == 1
}

func (p *GamePerson) Type() int {
	return int(p.combo & 0xC0 >> 6)
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
