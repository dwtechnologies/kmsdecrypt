package kmsdecrypt

import (
	"os"
	"strings"
)

// DecryptEnvAuto will decrypt values from ENV variables automatically. It will check that the ENV value is divisible by 4, otherwise the ENV vill be ignored.
// Please note that any KMS decryption error will be treated as the ENV was not encrypted. So it can potentially be dangeroud, so use with causion.
// Returns a map[string]string of decrypted Key-Value pairs.
func (d *KmsDecrypter) DecryptEnvAuto() map[string]string {
	envs := os.Environ()
	resultChannel := make(chan resultKeyValue)
	count := 0

	for _, env := range envs {
		slice := strings.SplitN(env, "=", 2)
		if len(slice) != 2 {
			continue
		}

		key := slice[0]
		value := slice[1]

		// If key isn't divisible with 4, ignore it as it's not base64 encoded.
		if (len(key) % 4) != 0 {
			continue
		}

		count++
		go d.decryptKeyValue(&key, &value, resultChannel)
	}

	// Wait for all go-routines to finish.
	result := make(map[string]string)
	for i := 0; i < count; i++ {
		res := <-resultChannel

		// This is what makes the DecryptEnvAuto potentially dangerous.
		if res.err != nil {
			continue
		}

		result[res.key] = res.val
	}
	return result
}
