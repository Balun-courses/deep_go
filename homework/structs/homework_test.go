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
		if len(name) > 42 {
			return
		}
		person.name = [42]byte{}

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
	pozGo := 22
	return func(person *GamePerson) {
		bytee := uint32(mana) << pozGo
		mask := ^(uint32(0x3FF) << pozGo)
		person.stats = person.stats&mask | bytee
	}
}

func WithHealth(health int) func(*GamePerson) {
	pozGo := 12
	return func(person *GamePerson) {
		bytee := uint32(health) << pozGo
		mask := ^(uint32(0x3FF) << pozGo)
		person.stats = person.stats&mask | bytee
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		bytee := uint16(respect) << 12
		person.skills = person.skills&0x0FFF | bytee
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		bytee := uint16(strength) << 8
		person.skills = person.skills&0xFF0F | bytee
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		bytee := uint16(experience) << 4
		person.skills = person.skills&0xFFF0 | bytee
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		bytee := uint16(level)
		person.skills = person.skills&0xFFF0 | bytee
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		mask := uint32(1) << 7
		person.stats |= mask
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		mask := uint32(1) << 6
		person.stats |= mask
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		mask := uint32(1) << 5
		person.stats |= mask
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		bytee := uint32(personType)
		person.stats = (person.stats & ^uint32(0xF)) | bytee
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
	name    [42]byte
	skills  uint16 // respect, strength, experience, level (по 4 bit)
	stats   uint32 // mana, health (по 10 bit), house, gun, family (выделено 4 bit по 1 bit), typePerson(выделено 4 bit)
}

func NewGamePerson(options ...Option) GamePerson {

	p := GamePerson{}
	for _, opt := range options {
		opt(&p)
	}
	return p
}

func (p *GamePerson) Name() string {
	i := 0
	for {
		if i == len(p.name) || p.name[i] == 0 {
			break
		}
		i++
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
	pozGo := 22
	mask := uint32(0x3FF) << pozGo
	return int((p.stats & mask) >> pozGo)
}

func (p *GamePerson) Health() int {
	pozGo := 12
	mask := uint32(0x3FF) << pozGo
	return int((p.stats & mask) >> pozGo)
}

func (p *GamePerson) Respect() int {
	return int((p.skills & 0xF000) >> 12)
}

func (p *GamePerson) Strength() int {
	return int((p.skills & 0x0F00) >> 8)
}

func (p *GamePerson) Experience() int {
	return int((p.skills & 0x00F0) >> 4)
}

func (p *GamePerson) Level() int {
	return int(p.skills & 0x000F)
}

func (p *GamePerson) HasHouse() bool {
	return (p.stats>>7)&1 != 0
}

func (p *GamePerson) HasGun() bool {
	return (p.stats>>6)&1 != 0
}

func (p *GamePerson) HasFamilty() bool {
	return (p.stats>>5)&1 != 0
}

func (p *GamePerson) Type() int {
	return int(p.stats & 0xF)
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
