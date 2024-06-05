package service

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/google/uuid"
)

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description"`
	Distance    int           `json:"distance" binding:"required"`
	Radius      float64       `json:"radius" binding:"required"`
	Mass        float64       `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type" binding:"required"`
}

var exoplanets = make(map[string]Exoplanet)
var mutex = &sync.Mutex{}

func AddExoplanet(e Exoplanet) Exoplanet {
	e.ID = uuid.New().String()
	mutex.Lock()
	exoplanets[e.ID] = e
	mutex.Unlock()
	return e
}

func ListExoplanets() []Exoplanet {
	mutex.Lock()
	defer mutex.Unlock()
	var exoplanetList []Exoplanet
	for _, exoplanet := range exoplanets {
		exoplanetList = append(exoplanetList, exoplanet)
	}
	return exoplanetList
}

func GetExoplanetByID(id string) (Exoplanet, error) {
	mutex.Lock()
	defer mutex.Unlock()
	exoplanet, exists := exoplanets[id]
	if !exists {
		return Exoplanet{}, errors.New("exoplanet not found")
	}
	return exoplanet, nil
}

func UpdateExoplanet(id string, updatedExoplanet Exoplanet) (Exoplanet, error) {
	mutex.Lock()
	defer mutex.Unlock()
	exoplanet, exists := exoplanets[id]
	if !exists {
		return Exoplanet{}, errors.New("exoplanet not found")
	}

	updatedExoplanet.ID = id
	if updatedExoplanet.Name == "" {
		updatedExoplanet.Name = exoplanet.Name
	}
	if updatedExoplanet.Description == "" {
		updatedExoplanet.Description = exoplanet.Description
	}
	if updatedExoplanet.Distance == 0 {
		updatedExoplanet.Distance = exoplanet.Distance
	}
	if updatedExoplanet.Radius == 0 {
		updatedExoplanet.Radius = exoplanet.Radius
	}
	if updatedExoplanet.Type == "" {
		updatedExoplanet.Type = exoplanet.Type
	}
	if updatedExoplanet.Type == Terrestrial && updatedExoplanet.Mass == 0 {
		updatedExoplanet.Mass = exoplanet.Mass
	}

	if err := ValidateExoplanet(updatedExoplanet); err != nil {
		return Exoplanet{}, err
	}

	exoplanets[id] = updatedExoplanet
	return updatedExoplanet, nil
}

func DeleteExoplanet(id string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := exoplanets[id]; !exists {
		return errors.New("exoplanet not found")
	}
	delete(exoplanets, id)
	return nil
}

func FuelEstimation(id string, crewCapacity int) (float64, error) {
	exoplanet, err := GetExoplanetByID(id)
	if err != nil {
		return 0, err
	}

	var gravity float64
	switch exoplanet.Type {
	case GasGiant:
		gravity = 0.5 / math.Pow(exoplanet.Radius, 2)
	case Terrestrial:
		gravity = exoplanet.Mass / math.Pow(exoplanet.Radius, 2)
	default:
		return 0, fmt.Errorf("invalid exoplanet type")
	}

	fuel := float64(exoplanet.Distance) / math.Pow(gravity, 2) * float64(crewCapacity)
	return fuel, nil
}

func ValidateExoplanet(e Exoplanet) error {
	if e.Distance < 10 || e.Distance > 1000 {
		return fmt.Errorf("distance must be between 10 and 1000 light years")
	}
	if e.Radius < 0.1 || e.Radius > 10 {
		return fmt.Errorf("radius must be between 0.1 and 10 Earth-radius units")
	}
	if e.Type == Terrestrial && (e.Mass < 0.1 || e.Mass > 10) {
		return fmt.Errorf("mass must be between 0.1 and 10 Earth-Mass units for terrestrial planets")
	}
	if e.Type != GasGiant && e.Type != Terrestrial {
		return fmt.Errorf("type must be either GasGiant or Terrestrial")
	}
	return nil
}
