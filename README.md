# SBDB CAD

A Go Client for Nasa's Small Body Database Close-Approach API

## Installation

```bash
go get github.com/evancaplan/sbdb_cad
```

## Search Options
```go
type SmallBodyOptions struct {
    
    	// Default value: "now" for current date
    	// Exludes dates earlier than this date
	DateMin             string `json:"dateMin"`

	// Default value: "+60" to add 60 days to current date
	// Excludes data later than this date
	DateMax             string `json:"dateMax"`
	
	//Default Value: none; Default unit: au
	//Excludes data with an approach greater than value
	DistanceMin         string `json:"distanceMin"`
	
	// Default Value: .05; Default unit: au
	// Excludes data with an approach less than value
	DistanceMax         string `json:"distanceMax"`
	
	// Default Value: none
	// Excludes H values smaller than this
	HMin                string `json:"hMin"`
	
	// Default Value: none
	// Excludes H values greater than this
	HMax                string `json:"hMax"`
	
	// Default Value: none; Units: km/s
	// Exlucdes data with Velocity Infinity less than this value
	VelocityInfMax      string `json:"velocityInfMax"`
	
	// Default Value: none; Units: km/s
	// Exlucdes data with Velocity Infinity greater than this value
	VelocityInfMin      string `json:"velocityInfMin"`
	
	//Default Value: none; Units: km/s
	// Exlucdes data with Velocity Relative less than this value
	VelocityRelativeMax string `json:"velocityRelativeMax"`
	
	// Default Value: none; Units: km/s
	// Exlucdes data with Velocity Relative greater than this value
	VelocityRelativeMin string `json:"velocityRelativeMin"`
	
	// Default Value: none
	// Limits data to objects of specified orbit class
	// See https://ssd-api.jpl.nasa.gov/doc/cad.html#sbdb_class_table
	Class               string `json:"class"`
	
	// Default Value: false
	// Limit data to  potentially hazardous asteroids
	Pha                 bool   `json:"pha"`
	
	// Default Value: false
	// Limit data to near-earth asteroids
	Nea                 bool   `json:"nea"`
	
	// Default Value: false
	// Limit data to comets
	Comet               bool   `json:"comet"`
	
	// Default Value: false
	// Limit data to near-earth asteroids and comets
	NeaComet            bool   `json:"neaComet"`
	
	// Default Value: false
	// Limit data to near-earth objects
	Neo                 bool   `json:"neo"`
	
	// Default Value: none
	// Limits data to objects of the specific kind
	// a = asteroid, an = numbered-asteroids, au = unnumbered-asteroids
	// c = comets, cn = numbered-comets, cu = unnumbered-comets
	// n = numbered-objects, u = unnumbered-objects
	Kind                string `json:"kind"`
	
	// Default Value: none
	// Filters data for objects with a matching SPK-ID
	Spk                 string `json:"spk"`
	
	// Default Value: none
	// Filters data for objects with a matching designation
	Designation         string `json:"designation"`
	
	// Default Value: "Earth"
	// Limits data to close-approaches to the specified body
	// See https://ssd-api.jpl.nasa.gov/doc/cad.html#cad_body_table
	Body                string `json:"body"`
	
	// Default Value: "date"
	// Sorts data by any of the fields in this struct
	Sort                string `json:"sort"`
	
	// Default Value: none
	// Limits data to first N results
	Limit               string `json:"limit"`
	
	// Default Value: none
	// Include the full-format object name/designation
	FullName            bool   `json:"fullName"`
}
```

## Usage

```go
imports(
    sbdb_cad "github.com/evancaplab/sbdb_cad"
)

func FindSmallBodyCloseApproachData(sbo sbdb_cad.SmallBodyOptions) ([]sbdb.SBCAD, error) {
	sbCADService := sbdb_cad.NewSbCADService()
	bodies, err := sbCADService.FindSbCADBy(sbo)
	if err != nil {
    		return err, nil
	}
	return bodies
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
