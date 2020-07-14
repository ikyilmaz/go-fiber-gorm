package validators

type IValidator interface {
	Validate() (err error)
}
