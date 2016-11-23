package selecter

type Selecter interface {
	Select(*Population) *Population
}
