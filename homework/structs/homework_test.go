package main

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const (
	maxNameLen           = 42
	maxCoord      int    = 2_000_000_000
	maxGold       uint32 = 2_000_000_000
	maxMana       uint32 = 1_000
	maxHealth     uint32 = 1_000
	maxRespect    uint8  = 10
	maxStrength   uint8  = 10
	maxExperience uint8  = 10
	maxLevel      uint8  = 10
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = name[0:maxNameLen]
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		if x > maxCoord {
			person.x = maxCoord
		} else if x < -maxCoord {
			person.x = -maxCoord
		} else {
			person.x = x
		}

		if y > maxCoord {
			person.y = maxCoord
		} else if y < -maxCoord {
			person.y = -maxCoord
		} else {
			person.y = y
		}

		if z > maxCoord {
			person.z = maxCoord
		} else if z < -maxCoord {
			person.z = -maxCoord
		} else {
			person.z = z
		}
	}
}

func WithGold(gold uint32) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = gold

		if gold > maxGold {
			person.gold = maxGold
		}
	}
}

func WithMana(mana uint32) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mana = mana

		if mana > maxMana {
			person.mana = maxMana
		}
	}
}

func WithHealth(health uint32) func(*GamePerson) {
	return func(person *GamePerson) {
		person.health = health

		if health > maxHealth {
			person.health = maxHealth
		}
	}
}

func WithRespect(respect uint8) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respect = respect

		if respect > maxRespect {
			person.respect = maxRespect
		}
	}
}

func WithStrength(strength uint8) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strength = strength

		if strength > maxStrength {
			person.strength = maxStrength
		}
	}
}

func WithExperience(experience uint8) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experience = experience

		if experience > maxExperience {
			person.experience = maxExperience
		}
	}
}

func WithLevel(level uint8) func(*GamePerson) {
	return func(person *GamePerson) {
		person.level = level

		if level > maxLevel {
			person.level = maxLevel
		}
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasHouse = true
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasGun = true
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hasFamily = true
	}
}

func WithType(personType uint8) func(*GamePerson) {
	return func(person *GamePerson) {
		person.personType = personType
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	name       string
	x, y, z    int
	gold       uint32
	mana       uint32
	health     uint32
	respect    uint8
	strength   uint8
	experience uint8
	level      uint8
	personType uint8
	hasHouse   bool
	hasGun     bool
	hasFamily  bool
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}

	for _, option := range options {
		option(&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	return p.name
}

func (p *GamePerson) X() int {
	return p.x
}

func (p *GamePerson) Y() int {
	return p.y
}

func (p *GamePerson) Z() int {
	return p.z
}

func (p *GamePerson) Gold() uint32 {
	return p.gold
}

func (p *GamePerson) Mana() uint32 {
	return p.mana
}

func (p *GamePerson) Health() uint32 {
	return p.health
}

func (p *GamePerson) Respect() uint8 {
	return p.respect
}

func (p *GamePerson) Strength() uint8 {
	return p.strength
}

func (p *GamePerson) Experience() uint8 {
	return p.experience
}

func (p *GamePerson) Level() uint8 {
	return p.level
}

func (p *GamePerson) HasHouse() bool {
	return p.hasHouse
}

func (p *GamePerson) HasGun() bool {
	return p.hasGun
}

func (p *GamePerson) HasFamily() bool {
	return p.hasFamily
}

func (p *GamePerson) Type() uint8 {
	return p.personType
}

func TestGamePerson(t *testing.T) {
	require.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

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
	assert.Equal(t, -maxCoord, person.X())
	assert.Equal(t, maxCoord, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, maxGold, person.Gold())
	assert.Equal(t, uint32(mana), person.Mana())
	assert.Equal(t, uint32(health), person.Health())
	assert.Equal(t, uint8(respect), person.Respect())
	assert.Equal(t, uint8(strength), person.Strength())
	assert.Equal(t, uint8(experience), person.Experience())
	assert.Equal(t, uint8(level), person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamily())
	assert.False(t, person.HasGun())
	assert.Equal(t, uint8(personType), person.Type())
}
