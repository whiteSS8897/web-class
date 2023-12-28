package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Setting struct {
	BackEnd BackEndSetting `json:"Backend"`
}

//	type BackEndSetting struct {
//		BackendPort string `json:"BACKENDPORT"`
//		UserName    string `json:"USERNAME"`
//		PassWord    string `json:"PASSWORD"`
//		NetWorkType string `json:"NETWORKTYPE"`
//		DB          DatabaseConfig
//	}
type BackEndSetting struct {
	BackendPort  string `json:"BACKENDPORT"`
	UserName     string `json:"USERNAME"`
	PassWord     string `json:"PASSWORD"`
	NetWorkType  string `json:"NETWORKTYPE"`
	DBServerAddr string `json:"DBSERVERADDR"`
	DBPort       int    `json:"DBPORT"`
	DataBase     string `json:"DATABASE"`
}

// type DatabaseConfig struct {
// 	DBServerAddr string `json:"DBSERVERADDR"`
// 	DBPort       int    `json:"DBPORT"`
// 	DataBase     string `json:"DATABASE"`
// }

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token_btof struct {
	AccessToken string
	Signature   string
}

type Token_ftob struct {
	Token string `json:"token"`
}

var BACKENDPORt string  //default 8111
var USERNAMe string     //"user"
var PASSWORd string     //"passwd"
var NETWORKTYPe string  //"tcp/udp"
var DBSERVERADDr string //"0.0.0.0"
var DBPORt int          //0
var DATABASe string     //"users...etc"

var secretKey = []byte("12345") // 密鑰

func setTheConstDefault() {
	fmt.Println("[starting] start to set up const..")
	//讀取設定檔(json file)
	jsonFile, err := os.Open("../public/SettingConfig.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[starting] 成功開啟json設定檔，讀取中請稍後...")
	defer jsonFile.Close()

	byteValue, _ := os.ReadFile("../public/SettingConfig.json")
	var setting Setting
	json.Unmarshal(byteValue, &setting)

	BACKENDPORt = setting.BackEnd.BackendPort
	USERNAMe = setting.BackEnd.UserName
	PASSWORd = setting.BackEnd.PassWord
	NETWORKTYPe = setting.BackEnd.NetWorkType
	DBSERVERADDr = setting.BackEnd.DBServerAddr
	DBPORt = int(setting.BackEnd.DBPort)
	DATABASe = setting.BackEnd.DataBase
	fmt.Println("[starting] const set up done!")
}

var db *sql.DB // 資料庫連線全域變數

func CheckIdentity(c *gin.Context) {
	account := Account{}
	c.BindJSON(&account)
	fmt.Printf("%v", &account)
	c.JSON(http.StatusOK, gin.H{
		"username": account.Username,
		"password": account.Password,
	})
	fmt.Printf("usernam=%s   ,password=%s \n", account.Username, account.Password)

	var dbSalt, dbHash string
	err := db.QueryRow("SELECT user_psw_salt, user_psw_hash FROM users WHERE binary user_acc=?", account.Username).Scan(&dbSalt, &dbHash)
	if err != nil {
		fmt.Println("DB Query error")
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	inputHash := calculateSHA256Hash(account.Password, dbSalt)
	fmt.Println(inputHash)
	// 檢查帳號密碼是否有效
	isValid := inputHash == dbHash

	if isValid {
		fmt.Println("login success")
		token, err := generateAccessToken(account.Username, account.Password, db)
		if err != nil {
			fmt.Println("Failed to generate access token:", err)
			return
		}
		accessToken := token.AccessToken
		fmt.Println("Token:", token.Signature)

		payload, err := extractPayload(accessToken)
		if err != nil {
			fmt.Println("Failed to extract payload:", err)
			return
		}

		username, expiration, err := extractUsernameAndExpiration(payload)
		if err != nil {
			fmt.Println("Failed to extract username and expiration:", err)
			return
		}

		fmt.Println("Username:", username)
		fmt.Println("Expiration:", expiration)
		fmt.Println(token.Signature)

		c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token.Signature})

	} else {
		fmt.Println("wrong username or passwd")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password", "token": "error"})
	}

}

func calculateSHA256Hash(password, salt string) string {
	data := []byte(password + salt)
	hash := sha256.Sum256(data)
	hashedPassword := hex.EncodeToString(hash[:])
	return hashedPassword
}

// 產生 Token
func generateAccessToken(username, password string, db *sql.DB) (Token_btof, error) {
	expiration := time.Now().UTC().Add(time.Hour * 24) // Token有效時間

	// Token 建立
	payload := fmt.Sprintf(`{"username": "%s", "exp": "%s"}`, username, expiration.Format(time.RFC3339))
	signature := generateHMAC(payload, secretKey)
	accessToken := fmt.Sprintf("%s.%s", payload, signature)

	err := saveTokenToDatabase(username, signature, expiration, db)
	if err != nil {
		return Token_btof{}, err
	}

	return Token_btof{AccessToken: accessToken, Signature: signature}, nil
}

// 產生HMAC
func generateHMAC(data string, secretKey []byte) string {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(data))
	signature := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(signature)
}

// 提取payload
func extractPayload(accessToken string) (string, error) {
	tokenParts := strings.Split(accessToken, ".")
	if len(tokenParts) != 2 {
		return "", fmt.Errorf("Invalid access token format")
	}

	return tokenParts[0], nil
}

