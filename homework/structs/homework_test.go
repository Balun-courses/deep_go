package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"unsafe"
)

// Вспомогательные функции для битовых операций
func getBit(b byte, pos uint) bool {
	return (b>>pos)&1 == 1
}

func setBit(b byte, pos uint, val bool) byte {
	if val {
		return b | (1 << pos)
	}
	return b & ^(1 << pos)
}

// Вспомогательные функции для 3-х байтов на ману и здоровье
func getMana(b [3]byte) int {
	res := 0
	res += int(b[0])
	res += int(b[1]%(1<<2)) * (1 << 8)
	return res
}

func setMana(b [3]byte, mana int) [3]byte {
	b[0] += byte(mana % (1 << 8))
	b[1] += byte(mana / (1 << 8))
	return b
}

func getHealth(b [3]byte) int {
	res := 0
	res += int(b[1] / (1 << 2))
	res += int(b[2]) * (1 << 6)
	return res
}

func setHealth(b [3]byte, health int) [3]byte {
	b[1] += byte((health % (1 << 6)) * (1 << 2))
	b[2] += byte(health / (1 << 6))
	return b
}

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		res := [42]byte{}
		for i := 0; i < min(len(name), 42); i++ {
			res[i] = name[i]
		}
		person.name = res
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
		person.manaAndHealth = setMana(person.manaAndHealth, mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.manaAndHealth = setHealth(person.manaAndHealth, health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respectAndStrength = uint8(respect)*(1<<4) + (person.respectAndStrength)%(1<<4)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respectAndStrength = uint8(strength) + (person.respectAndStrength)/(1<<4)*(1<<4)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceAndLevel = uint8(experience)*(1<<4) + (person.experienceAndLevel)%(1<<4)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceAndLevel = uint8(level) + (person.experienceAndLevel)/(1<<4)*(1<<4)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.houseGunFamilyType = setBit(person.houseGunFamilyType, 0, true)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.houseGunFamilyType = setBit(person.houseGunFamilyType, 1, true)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.houseGunFamilyType = setBit(person.houseGunFamilyType, 2, true)
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		switch personType {
		case BuilderGamePersonType: // 00
		case BlacksmithGamePersonType: // 01
			person.houseGunFamilyType = setBit(person.houseGunFamilyType, 5, true)
		case WarriorGamePersonType: // 10
			person.houseGunFamilyType = setBit(person.houseGunFamilyType, 4, true)
		}
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	// Выравнивание в структуре по 4 байта

	name               [42]byte // 42 байта
	experienceAndLevel uint8    // 1 байт (0-3 бит - опыт, 4-7 - уровень)
	respectAndStrength uint8    // 1 байт (0-3 бит - уважение, 4-7 - сила)

	x    int32  // 4 байта
	y    int32  // 4 байта
	z    int32  // 4 байта
	gold uint32 // 4 байта

	manaAndHealth      [3]byte // 3 байта (0-9 бит - мана, 10-19 - здоровье)
	houseGunFamilyType byte    // 1 байт (1 бит - дом, 2 - оружие, 3 - семья, 4 и 5 - тип)
}

func NewGamePerson(options ...Option) GamePerson {
	gp := GamePerson{}

	for _, option := range options {
		option(&gp)
	}

	return gp
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
	return getMana(p.manaAndHealth)
}

func (p *GamePerson) Health() int {
	return getHealth(p.manaAndHealth)
}

func (p *GamePerson) Respect() int {
	return int(p.respectAndStrength / (1 << 4))
}

func (p *GamePerson) Strength() int {
	return int(p.respectAndStrength % (1 << 4))
}

func (p *GamePerson) Experience() int {
	return int(p.experienceAndLevel / (1 << 4))
}

func (p *GamePerson) Level() int {
	return int(p.experienceAndLevel % (1 << 4))
}

func (p *GamePerson) HasHouse() bool {
	return getBit(p.houseGunFamilyType, 0)
}

func (p *GamePerson) HasGun() bool {
	return getBit(p.houseGunFamilyType, 1)
}

func (p *GamePerson) HasFamily() bool {
	return getBit(p.houseGunFamilyType, 2)
}

func (p *GamePerson) Type() int {
	if getBit(p.houseGunFamilyType, 4) {
		return WarriorGamePersonType
	} else if getBit(p.houseGunFamilyType, 5) {
		return BlacksmithGamePersonType
	}
	return BuilderGamePersonType
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
