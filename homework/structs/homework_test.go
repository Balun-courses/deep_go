package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const (
	NameLength = 42
)

const (
	// attrs
	ManaOffset      = 22
	HealthOffset    = 12
	HasHouseOffset  = 11
	HasGunOffset    = 10
	HasFamilyOffset = 9
	TypeOffset      = 7

	// props
	RespectOffset    = 12
	StrengthOffset   = 8
	ExperienceOffset = 4
)

const (
	TenBitMask  = 0x3FF
	FourBitMask = 0xF
	TwoBitMask  = 0b11
	OneBitMask  = 0b1
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		for i := 0; i < NameLength; i++ {
			person.name[i] = name[i]
		}
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
		person.attrs |= uint32(mana) << ManaOffset
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.attrs |= uint32(health) << HealthOffset
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.props |= uint16(respect) << RespectOffset
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.props |= uint16(strength) << StrengthOffset
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.props |= uint16(experience) << ExperienceOffset
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.props |= uint16(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.attrs |= 1 << HasHouseOffset
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.attrs |= 1 << HasGunOffset
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.attrs |= 1 << HasFamilyOffset
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.attrs |= uint32(personType) << TypeOffset
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	x, y, z int32
	gold    uint32
	attrs   uint32
	props   uint16
	name    [NameLength]byte
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
	return int(p.attrs >> ManaOffset)
}

func (p *GamePerson) Health() int {
	return int((p.attrs >> HealthOffset) & TenBitMask)
}

func (p *GamePerson) Respect() int {
	return int(p.props >> RespectOffset)
}

func (p *GamePerson) Strength() int {
	return int((p.props >> StrengthOffset) & FourBitMask)
}

func (p *GamePerson) Experience() int {
	return int((p.props >> ExperienceOffset) & FourBitMask)
}

func (p *GamePerson) Level() int {
	return int(p.props & FourBitMask)
}

func (p *GamePerson) HasHouse() bool {
	return (p.attrs>>HasHouseOffset)&OneBitMask == 1
}

func (p *GamePerson) HasGun() bool {
	return (p.attrs>>HasGunOffset)&OneBitMask == 1
}

func (p *GamePerson) HasFamilty() bool {
	return (p.attrs>>HasFamilyOffset)&OneBitMask == 1
}

func (p *GamePerson) Type() int {
	return int((p.attrs >> TypeOffset) & TwoBitMask)
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
