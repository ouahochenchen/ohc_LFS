package err_code

type MyError struct {
	Msg string
}

func (m *MyError) Error() string {
	return m.Msg
}
