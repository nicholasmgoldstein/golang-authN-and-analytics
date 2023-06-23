package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Insert user into database
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate JWT
	token, err := generateJWT(user.Email)
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	// Return token as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	// Get user from database based on email
	var user models.User
	if err := db.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare stored hashed password with provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate and send JWT if credentials are valid
	token, err := generateJWT(user.Email)
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Invalidate JWT or delete token from client side
	// You can't actually "invalidate" a JWT, but you can set the token's expiration time
	// to a past date so that it is no longer valid. You could also delete the token from
	// the client side by deleting it from the cookie or local storage.

	// For example, to set the expiration time of the JWT to a past date:
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Get JWT from cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value

	// Parse JWT claims
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get user information from database based on JWT claims
	var user models.User
	if err := db.Where("email = ?", claims.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Return user information as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func generateJWT(email string) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
