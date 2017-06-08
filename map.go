package kmsdecrypt

// DecryptMap will decrypt values from map[string]string m.
// Returns a map[string]string of decrypted Key-Value pairs.
func (d *KmsDecrypter) DecryptMap(m map[string]string) (map[string]string, error) {
	resultChannel := make(chan resultKeyValue)
	count := len(m)

	for key, value := range m {
		go d.decryptKeyValue(&key, &value, resultChannel)
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