// 從payload提取帳號與時間
func extractUsernameAndExpiration(payload string) (string, string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		return "", "", fmt.Errorf("Failed to decode payload")
	}

	username, ok := data["username"].(string)
	if !ok {
		return "", "", fmt.Errorf("Failed to extract username")
	}

	expiration, ok := data["exp"].(string)
	if !ok {
		return "", "", fmt.Errorf("Failed to extract expiration")
	}

	return username, expiration, nil
}

func saveTokenToDatabase(username, signature string, expiration time.Time, db *sql.DB) error {
	// 檢查資料庫中是否已存在該使用者的 Token 記錄
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tokens WHERE Username = ?", username).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// 該使用者的 Token 記錄已存在則執行更新操作
		stmt, err := db.Prepare("UPDATE tokens SET Token = ?, Time = ? WHERE Username = ?")
		if err != nil {
			return err
		}
		defer stmt.Close()

		// 執行更新
		_, err = stmt.Exec(signature, expiration.Format("2006-01-02 15:04:05"), username)
		if err != nil {
			return err
		}
	} else {
		// 該使用者的 Token 記錄不存在則執行插入操作
		stmt, err := db.Prepare("INSERT INTO tokens (Username, Token, Time) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		// 執行插入
		_, err = stmt.Exec(username, signature, expiration.Format("2006-01-02 15:04:05"))
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckToken(c *gin.Context) {
	token := Token_ftob{}
	c.BindJSON(&token)
	fmt.Printf("%v", &token)
	c.JSON(http.StatusOK, gin.H{
		"token": token.Token,
	})
	fmt.Printf("token=%s", token.Token)       // 從前端的請求中獲取 token
	isValid := validateToken(token.Token, db) // 驗證 token 是否有效

	if isValid {
		fmt.Println("you have authority")
		username, err := getUsernameFromToken(token.Token) // 從資料庫獲取對應的使用者名稱
		if err != nil {
			fmt.Printf("error retrieving username: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		nickname, err := getNickname(username)
		if err != nil {
			fmt.Printf("error retrieving nickname: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// Token 验证成功，返回用户名
		c.JSON(http.StatusOK, gin.H{"message": "Token is valid", "username": username, "nickname": nickname})
		return
	} else {
		// Token 验证失败
		fmt.Println("you don't have authority", token.Token)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
}

func validateToken(token string, db *sql.DB) bool {
	// 執行查詢，檢查 token 是否存在於資料庫中
	var expiration string
	err := db.QueryRow("SELECT Time FROM tokens WHERE token = ?", token).Scan(&expiration)
	if err != nil {
		if err == sql.ErrNoRows {
			// Token 不存在於資料庫中
			fmt.Println("not exist")
			return false
		}
		// 處理資料庫查詢錯誤
		fmt.Println("Here")
		return false
	}

	// 檢查過期時間是否超過當前時間
	expirationTime, err := time.Parse("2006-01-02 15:04:05", expiration)
	if err != nil {
		// 處理時間解析錯誤
		fmt.Println("Error")
		return false
	}
	now := time.Now().UTC()
	if expirationTime.Before(now) {
		fmt.Println("access out time")
		return false // token 已過期
	}

	fmt.Println("access in time")
	return true // token 有效
}

func getUsernameFromToken(token string) (string, error) {
	// 資料庫連接
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAMe, PASSWORd, NETWORKTYPe, DBSERVERADDr, DBPORt, DATABASe))
	if err != nil {
		// 處理資料庫連接錯誤
		return "", err
	}
	defer db.Close()

	// 執行查詢，檢查 token 是否存在於資料庫中並獲取對應的使用者名稱
	var username string
	err = db.QueryRow("SELECT Username FROM tokens WHERE token = ?", token).Scan(&username)
	if err != nil {
		// 處理資料庫查詢錯誤
		return "", err
	}

	return username, nil
}

func getNickname(username string) (string, error) {
	// 在這裡執行查詢，獲取對應的 nickname
	var nickname string
	err := db.QueryRow("SELECT nickname FROM usersdata WHERE username=?", username).Scan(&nickname)
	if err != nil {
		return "", err
	}
	return nickname, nil
}

func UpdateNickname(c *gin.Context) {
	// 解析請求中的JSON數據，包括新的暱稱
	var request struct {
		Username    string `json:"userName"`
		NewNickname string `json:"newNickname"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("username:", request.Username)
	fmt.Println("暱稱:", request.NewNickname)

	err := saveNewNicknameToDatabase(request.Username, request.NewNickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功訊息
	c.JSON(http.StatusOK, gin.H{"message": "Nickname updated successfully"})
}

func saveNewNicknameToDatabase(username, newNickname string) error {
	query := "UPDATE usersdata SET nickname = ? WHERE username = ?"
	_, err := db.Exec(query, newNickname, username)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	setTheConstDefault()
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAMe, PASSWORd, NETWORKTYPe, DBSERVERADDr, DBPORt, DATABASe))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	// 配置 CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	// 添加路由和對應函數
	router.POST("/login/", CheckIdentity)
	router.POST("/checkToken/", CheckToken)
	router.POST("/updateNickname/", UpdateNickname)
	router.Run(":" + BACKENDPORt)

}
