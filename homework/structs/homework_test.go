package main

import (
	"encoding/binary"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const (
	maxNameLen = 42

	minCoordinate = -2_000_000_000
	maxCoordinate = 2_000_000_000

	maxGold   = 2_000_000_000
	maxMana   = 1000
	maxHealth = 1000

	maxRespect    = 10
	maxStrength   = 10
	maxExperience = 10
	maxLevel      = 10

	// битовые параметры для статов
	goldShift       = 0
	manaShift       = 31
	healthShift     = 41
	respectShift    = 51
	strengthShift   = 55
	experienceShift = 59
	levelShift      = 63

	goldMask       uint64 = (1 << 31) - 1
	manaMask       uint64 = (1 << 10) - 1
	healthMask     uint64 = (1 << 10) - 1
	respectMask    uint64 = (1 << 4) - 1
	strengthMask   uint64 = (1 << 4) - 1
	experienceMask uint64 = (1 << 4) - 1

	// битовые параметры для флагов
	levelFlagsShift = 0
	houseFlag       = 1 << 3
	gunFlag         = 1 << 4
	familyFlag      = 1 << 5
	typeShift       = 6
	typeMask        = 0b11 << typeShift
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > maxNameLen {
			return
		}

		for i := 0; i < len(name); i++ {
			if !isLatin(name[i]) {
				return
			}
		}

		person.name = [42]byte{}

		copy(person.name[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.setCoordinate(0, x)
		person.setCoordinate(1, y)
		person.setCoordinate(2, z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		if gold < 0 || gold > maxGold {
			return
		}

		person.setStats(goldMask, goldShift, uint64(gold))
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		if mana < 0 || mana > maxMana {
			return
		}

		person.setStats(manaMask, manaShift, uint64(mana))
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		if health < 0 || health > maxHealth {
			return
		}

		person.setStats(healthMask, healthShift, uint64(health))
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		if respect < 0 || respect > maxRespect {
			return
		}

		person.setStats(respectMask, respectShift, uint64(respect))
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		if strength < 0 || strength > maxStrength {
			return
		}

		person.setStats(strengthMask, strengthShift, uint64(strength))
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		if experience < 0 || experience > maxExperience {
			return
		}

		person.setStats(experienceMask, experienceShift, uint64(experience))
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		if level < 0 || level > maxLevel {
			return
		}

		person.stats = (person.stats &^ (uint64(1) << levelShift)) |
			(uint64(level&1) << levelShift)

		person.flags = (person.flags &^ 0b111) |
			uint8((level>>1)&0b111)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= houseFlag
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= gunFlag
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= familyFlag
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		if personType < BuilderGamePersonType ||
			personType > WarriorGamePersonType {
			return
		}

		person.flags = (person.flags &^ typeMask) |
			uint8(personType<<typeShift)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	stats uint64
	flags uint8

	name [42]byte

	coordinates [12]byte
}

func NewGamePerson(options ...Option) GamePerson {
	var person GamePerson

	for _, option := range options {
		option(&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	n := 0

	for n < len(p.name) && p.name[n] != 0 {
		n++
	}

	return string(p.name[:n])
}

func (p *GamePerson) X() int {
	return p.coordinate(0)
}

func (p *GamePerson) Y() int {
	return p.coordinate(1)
}

func (p *GamePerson) Z() int {
	return p.coordinate(2)
}

func (p *GamePerson) Gold() int {
	return int(p.getStats(goldMask, goldShift))
}

func (p *GamePerson) Mana() int {
	return int(p.getStats(manaMask, manaShift))
}

func (p *GamePerson) Health() int {
	return int(p.getStats(healthMask, healthShift))
}

func (p *GamePerson) Respect() int {
	return int(p.getStats(respectMask, respectShift))
}

func (p *GamePerson) Strength() int {
	return int(p.getStats(strengthMask, strengthShift))
}

func (p *GamePerson) Experience() int {
	return int(p.getStats(experienceMask, experienceShift))
}

func (p *GamePerson) Level() int {
	low := int((p.stats >> levelShift) & 1)
	high := int(p.flags & 0b111)

	return low | (high << 1)
}

func (p *GamePerson) HasHouse() bool {
	return p.flags&houseFlag != 0
}

func (p *GamePerson) HasGun() bool {
	return p.flags&gunFlag != 0
}

func (p *GamePerson) HasFamilty() bool {
	return p.flags&familyFlag != 0
}

func (p *GamePerson) Type() int {
	return int((p.flags & typeMask) >> typeShift)
}

func (p *GamePerson) setStats(mask uint64, shift uint, value uint64) {
	p.stats = (p.stats &^ (mask << shift)) |
		((value & mask) << shift)
}

func (p *GamePerson) getStats(mask uint64, shift uint) uint64 {
	return (p.stats >> shift) & mask
}

func (p *GamePerson) setCoordinate(index int, value int) {
	if value < minCoordinate || value > maxCoordinate {
		return
	}

	binary.LittleEndian.PutUint32(
		p.coordinates[index*4:index*4+4],
		uint32(int32(value)),
	)
}

func (p *GamePerson) coordinate(index int) int {
	return int(int32(binary.LittleEndian.Uint32(
		p.coordinates[index*4 : index*4+4],
	)))
}

func isLatin(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	const x, y, z = -2_000_000_000, 2_000_000_000, 0
	const name = "aaaaaaaaaaaaabbbbbbbbbbbbbcccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = 2_000_000_000
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
