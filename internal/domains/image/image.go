package image

// ImageRepoContract defines set of methods for any type who wants play role as image repository
type ImageRepoContract interface {
	GetImageByImgables(imageable ImageableContract) (*Image, error)
	GetImage(id int) (*Image, error)
}

// ImageableContract defines set of methods that every thing that wants to have images should obey these methods
type ImageableContract interface {
	ImageableType() string
	ImageableId() int
}
