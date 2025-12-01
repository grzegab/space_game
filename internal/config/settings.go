package config

import (
	"math"
	"time"
)

const (
	LaserSpawnOffset     = 20.0
	DyingAnimationAmount = 50 * time.Millisecond
	DyingAnimationFrames = 12
	StarsNumber          = 2000
	ExhaustSpawnOffset   = -50.0
	ShieldOffset         = 20.0
)

// Intro settings
const (
	FadingOutLength = 40
	FadingInLength  = 15
	IntroDuration   = 3 * time.Second
)

// Speed settings
const (
	BaseVelocity     = 0.6
	RotationSpeed    = math.Pi
	RotationSpeedMin = 0.03
	RotationSpeedMax = -0.02
	MaxAcceleration  = 5.0
	AccelerationStep = 0.7
	AlienLaserSpeed  = 500.0
	LaserSpeed       = 900.0
)

// Timer settings
const (
	ShieldDuration      = 5 * time.Second
	HyperspaceCooldown  = 10 * time.Second
	DriftTime           = 30 * time.Second
	BurstCoolDown       = 1 * time.Second
	ShootCoolDown       = 200 * time.Millisecond
	BaseMeteorSpawnTime = 3 * time.Second
)

// Player settings
const (
	MaxShotsPerBurst = 2
	NumberOfLives    = 3
	NumberOfShields  = 3
)
