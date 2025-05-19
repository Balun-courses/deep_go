package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func StringToByte(str string) []byte {
	if len(str) == 0 {
		return nil
	}

	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func ByteToString(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	return unsafe.String(unsafe.SliceData(data), len(data))
}

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > 42 {
			name = name[:42]
		}
		//len 41-46 bits
		person.rest[5] = byte(len(name)&0x3F)<<2 | person.rest[5]&0x01
		copy(person.name[:], StringToByte(name))
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
		// 0-9 bits - mana
		person.rest[0] = byte(mana & 0xFF)
		person.rest[1] = (person.rest[1] & 0xFC) | byte((mana>>8)&0x03)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		// 10‑19 bits - health
		person.rest[1] = (person.rest[1] & 0x03) | byte((health&0x3FF)<<2)
		person.rest[2] = (person.rest[2] & 0xF0) | byte((health>>6)&0x0F)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		// 20‑23 bits - respect
		person.rest[2] = (person.rest[2] & 0x0F) | byte((respect&0xF)<<4)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		// 24‑27 bits - strength
		person.rest[3] = (person.rest[3] & 0xF0) | byte(strength&0xF)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		// 28‑31 bits - experience
		person.rest[3] = (person.rest[3] & 0x0F) | byte((experience&0xF)<<4)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		// 32‑35 bits - level
		person.rest[4] = (person.rest[4] & 0xF0) | byte(level&0xF)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		// 38 - bit
		person.rest[4] |= 0x40
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		// 39 - bit
		person.rest[4] |= 0x80
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		// 40 - bit
		person.rest[5] |= 0x01
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		// 36‑37 bits
		person.rest[4] = (person.rest[4] & 0xCF) | byte((personType&3)<<4)
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
	rest    [6]byte
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}

	for _, option := range options {
		option(&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	//len 41-46 bits
	lenName := int(p.rest[5] >> 2)
	if lenName == 0 {
		return ""
	} else {
		return ByteToString(p.name[:lenName])
	}
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
	// 0-9 bits - mana
	return int(p.rest[0]) | int(p.rest[1]&0x03)<<8
}

func (p *GamePerson) Health() int {
	// 10‑19 bits - health
	return int(p.rest[1]>>2) | int(p.rest[2]&0x0F)<<6
}

func (p *GamePerson) Respect() int {
	// 20‑23 bits - respect
	return int(p.rest[2] >> 4)
}

func (p *GamePerson) Strength() int {
	// 24‑27 bits - strength
	return int(p.rest[3] & 0x0F)
}

func (p *GamePerson) Experience() int {
	// 28‑31 bits - experience
	return int(p.rest[3] >> 4)
}

func (p *GamePerson) Level() int {
	// 32‑35 bits - level
	return int(p.rest[4] & 0x0F)
}

func (p *GamePerson) HasHouse() bool {
	// 38 - bit
	return p.rest[4]&0x40 != 0

}

func (p *GamePerson) HasGun() bool {
	// 39 - bit
	return p.rest[4]&0x80 != 0
}

func (p *GamePerson) HasFamilty() bool {
	// 40 - bit
	return p.rest[5]&0x01 != 0
}

func (p *GamePerson) Type() int {
	// 36‑37 bits
	return int((p.rest[4] >> 4) & 0x03)
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
