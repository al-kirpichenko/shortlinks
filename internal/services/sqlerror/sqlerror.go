package sqlerror

func GetSQLState(err error) string {
	type checker interface {
		SQLState() string
	}
	pe := err.(checker)

	return pe.SQLState()
}
