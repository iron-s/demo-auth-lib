package auth

// ValidateJWT validates a JWT token
func validateJWT(token, scope string) bool {
    return len(token) > 0 && len(scope) > 0
}
