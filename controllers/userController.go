package controllers

import (
    "database/sql"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/dhekerlaadhibi/LearnGo/jwt/database"
    "github.com/dhekerlaadhibi/LearnGo/jwt/models"
    "github.com/go-playground/validator/v10"
    "golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func VerifyPassword(userPassword string, providedPassword string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
    if err != nil {
        return false, err
    }
    return true, nil
}

func Signup() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur de liaison JSON: " + err.Error()})
            return
        }

        validationErr := validate.Struct(user)
        if validationErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur de validation des données: " + validationErr.Error()})
            return
        }

        // Hash du mot de passe
        hashedPassword, err := HashPassword(*user.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du hachage du mot de passe: " + err.Error()})
            return
        }

        // Connexion à la base de données
        db := database.GetDB()

        // Exécution de la requête SQL pour insérer un nouvel utilisateur
        result, err := db.Exec("INSERT INTO users (first_name, last_name, password, email, phone, user_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
            user.First_name, user.Last_name, hashedPassword, *user.Email, *user.Phone, *user.User_type, time.Now(), time.Now())
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'insertion de l'utilisateur dans la base de données: " + err.Error()})
            return
        }

        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Aucune ligne n'a été insérée dans la base de données"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Utilisateur créé avec succès"})
    }
}
func Login() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur de liaison JSON"})
            return
        }

        // Connexion à la base de données
        db := database.GetDB()

        // Recherche de l'utilisateur par email
        var storedPassword string
        err := db.QueryRow("SELECT password FROM users WHERE email = $1", user.Email).Scan(&storedPassword)
        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou mot de passe incorrect"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la recherche de l'utilisateur"})
            return
        }

        // Vérification du mot de passe
        passwordIsValid, err := VerifyPassword(storedPassword, *user.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la vérification du mot de passe"})
            return
        }

        if !passwordIsValid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou mot de passe incorrect"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Connexion réussie", "user": user})
    }
}



func GetUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Connexion à la base de données
        db := database.GetDB()

        // Exécution de la requête SQL pour récupérer tous les utilisateurs
        rows, err := db.Query("SELECT * FROM users")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des utilisateurs: " + err.Error()})
            return
        }
        defer rows.Close()

        var users []models.User
        for rows.Next() {
            var user models.User
            err := rows.Scan(&user.ID, &user.First_name, &user.Last_name, &user.Password, &user.Email, &user.Phone, &user.User_type, &user.Created_at, &user.Updated_at)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture des utilisateurs: " + err.Error()})
                return
            }
            users = append(users, user)
        }

        c.JSON(http.StatusOK, users)
    }
}


func GetUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        userId := c.Param("user_id")

        // Connexion à la base de données
        db := database.GetDB()

        // Exécution de la requête SQL pour récupérer l'utilisateur par ID
        var user models.User
        err := db.QueryRow("SELECT * FROM users WHERE id = $1", userId).Scan(
            &user.ID, &user.First_name, &user.Last_name, &user.Password, &user.Email, &user.Phone, &user.User_type, &user.Created_at, &user.Updated_at)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Utilisateur non trouvé: " + err.Error()})
            return
        }

        c.JSON(http.StatusOK, user)
    }
}
