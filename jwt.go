package jwt

import (
    "github.com/dgrijalva/jwt-go"
)

// Generate a JWT token for a user
func generateToken(user *User) (string, error) {
    // Create a new token object with the user ID as a claim
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": user.ID,
    })

    // Generate a signed token string
    tokenString, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// Validate a JWT token and extract the user ID
func validateToken(tokenString string) (int, error) {
    // Parse the token and extract the claims
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Make sure the token is signed with the correct algorithm
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        // Return the secret key used to sign the token
        return []byte("your-secret-key"), nil
    })

    if err != nil {
        return 0, err
    }

    // Extract the user ID from the claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if id, ok := claims["id"].(int); ok {
            return id, nil
        }
    }

    return 0, fmt.Errorf("invalid token")
}