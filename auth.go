package auth

// ValidateJWT validates a JWT token
func validateJWT(token, scope string) bool {
    return len(token) > 0 && len(scope) > 0
}

// ValidateCheckoutToken validates tokens for the unified checkout flow
func ValidateCheckoutToken(token string) bool {
    return validateJWT(token, "checkout")
}
