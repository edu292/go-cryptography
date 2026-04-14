package main

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	for text := range textCh {
		key, ok := <-keyCh
		if !ok {
			return
		}

		result <- text ^ key

	}
}
