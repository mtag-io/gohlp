package cloudinary

func New() *Class {
	this := Class{}
	this.credentials()
	return &this
}
