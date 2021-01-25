// Package erratum implements the error handler for the error handling exercise
package erratum

// Use handles specific errors when opening a resource
func Use(o ResourceOpener, input string) (err error) {
	// try to open the resource
	r, err := o()

	// keep trying to open the resource if we have a TransientError
	if _, ok := err.(TransientError); ok {
		return Use(o, input)
	}

	// if we encounter another error, just return it
	if err != nil {
		return err
	}

	// no error so far, r successfully opened, so defer the closing
	defer r.Close()

	// prepare to catch possible panic with recover
	defer func() {
		if p := recover(); p != nil {
			if frobErr, ok := p.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			if otherErr, ok := p.(error); ok {
				err = otherErr
			}
		}
	}()

	// now see if Frob works...
	r.Frob(input)

	return err
}
