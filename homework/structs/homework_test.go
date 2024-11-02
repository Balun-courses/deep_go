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
		// need to implement
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	// need to implement
}

func NewGamePerson(options ...Option) GamePerson {
	// need to implement
	return GamePerson{}
}

func (p *GamePerson) Name() string {
	// need to implement
	return ""
}

func (p *GamePerson) X() int {
	// need to implement
	return 0
}

func (p *GamePerson) Y() int {
	// need to implement
	return 0
}

func (p *GamePerson) Z() int {
	// need to implement
	return 0
}

func (p *GamePerson) Gold() int {
	// need to implement
	return 0
}

func (p *GamePerson) Mana() int {
	// need to implement
	return 0
}

func (p *GamePerson) Health() int {
	// need to implement
	return 0
}

func (p *GamePerson) Respect() int {
	// need to implement
	return 0
}

func (p *GamePerson) Strength() int {
	// need to implement
	return 0
}

func (p *GamePerson) Experience() int {
	// need to implement
	return 0
}

func (p *GamePerson) Level() int {
	// need to implement
	return 0
}

func (p *GamePerson) HasHouse() bool {
	// need to implement
	return false
}

func (p *GamePerson) HasGun() bool {
	// need to implement
	return false
}

func (p *GamePerson) HasFamilty() bool {
	// need to implement
	return false
}

func (p *GamePerson) Type() int {
	// need to implement
	return 0
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
