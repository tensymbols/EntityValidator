package AdValidator

import "errors"

func Validate(title string, text string) error {
	if len(title) >= 100 ||
		title == "" ||
		len(text) >= 500 ||
		text == "" {
		return errors.New("invalid ad")
	}
	return nil
}
