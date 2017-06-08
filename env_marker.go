package kmsdecrypt

import (
	"fmt"
	"os"
	"strings"
)

// DecryptEnv will return a map[string]string decrypted Key-Value pairs from ENV variables that includes the supplied Marker m in it's Key.
// If m is an empty string it will default to "KMS_DECRYPT". The returned maps Key will have the marker removed from it's Key-name.
// Returns map[string]string and error.
func (d *KmsDecrypter) DecryptEnv(m string) (map[string]string, error) {
	if m == "" {
		m = "KMS_DECRYPT"
	}

	envs := os.Environ()
	resultChannel := make(chan resultKeyValue)
	count := 0

	for _, env := range envs {
		slice := strings.SplitN(env, "=", 2)
		if len(slice) != 2 {
			continue
		}

		newkey := ""
		key := slice[0]
		value := slice[1]

		// If marker wasn't found, continue to next env var.
		if !strings.Contains(key, m) {
			continue
		}

		count++

		// Create the new Key name.
		surrounding := fmt.Sprintf("_%v_", m)
		leading := fmt.Sprintf("_%v", m)
		trailing := fmt.Sprintf("%v_", m)

		switch {
		case strings.Contains(key, surrounding):
			newkey = strings.Replace(key, surrounding, "", 1)

		case strings.Contains(key, leading):
			newkey = strings.Replace(key, leading, "", 1)

		case strings.Contains(key, trailing):
			newkey = strings.Replace(key, trailing, "", 1)

		default:
			newkey = strings.Replace(key, m, "", 1)
		}

		go d.decryptKeyValue(&newkey, &value, resultChannel)

	}

	// Wait for all go-routines to finish.
	result := make(map[string]string)
	for i := 0; i < count; i++ {
		res := <-resultChannel

		if res.err != nil {
			return result, res.err
		}

		result[res.key] = res.val
	}

	return result, nil
}
