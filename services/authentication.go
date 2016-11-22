package authentication

import(
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "time"
    "crypto/md5"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "bufio"
    "os"
    "fmt"
    "io"
    "encoding/hex"
    "../settings"
    "../models"
)

type (
    Claims struct {
        ID uint
        Name string
        Email string
        jwt.StandardClaims
    }
    AuthKeys struct {
        privateKey *rsa.PrivateKey
        PublicKey  *rsa.PublicKey
    }
)

var (
    Keys = AuthKeys{}
    LoggedUser = models.User{}
    //UserClaims = *Claims{}
)

func init(){
    Keys.privateKey = getPrivateKey()
    Keys.PublicKey = getPublicKey()
}

func respondWithError(code int, message string, c *gin.Context) {
    resp := map[string]string{"error": message}
    c.JSON(code, resp)
    c.Abort()
}

func IsAuthorized() gin.HandlerFunc{
    return func(c *gin.Context){
        token := c.Request.Header.Get("AuthToken")
        UserClaims, isValid := ValidateToken(token);
        if !isValid {
            respondWithError(401, "Unauthorized", c)
        } else if !models.CheckToken(token) {
            respondWithError(401, "Token expired", c)
        } else {
            LoggedUser = models.UserByToken(token, UserClaims.ID)
            //fmt.Println(LoggedUser)
            //fmt.Println(UserClaims)
            c.Next()
        }
    }
}

func getPrivateKey() *rsa.PrivateKey {
    privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
    if err != nil {
        panic(err)
    }
    defer privateKeyFile.Close()
    pemFileInfo, _ := privateKeyFile.Stat()
    var size int64 = pemFileInfo.Size()
    pemBytes := make([]byte, size)
    buffer := bufio.NewReader(privateKeyFile)
    _, err = buffer.Read(pemBytes)
    data, _ := pem.Decode([]byte(pemBytes))
    privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
    if err != nil {
        panic(err)
    }
    return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
    publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
    if err != nil {
        panic(err)
    }
    defer publicKeyFile.Close()
    pemFileInfo, _ := publicKeyFile.Stat()
    var size int64 = pemFileInfo.Size()
    pemBytes := make([]byte, size)
    buffer := bufio.NewReader(publicKeyFile)
    _, err = buffer.Read(pemBytes)
    data, _ := pem.Decode([]byte(pemBytes))
    publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
    if err != nil {
        panic(err)
    }
    rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
    if !ok {
        panic(err)
    }
    return rsaPub
}

func GetPassword(password string) string{
    h := md5.New()
    io.WriteString(h, password)
    io.WriteString(h, settings.Get().Salt)
    return hex.EncodeToString(h.Sum(nil))
}

func GenerateToken(user models.User) (jwtToken string){
    claims := Claims{
        user.ID,
        user.Name,
        user.Email,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Unix()+settings.Get().AuthTokenExpire,
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
    jwtToken, err := token.SignedString(Keys.privateKey)
    if err != nil{
        fmt.Println("Can't generate JWT token:", err)
    }
    return
}

func ValidateToken(jwtToken string) (*Claims, bool){
    if jwtToken == ""{
        return &Claims{}, false
    }
    token, _ := jwt.ParseWithClaims(jwtToken,&Claims{}, func(token *jwt.Token) (interface{}, error) {
        return Keys.PublicKey, nil
    })
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, true
    } else {
        return &Claims{}, false
    }
}

func GetTokenRemainValidility(jwtToken string) int64{
    if claims, valid := ValidateToken(jwtToken); valid ==true {
        return claims.ExpiresAt - time.Now().Unix();
    }
    return 0
}