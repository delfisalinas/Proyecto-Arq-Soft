package users

import (
	"backend/domain/users"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key") //Define una clave secreta para firmar los tokens JWT.

func CheckPasswordHash(password, hash string) bool { //Compara la contraseña con el hash almacenado en la base de datos.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil //Si no hay error, la contraseña es correcta.
}

func GenerateJWT(username, userType string) (string, error) { //Genera un token JWT con el nombre de usuario y el tipo de usuario.
	expirationTime := time.Now().Add(24 * time.Hour) //El token expira en 24 horas.
	claims := &users.Claims{                         //Define los claims del token.
		Username: username,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{ //Define los claims estándar del token.
			ExpiresAt: expirationTime.Unix(), //La fecha de expiración del token.
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Crea un token JWT firmado con el algoritmo HS256.
	return token.SignedString(jwtKey)                          //Firma el token con la clave secreta y lo convierte en una cadena.
}

func Login(request users.LoginRequest) users.LoginResponse { //Función para el inicio de sesión de un usuario, verificar la contraseña y generar un token JWT si las credenciales son correctas.

	//COMPLETAR, validar contra la base de datos

	/*func Login(db *gorm.DB, request users.LoginRequest) (users.LoginResponse, error) {
	    var user users.User

	    // Buscar el usuario en la base de datos
	    if err := db.Where("username = ?", request.Username).First(&user).Error; err != nil {
	        if errors.Is(err, gorm.ErrRecordNotFound) {
	            return users.LoginResponse{}, errors.New("usuario no encontrado")
	        }
	        return users.LoginResponse{}, err
	    }

	    // Verificar la contraseña
	    if !CheckPasswordHash(request.Password, user.Password) {
	        return users.LoginResponse{}, errors.New("contraseña incorrecta")
	    }

	    // Generar el token JWT
	    token, err := GenerateJWT(user.Username, user.UserType)
	    if err != nil {
	        return users.LoginResponse{}, err
	    }

	    return users.LoginResponse{
	        Token: token,
	    }, nil
	} */

	return users.LoginResponse{
		Token: "abcdef123456",
	}
}

// Register crea un nuevo usuario en la base de datos.
func Register(db *gorm.DB, request users.RegisterRequest) (users.RegisterResponse, error) {
	// Verificar si el usuario ya existe
	var existingUser users.User
	if err := db.Where("username = ? OR email = ?", request.Username, request.Email).First(&existingUser).Error; err == nil {
		return users.RegisterResponse{}, errors.New("username or email already exists")
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return users.RegisterResponse{}, err
	}

	// Crear el nuevo usuario
	user := users.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
		UserType: request.UserType,
	}

	// Guardar el usuario en la base de datos
	if err := db.Create(&user).Error; err != nil {
		return users.RegisterResponse{}, err
	}

	// Responder con los datos del usuario
	return users.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		UserType: user.UserType,
	}, nil
}
