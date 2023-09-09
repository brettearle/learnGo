package maps

const (
	ErrNotFound           = DictionaryErr("entry not found")
	ErrEntryExistsAlready = DictionaryErr("entry already exists")
	ErrNoEntry            = DictionaryErr("entry doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrEntryExistsAlready
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case ErrNotFound:
		return ErrNoEntry
	default:
		return err
	}
	return nil
}
