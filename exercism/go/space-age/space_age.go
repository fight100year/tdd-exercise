package space

// Planet is palnet for earth years
type Planet string

// plant's earth years
const (
	Earth   = 1.0
	Mercury = 0.2408467
	Venus   = 0.61519726
	Mars    = 1.8808158
	Jupiter = 11.862615
	Saturn  = 29.447498
	Uranus  = 84.016846
	Neptune = 164.79132

	EarthSeconds = 31557600
)

func getPlanet(planet Planet) float64 {
	switch planet {
	case "Earth":
		return Earth
	case "Mercury":
		return Mercury
	case "Venus":
		return Venus
	case "Mars":
		return Mars
	case "Jupiter":
		return Jupiter
	case "Saturn":
		return Saturn
	case "Uranus":
		return Uranus
	case "Neptune":
		return Neptune
	default:
		return 1.0
	}
}

// Age calc the year of planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / EarthSeconds / getPlanet(planet)
}
