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
		// Сохраняем ману в нижних 10 битах
		person.manaHealthFlags |= uint32(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		// Сохраняем здоровье в битах с 10 по 19
		person.manaHealthFlags |= uint32(health) << 10
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		// Сохраняем уважение в младших 4 битах
		person.rselAttrs |= uint16(respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		// Сохраняем силу в битах с 4 по 7
		person.rselAttrs |= uint16(strength) << 4
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		// Сохраняем опыт в битах с 8 по 11
		person.rselAttrs |= uint16(experience) << 8
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		// Сохраняем уровень в битах с 12 по 15
		person.rselAttrs |= uint16(level) << 12
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		// Флаг дома - 20-й бит
		person.manaHealthFlags |= uint32(1) << 20
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		// Флаг оружия - 22-й бит
		person.manaHealthFlags |= uint32(1) << 22
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		// Флаг семьи - 21-й бит
		person.manaHealthFlags |= uint32(1) << 21
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		// Тип игрока - биты с 23 по 30
		person.manaHealthFlags |= uint32(personType) << 23
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	x               int32
	y               int32
	z               int32
	gold            uint32
	manaHealthFlags uint32
	rselAttrs       uint16
	name            [42]byte
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}

	for _, option := range options {
		option(&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	return string(p.name[:])
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
	return int(p.manaHealthFlags & 0x3FF) // Мана в нижних 10 битах
}

func (p *GamePerson) Health() int {
	return int((p.manaHealthFlags >> 10) & 0x3FF) // Здоровье в битах с 10 по 19
}

func (p *GamePerson) Respect() int {
	return int(p.rselAttrs & 0xF) // Уважение в младших 4 битах
}

func (p *GamePerson) Strength() int {
	return int((p.rselAttrs >> 4) & 0xF) // Сила в битах с 4 по 7
}

func (p *GamePerson) Experience() int {
	return int((p.rselAttrs >> 8) & 0xF) // Опыт в битах с 8 по 11
}

func (p *GamePerson) Level() int {
	return int((p.rselAttrs >> 12) & 0xF) // Уровень в битах с 12 по 15
}

func (p *GamePerson) HasHouse() bool {
	return (p.manaHealthFlags & (1 << 20)) != 0
}

func (p *GamePerson) HasGun() bool {
	return (p.manaHealthFlags & (1 << 22)) != 0
}

func (p *GamePerson) HasFamily() bool {
	return (p.manaHealthFlags & (1 << 21)) != 0
}

func (p *GamePerson) Type() int {
	return int((p.manaHealthFlags >> 23) & 0xFF) // Тип игрока в битах с 23 по 30
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
	assert.True(t, person.HasFamily())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
