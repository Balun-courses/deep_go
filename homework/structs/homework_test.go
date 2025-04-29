package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const (
	// base mask
	manaBitsSize   = 10
	healthBitsSize = 10
	houseBitsSize  = 1
	gunBitsSize    = 1
	familyBitsSize = 1
	typeBitsSize   = 2

	// skill mask
	respectBitsSize    = 4
	strengthBitsSize   = 4
	experienceBitsSize = 4
	levelBitsSize      = 4
)

const (
	// base mask
	manaBitsOffset   = 0
	healthBitsOffset = 10
	houseBitsOffset  = 20
	gunBitsOffset    = 21
	familyBitsOffset = 22
	typeBitsOffset   = 23

	// skill mask
	respectBitsOffset    = 0
	strengthBitsOffset   = 4
	experienceBitsOffset = 8
	levelBitsOffset      = 12
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.nickname = *(*[42]byte)(unsafe.Pointer(unsafe.StringData(name)))
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
		person.golds = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.baseMask = setBits[uint32](person.baseMask, uint32(mana), manaBitsSize, manaBitsOffset)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.baseMask = setBits[uint32](person.baseMask, uint32(health), healthBitsSize, healthBitsOffset)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.skillMask = setBits[uint16](person.skillMask, uint16(respect), respectBitsSize, respectBitsOffset)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.skillMask = setBits[uint16](person.skillMask, uint16(strength), strengthBitsSize, strengthBitsOffset)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.skillMask = setBits[uint16](person.skillMask, uint16(experience), experienceBitsSize, experienceBitsOffset)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.skillMask = setBits[uint16](person.skillMask, uint16(level), levelBitsSize, levelBitsOffset)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.baseMask = setBits[uint32](person.baseMask, uint32(1), houseBitsSize, houseBitsOffset)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.baseMask = setBits[uint32](person.baseMask, uint32(1), gunBitsSize, gunBitsOffset)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.baseMask = setBits[uint32](person.baseMask, uint32(1), familyBitsSize, familyBitsOffset)
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.baseMask = setBits[uint32](person.baseMask, uint32(personType), typeBitsSize, typeBitsOffset)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	x         int32    // 4
	y         int32    // 4
	z         int32    // 4
	golds     uint32   // 4
	nickname  [42]byte // 42
	skillMask uint16   // 2 | 4 bits - lvl,  4 bits - exp, 4 bits - strength, 4 bits - respect
	baseMask  uint32   // 4 | 2 bits - type, 1 bit - family, 1 bit - gun, 1 bit - home, 10 bits - hp, 10 bits - mana
}

func getBits[T ~uint16 | ~uint32](x T, limit, offset uint8) T {
	return (x >> offset) & ((1 << limit) - 1)
}

func setBits[T ~uint16 | ~uint32](x, set T, limit, offset uint8) T {
	mask := (T(1)<<limit - 1) << offset
	// 1. очищаем место куда зотим вставить
	// 2. обрезаем нужную нам участок битов
	return (x & ^mask) | ((set & (T(1)<<limit - 1)) << offset)
}

func NewGamePerson(options ...Option) GamePerson {
	gp := GamePerson{}
	for _, option := range options {
		option(&gp)
	}
	return gp
}

func (p *GamePerson) Name() string {
	return unsafe.String(&p.nickname[0], len(p.nickname))
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
	return int(p.golds)
}

func (p *GamePerson) Mana() int {
	return int(getBits[uint32](p.baseMask, manaBitsSize, manaBitsOffset))
}

func (p *GamePerson) Health() int {
	return int(getBits[uint32](p.baseMask, healthBitsSize, healthBitsOffset))
}

func (p *GamePerson) Respect() int {
	return int(getBits[uint16](p.skillMask, respectBitsSize, respectBitsOffset))
}

func (p *GamePerson) Strength() int {
	return int(getBits[uint16](p.skillMask, strengthBitsSize, strengthBitsOffset))
}

func (p *GamePerson) Experience() int {
	return int(getBits[uint16](p.skillMask, experienceBitsSize, experienceBitsOffset))

}

func (p *GamePerson) Level() int {
	return int(getBits[uint16](p.skillMask, levelBitsSize, levelBitsOffset))
}

func (p *GamePerson) HasHouse() bool {
	return getBits[uint32](p.baseMask, houseBitsSize, houseBitsOffset) == 1

}

func (p *GamePerson) HasGun() bool {
	return getBits[uint32](p.baseMask, gunBitsSize, gunBitsOffset) == 1
}

func (p *GamePerson) HasFamilty() bool {
	return getBits[uint32](p.baseMask, familyBitsSize, familyBitsOffset) == 1
}

func (p *GamePerson) Type() int {
	return int(getBits[uint32](p.baseMask, typeBitsSize, typeBitsOffset))
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
