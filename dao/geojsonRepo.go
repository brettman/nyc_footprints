package dao

// GeoJSONRepo - a put your db abstractions here
type GeoJSONRepo struct {
	dbctx FootprintsDao
}

// NewGeoJSONRepo - initilaze a new repo
func NewGeoJSONRepo() *GeoJSONRepo {
	fpdao := *NewDao()
	return &GeoJSONRepo{dbctx: fpdao}
}
