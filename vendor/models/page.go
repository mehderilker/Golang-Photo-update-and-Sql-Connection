package models

type SimplePage struct {
	Title string
	Description string
	Category []Part
}

type Home struct {
	Page SimplePage
	FileName string
	Status bool

}
type Part struct {
	Category string
}

type PhotoPage struct {
	Page SimplePage
	PhotoList []Photo
}

type PhotoUpdate struct {
	Page SimplePage
	PhotoInfo Photo
}
type PhotoDelete struct {
	Page SimplePage
	PhotoInfo Photo
	Response Response
}

