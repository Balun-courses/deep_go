package main

import (
	"fmt"
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const (
	mask         uint32 = 0b1000000000_0000000000_0000000000_00
	manaMask     uint32 = 0b1111111111_0000000000_0000000000_00
	healthMask   uint32 = 0b0000000000_1111111111_0000000000_00
	strengthMask uint32 = 0b0000000000_0000000000_1111000000_00
	expMask      uint32 = 0b0000000000_0000000000_0000111100_00
	respectMask  uint32 = 0b0000000000_0000000000_0000000011_11

	lvlMask    uint16 = 0b1111000000_000000
	homeMask   uint16 = 0b0000100000_000000
	weaponMask uint16 = 0b0000010000_000000
	familyMask uint16 = 0b0000001000_000000
	typeMask   uint16 = 0b0000000110_000000
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		i := 0
		fmt.Println("LEN = ", len(name))
		for ; i < len(name); i++ {
			person.NameV[i] = name[i]
		}

		for ; i < 42; i++ {
			person.NameV[i] = 0
		}
		// need to implement
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
		person.XV = int32(x)
		person.YV = int32(y)
		person.ZV = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
		person.GoldV = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	if mana > 1000 || mana < 0 {
		return nil
	}
	return func(person *GamePerson) {
		// need to implement
		person.ManaHealthRespectStrengthExp &= ^manaMask
		person.ManaHealthRespectStrengthExp |= (uint32(mana) << 22)
	}
}

func WithHealth(health int) func(*GamePerson) {
	if health > 1000 || health < 0 {
		return nil
	}
	return func(person *GamePerson) {
		// need to implement
		person.ManaHealthRespectStrengthExp &= ^healthMask
		person.ManaHealthRespectStrengthExp |= (uint32(health) << 12)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	if respect > 10 || respect < 0 {
		return nil
	}

	return func(person *GamePerson) {
		// need to implement
		person.ManaHealthRespectStrengthExp &= ^respectMask
		person.ManaHealthRespectStrengthExp |= (uint32(respect))
	}
}

func WithStrength(strength int) func(*GamePerson) {
	if strength > 10 || strength < 0 {
		return nil
	}

	return func(person *GamePerson) {
		// need to implement
		person.ManaHealthRespectStrengthExp &= ^strengthMask
		person.ManaHealthRespectStrengthExp |= (uint32(strength) << 8)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	if experience > 10 || experience < 0 {
		return nil
	}

	return func(person *GamePerson) {
		// need to implement
		person.ManaHealthRespectStrengthExp &= ^expMask
		person.ManaHealthRespectStrengthExp |= (uint32(experience) << 4)
	}
}

func WithLevel(level int) func(*GamePerson) {
	if level > 10 || level < 0 {
		return nil
	}

	return func(person *GamePerson) {
		// need to implement
		person.LvlHomeWeaponFamilyType &= ^lvlMask
		person.LvlHomeWeaponFamilyType |= (uint16(level) << 12)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
		person.LvlHomeWeaponFamilyType &= ^homeMask
		person.LvlHomeWeaponFamilyType |= (uint16(1) << 11)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
		person.LvlHomeWeaponFamilyType &= ^weaponMask
		person.LvlHomeWeaponFamilyType |= (uint16(1) << 10)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		// need to implement
		person.LvlHomeWeaponFamilyType &= ^familyMask
		person.LvlHomeWeaponFamilyType |= (uint16(1) << 9)
	}
}

func WithType(personType int) func(*GamePerson) {
	if personType > 3 || personType < 0 {
		return nil
	}

	return func(person *GamePerson) {
		// need to implement
		person.LvlHomeWeaponFamilyType &= ^typeMask
		person.LvlHomeWeaponFamilyType |= (uint16(personType) << 7)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	NameV                        [42]byte // 4
	LvlHomeWeaponFamilyType      uint16   // 4 | 32 | 4, 1, 1, 1, 2
	XV                           int32    // 4
	YV                           int32    // 4
	ZV                           int32    // 4
	GoldV                        uint32   // 4
	ManaHealthRespectStrengthExp uint32   // 4 | 32 | 10, 10, 4, 4, 4, 4
	// need to implement
}

func NewGamePerson(options ...Option) GamePerson {
	// need to implement
	game := GamePerson{}
	for _, opt := range options {
		opt(&game)
	}
	return game
}

func (p *GamePerson) Name() string {
	// need to implement
	return string(p.NameV[:])
}

func (p *GamePerson) X() int {
	// need to implement
	return int(p.XV)
}

func (p *GamePerson) Y() int {
	// need to implement
	return int(p.YV)
}

func (p *GamePerson) Z() int {
	// need to implement
	return int(p.ZV)
}

func (p *GamePerson) Gold() int {
	// need to implement
	return int(p.GoldV)
}

func (p *GamePerson) Mana() int {
	// need to implement
	return int((p.ManaHealthRespectStrengthExp & manaMask) >> 22)
}

func (p *GamePerson) Health() int {
	// need to implement
	return int((p.ManaHealthRespectStrengthExp & healthMask) >> 12)
}

func (p *GamePerson) Respect() int {
	// need to implement
	return int((p.ManaHealthRespectStrengthExp & respectMask))
}

func (p *GamePerson) Strength() int {
	// need to implement
	return int((p.ManaHealthRespectStrengthExp & strengthMask) >> 8)
}

func (p *GamePerson) Experience() int {
	// need to implement
	return int((p.ManaHealthRespectStrengthExp & expMask) >> 4)
}

func (p *GamePerson) Level() int {
	// need to implement
	return int((p.LvlHomeWeaponFamilyType & lvlMask) >> 12)
}

func (p *GamePerson) HasHouse() bool {
	// need to implement
	return int((p.LvlHomeWeaponFamilyType&homeMask)>>11) > 0
}

func (p *GamePerson) HasGun() bool {
	// need to implement
	return int((p.LvlHomeWeaponFamilyType&weaponMask)>>10) > 0
}

func (p *GamePerson) HasFamilty() bool {
	// need to implement
	return ((p.LvlHomeWeaponFamilyType & familyMask) >> 9) > 0
}

func (p *GamePerson) Type() int {
	// need to implement
	return int((p.LvlHomeWeaponFamilyType & typeMask) >> 7)
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
