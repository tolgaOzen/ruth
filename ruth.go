package ruth


type Exception struct {
	Error interface{}
}

type Section struct {
	Exception
	innerTry     func()
}

func Try(innerTry func()) (section *Section) {

	section = &Section{innerTry: innerTry}

	defer func() {
		exception := Exception{}
		exception.Error = recover()
		section.Exception = exception
	}()

	section.innerTry()

	return
}

func (s *Section) Catch(innerCatch func(e Exception)) (section *Section) {

	if s.Exception.Error != nil {

		defer func() {

			if err := recover(); err != nil {
				s.Exception.Error = err
			} else {
				s.Exception.Error = nil
			}

			section = s
		}()

		innerCatch(s.Exception)
		return s
	}

	return
}


func (s *Section) TryAgain() (section *Section) {

	defer func() {
		if err := recover(); err != nil {
			s.Exception.Error = err
			section = s
		}
	}()

	s.innerTry()
	return
}


func (s *Section) Finally(innerFinally func()) {
	innerFinally()
}
