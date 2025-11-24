package car

import "fmt"

type CarModule struct {
	characteristics CarCharacteristics
	history         map[int]CarHistoryInfo
}

type CarCharacteristics struct {
	Brand string
	Model string
	Year  int
	Color string
}

type CarHistoryInfo struct {
	Description string
	Distance    int
	Sport       bool
}

type DriveCar interface {
	Drive(description string, distance int) CarHistoryInfo
	ChangeColor(newColor string) CarCharacteristics
}

// NewCarModule creates a CarModule with given characteristics and initialized history.
func NewCarModule(cc CarCharacteristics) *CarModule {
	return &CarModule{
		characteristics: cc,
		history:         make(map[int]CarHistoryInfo),
	}
}

// Drive records a drive entry and references the characteristics field so it is used.
func (cm *CarModule) Drive(description string, distance int) CarHistoryInfo {
	// include brand and model in description to demonstrate usage of characteristics
	fullDesc := cm.characteristics.Brand + " " + cm.characteristics.Model + ": " + description
	entry := CarHistoryInfo{
		Description: fullDesc,
		Distance:    distance,
	}
	// store in history with a new incremental key
	key := len(cm.history) + 1
	cm.history[key] = entry
	return entry
}

func (cm *CarModule) ChangeColor(newColor string) CarCharacteristics {
	fmt.Println("cm:", cm)
	fmt.Printf("type: %T, address: %p\n", cm, cm)
	cm.characteristics.Color = newColor
	return cm.characteristics
}
