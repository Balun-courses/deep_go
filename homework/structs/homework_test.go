package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Option func(*GamePerson)

// Константы для типа персонажа
const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

// Bitmask для flags
const (
	hasHouseMask  = byte(1 << 6)
	hasGunMask    = byte(1 << 5)
	hasFamilyMask = byte(1 << 4)
	typeBitsMask  = byte(0x03)
)

// GamePerson — компактная структура с максимальным размером 64 байта
type GamePerson struct {
	x                  int32
	y                  int32
	z                  int32
	gold               uint32
	name               [42]byte
	hpAndStr           [2]byte
	mpAndRep           [2]byte
	lvlAndExp          byte
	houseGunFamilyType byte
}

func setName(p *GamePerson, name string) {
	for i := range p.name {
		if i < len(name) {
			p.name[i] = name[i]
		} else {
			p.name[i] = 0
		}
	}
}

func WithName(name string) Option {
	return func(p *GamePerson) {
		setName(p, name)
	}
}

func WithCoordinates(x, y, z int) Option {
	return func(p *GamePerson) {
		p.x = int32(x)
		p.y = int32(y)
		p.z = int32(z)
	}
}

func WithGold(gold int) Option {
	return func(p *GamePerson) {
		p.gold = uint32(gold)
	}
}

func WithMana(mana int) Option {
	return func(p *GamePerson) {
		p.mpAndRep[0] = byte(mana)
		p.mpAndRep[1] = byte(mana >> 8)
	}
}

func WithHealth(health int) Option {
	return func(p *GamePerson) {
		p.hpAndStr[0] = byte(health)
		p.hpAndStr[1] = byte(health >> 8)
	}
}

func WithRespect(respect int) Option {
	return func(p *GamePerson) {
		p.mpAndRep[1] = byte(respect<<4) | p.mpAndRep[1]
	}
}

func WithStrength(strength int) Option {
	return func(p *GamePerson) {
		p.hpAndStr[1] = byte(strength<<4) | p.hpAndStr[1]
	}
}

func WithExperience(experience int) Option {
	return func(p *GamePerson) {
		p.lvlAndExp = (p.lvlAndExp & 0xF0) | byte(experience&0x0F)
	}
}

func WithLevel(level int) Option {
	return func(p *GamePerson) {
		p.lvlAndExp = (p.lvlAndExp & 0x0F) | byte((level&0x0F)<<4)
	}
}

func WithHouse() Option {
	return func(p *GamePerson) {
		p.houseGunFamilyType |= hasHouseMask
	}
}

func WithGun() Option {
	return func(p *GamePerson) {
		p.houseGunFamilyType |= hasGunMask
	}
}

func WithFamily() Option {
	return func(p *GamePerson) {
		p.houseGunFamilyType |= hasFamilyMask
	}
}

func WithType(frac int) Option {
	return func(p *GamePerson) {
		p.houseGunFamilyType = (p.houseGunFamilyType & ^typeBitsMask) | byte(frac&int(typeBitsMask))
	}
}

func NewGamePerson(options ...Option) GamePerson {
	var p GamePerson
	for _, opt := range options {
		opt(&p)
	}
	return p
}

func (p *GamePerson) Name() string {
	for i, b := range p.name {
		if b == 0 {
			return string(p.name[:i])
		}
	}
	return string(p.name[:])
}

// coords
func (p *GamePerson) X() int { return int(p.x) }
func (p *GamePerson) Y() int { return int(p.y) }
func (p *GamePerson) Z() int { return int(p.z) }

// HP && MP
func (p *GamePerson) Mana() int {
	return int(p.mpAndRep[0]) + int(p.mpAndRep[1]<<6)>>6*256
}

func (p *GamePerson) Health() int {
	return int(p.hpAndStr[0]) + int(p.hpAndStr[1]<<6)>>6*256
}

// Str && Rep
func (p *GamePerson) Respect() int {
	return int(p.mpAndRep[1] >> 4)
}

func (p *GamePerson) Strength() int {
	return int(p.hpAndStr[1] >> 4)
}

// lvl && exp
func (p *GamePerson) Experience() int {
	return int(p.lvlAndExp & 0x0F)
}

func (p *GamePerson) Level() int {
	return int(p.lvlAndExp >> 4)
}

// gold
func (p *GamePerson) Gold() int { return int(p.gold) }

// flags
func (p *GamePerson) HasHouse() bool {
	return p.houseGunFamilyType&hasHouseMask != 0
}

func (p *GamePerson) HasGun() bool {
	return p.houseGunFamilyType&hasGunMask != 0
}

func (p *GamePerson) HasFamilty() bool {
	return p.houseGunFamilyType&hasFamilyMask != 0
}

func (p *GamePerson) Type() int {
	return int(p.houseGunFamilyType & typeBitsMask)
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
