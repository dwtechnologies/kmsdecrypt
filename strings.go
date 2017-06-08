package kmsdecrypt

// DecryptStringSlice will decrypt values from []string s.
// Returns a []string of decrypted strings.
func (d *KmsDecrypter) DecryptStringSlice(s []string) ([]string, error) {
	resultChannel := make(chan resultString)
	count := len(s)

	for _, str := range s {
		go d.decryptString(&str, resultChannel)
	}

	// Wait for all go-routines to finish.
	result := []string{}
	for i := 0; i < count; i++ {
		res := <-resultChannel

		if res.err != nil {
			return result, res.err
		}

		result = append(result, res.str)
	}

	return result, nil
}

// DecryptString will string s. Returns a decrypted string.
func (d *KmsDecrypter) DecryptString(s string) (string, error) {
	resultChannel := make(chan resultString)
	count := 1

	go d.decryptString(&s, resultChannel)

	// Wait for the channel to send it's result.
	result := ""
	for i := 0; i < count; i++ {
		res := <-resultChannel

		if res.err != nil {
			return result, res.err
		}

		result = res.str
	}

	return result, nil
}
