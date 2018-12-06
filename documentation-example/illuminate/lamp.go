// Package illuminate shines the light
package illuminate

// Lamp can be turned on or off using the On() and Off() functions
type Lamp struct {
}

// On will make the lamp shine
func (l Lamp) On() string {
	return "shine"
}

// Off will make the lamp go dark
func (l Lamp) Off() string {
	return "dark"
}
