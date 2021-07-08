package space

const (
	MercuryOrbitalPeriod = 0.2408467
	VenusOrbitalPeriod   = 0.61519726
	EarthOrbitalPeriod   = 1.0
	MarsOrbitalPeriod    = 1.8808158
	JupiterOrbitalPeriod = 11.862615
	SaturnOrbitalPeriod  = 29.447498
	UranusOrbitalPeriod  = 84.016846
	NeptuneOrbitalPeriod = 164.79132
)

type Planet string

func Age(ageInSeconds float64, planet Planet) float64 {
	var planetAge float64
	ageInEarthYears := ageInSeconds / 31557600

	switch planet {
	case "Mercury":
		planetAge = ageInEarthYears / MercuryOrbitalPeriod
	case "Venus":
		planetAge = ageInEarthYears / VenusOrbitalPeriod
	case "Earth":
		planetAge = ageInEarthYears
	case "Mars":
		planetAge = ageInEarthYears / MarsOrbitalPeriod
	case "Jupiter":
		planetAge = ageInEarthYears / JupiterOrbitalPeriod
	case "Saturn":
		planetAge = ageInEarthYears / SaturnOrbitalPeriod
	case "Uranus":
		planetAge = ageInEarthYears / UranusOrbitalPeriod
	case "Neptune":
		planetAge = ageInEarthYears / NeptuneOrbitalPeriod
	default:
		planetAge = ageInEarthYears
	}

	return planetAge
}
