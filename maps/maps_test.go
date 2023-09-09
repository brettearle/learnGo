package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test value"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test value"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "added"

		err := dictionary.Add(key, value)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word errors", func(t *testing.T) {
		key := "test"
		value := "exists"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)
		assertError(t, err, ErrEntryExistsAlready)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	value := "to be deleted"
	dictionary := Dictionary{key: value}
	dictionary.Delete(key)
	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", key)
	}
}

func TestUpdate(t *testing.T) {

	t.Run("entry exists", func(t *testing.T) {
		key := "test"
		value := "new val"
		dictionary := Dictionary{key: "old val"}
		err := dictionary.Update(key, value)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("errors if no entry exists", func(t *testing.T) {
		key := "test"
		value := "non existent"
		dictionary := Dictionary{}
		err := dictionary.Update(key, value)
		assertError(t, err, ErrNoEntry)
	})
}

func assertDefinition(t testing.TB, d Dictionary, key string, want string) {
	t.Helper()
	got, err := d.Search(key)
	if err != nil {
		t.Fatal("should find added 'test' field containing 'added'")
	}
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got err %q want %q", got, want)
	}
}
