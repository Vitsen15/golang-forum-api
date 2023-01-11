package security

import "golang.org/x/crypto/bcrypt"

// HashPassword using the bcrypt hashing algorithm
func HashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// DoPasswordsMatch Check if two passwords match using Bcrypt CompareHashAndPassword
// which return nil on success and an error on failure.
func DoPasswordsMatch(hashedPassword string, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
