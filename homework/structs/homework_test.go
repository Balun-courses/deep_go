package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	// только ASCI символы
	return func(person *GamePerson) {
		for i := range len(name) {
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
		person.gold = int32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.manaAndRespect[0] = byte(mana)
		person.manaAndRespect[1] = byte(mana >> 8)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthAndStrength[0] = byte(health)
		person.healthAndStrength[1] = byte(health >> 8)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.manaAndRespect[1] = byte(respect<<4) | person.manaAndRespect[1]
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthAndStrength[1] = byte(strength<<4) | person.healthAndStrength[1]
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceAndLevel = byte(experience) | person.experienceAndLevel
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceAndLevel = byte(level<<4) | person.experienceAndLevel
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags = person.flags | 1
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags = person.flags | 4
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags = person.flags | 2
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags = byte(personType<<3) | person.flags
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

// размер структуры 64 байта
type GamePerson struct {
	// Имя пользователя [0…42] символов латиницы
	name [42]byte

	// Магическая сила (мана) [0…1000] 10 бит
	// Уважение [0…10] значений 4 бита
	manaAndRespect [2]byte

	// Здоровье [0…1000] 10 бит
	// Сила [0…10] значений 4 бита
	healthAndStrength [2]byte

	// Опыт [0…10] значений 4 бита
	// Уровень [0…10] значений 4 бита
	experienceAndLevel byte

	// Есть ли у игрока дом [true/false] значения 1 бит
	// Есть ли у игрока семья [true/false] значения 1 бит
	// Есть ли у игрока оружие [true/false] значения 1 бит
	// Тип игрока [строитель/кузнец/воин] значения 2 бита
	flags byte

	// Координаты по оси X, Y, Z [-2_000_000_000…2_000_000_000] 4 байта
	x, y, z int32

	// Золото [0…2_000_000_000] значений
	gold int32
}

func NewGamePerson(options ...Option) GamePerson {
	// need to implement
	person := GamePerson{}
	for _, o := range options {
		o(&person)
	}
	return person
}

func (p *GamePerson) Name() string {
	var length int
	for i := range len(p.name) {
		if p.name[i] == 0 {
			break
		}
		length++
	}
	return string(p.name[:length])
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
	return int(p.manaAndRespect[0]) + int(p.manaAndRespect[1]<<6)>>6*256
}

func (p *GamePerson) Health() int {
	return int(p.healthAndStrength[0]) + int(p.healthAndStrength[1]<<6)>>6*256
}

func (p *GamePerson) Respect() int {
	return int(p.manaAndRespect[1] >> 4)
}

func (p *GamePerson) Strength() int {
	return int(p.healthAndStrength[1] >> 4)

}

func (p *GamePerson) Experience() int {
	return int((p.experienceAndLevel << 4) >> 4)

}

func (p *GamePerson) Level() int {
	return int(p.experienceAndLevel >> 4)
}

func (p *GamePerson) HasHouse() bool {
	return p.flags&1 > 0
}

func (p *GamePerson) HasGun() bool {
	return p.flags&4 > 0
}

func (p *GamePerson) HasFamilty() bool {
	return p.flags&2 > 0
}

func (p *GamePerson) Type() int {
	return int(p.flags) >> 3
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
