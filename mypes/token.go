package mypes

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token ...
type Token struct {
	Model
	UserID     uint      `json:"user_id"`
	TokenKey   string    `json:"key"`
	Revoked    bool      `gorm:"default:'false'" json:"revoked"`
	ExpireTime time.Time `json:"expire_time"`
}

// Tokens ...
type Tokens []*Tokens

// JWTClaim Represents a basic type for JWT Claims
type JWTClaim struct {
	Key ClaimsKey   `json:"key"`
	Val interface{} `json:"val"`
}

// JWTClaims ...
type JWTClaims []*JWTClaim

// ClaimsKey ...
type ClaimsKey string

const (
	//ClaimUsername ....
	ClaimUsername ClaimsKey = "unam"
	// ClaimRole ...
	ClaimRole ClaimsKey = "rol"
	// ClaimExpTime ...
	ClaimExpTime ClaimsKey = "exp"
)

// ToString ...
func (c ClaimsKey) ToString() string {
	return string(c)
}

// CreateStandardToken ...
func CreateStandardToken(sectretKey, uname string, userID uint, role []int) (token *Token, e error) {
	now := time.Now()
	exp := now.Local().Add(time.Hour * time.Duration(5))
	// exp := now.Local().AddDate(1, 1, 1)
	jwtClaims := &JWTClaims{
		{ClaimUsername, uname},
		{ClaimRole, role},
		{ClaimExpTime, exp},
	}

	tokenKey, e := jwtClaims.GenerateToken(sectretKey)
	if e != nil {
		return
	}

	token = &Token{
		UserID:     userID,
		TokenKey:   tokenKey,
		Revoked:    false,
		ExpireTime: exp,
	}
	return
}

// CreateLongTimeToken ... creates a token with 1 year expire time
func CreateLongTimeToken(sectretKey, uname string, userID uint, role []int) (token *Token, e error) {

	now := time.Now()
	exp := now.Local().AddDate(1, 1, 1)

	jwtClaims := &JWTClaims{
		{ClaimUsername, uname},
		{ClaimRole, role},
		{ClaimExpTime, exp},
	}

	tokenKey, e := jwtClaims.GenerateToken(sectretKey)
	if e != nil {
		return
	}

	token = &Token{
		UserID:     userID,
		TokenKey:   tokenKey,
		Revoked:    false,
		ExpireTime: exp,
	}
	return
}

// GenerateToken ...
func (j JWTClaims) GenerateToken(secretKey string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Setting Claims
	for _, claim := range j {
		claims[string(claim.Key)] = claim.Val
	}

	// Generate encoded token.
	tokenKey, err := token.SignedString([]byte(secretKey))

	return tokenKey, err
}
