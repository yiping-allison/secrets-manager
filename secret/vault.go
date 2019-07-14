package secret

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/yiping-allison/secrets-manager/cipher"
)

// File creates an instance of a Vault file, where all your secret keys and values
// will be stored, including an encodingKey which is used to encrypt and decrypt
// your file contents
func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

// Vault is a struct with the basic userinfo including:
// encodingKey,
// filepath to home directory,
// mutex to prevent multiple accesses by goroutines, and an
// actual map of key, values in user's secret file
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

// Get will retrieve the value of the given key provided by the user
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}
	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}

// Set will allow the user to set a key, value pair into their secret file.
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.save()
	return err
}

// List will return the total amount of keys (slice of string) stored in the user's
// secret file.
func (v *Vault) List() ([]string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	ret := make([]string, 0, len(v.keyValues))
	if err != nil {
		return nil, err
	}
	for k := range v.keyValues {
		ret = append(ret, k)
	}
	return ret, nil
}

// Remove will take in a key and delete it from the Vault map
func (v *Vault) Remove(key string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	delete(v.keyValues, key)
	err = v.save()
	return err
}
