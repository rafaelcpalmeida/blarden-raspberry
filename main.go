package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "encoding/json"
    "errors"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

type Request struct {
    Message string `json:"message"`
}

type Payload struct {
    Key string `json:"key"`
    Timestamp int64 `json:"timestamp"`
}

func Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
    block, err := aes.NewCipher(key[:])
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    if len(ciphertext) < gcm.NonceSize() {
        return nil, errors.New("malformed cipher text")
    }

    return gcm.Open(nil,
        ciphertext[:gcm.NonceSize()],
        ciphertext[gcm.NonceSize():],
        nil,
    )
}

func api2apiValidationMiddleware() gin.HandlerFunc {
    api2apiToken := os.Getenv("API2API_TOKEN")
    aesKey := os.Getenv("AES_KEY")

    if api2apiToken == "" {
        log.Fatal("Please, set API2API_TOKEN environment variable")
    }

    if aesKey == "" {
        log.Fatal("Please, set AES_KEY environment variable")
    }

    cipherKeyString, _ := hex.DecodeString(aesKey)
    var cipherKey [32]byte
    copy(cipherKey[:], cipherKeyString)

    Request := Request{}

    return func(c *gin.Context) {
        err := c.BindJSON(&Request)

        if err != nil {
            log.Error("JSON was malformed")
        }

        encryptedString, _ := hex.DecodeString(Request.Message)
        plain, _ := Decrypt(encryptedString, &cipherKey)

        payload := Payload{}
        err = json.Unmarshal(plain, &payload)

        if err != nil {
            log.Error("JSON was malformed")
            c.AbortWithStatus(400)
            return
        }

        if payload.Key != api2apiToken {
            log.Error("Invalid key")
            c.AbortWithStatus(401)
            return
        }

        if payload.Timestamp < (time.Now().Unix() - 5) || payload.Timestamp > time.Now().Unix() {
            log.Error("Invalid timestamp")
            c.AbortWithStatus(401)
            return
        }

        c.Next()
    }
}

func main() {
    gin.SetMode(gin.ReleaseMode)

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        log.Info("I'm alive and healthy")
        c.String(200, "ALIVE")
    })

    r.POST("/open-door", api2apiValidationMiddleware(), func(c *gin.Context) {
        go OpenDoor()
        log.Info("I've opened the door")
        c.JSON(200, gin.H{
            "status": "open",
        })
    })

    r.POST("/garage-gate", api2apiValidationMiddleware(), func(c *gin.Context) {
        resp, err := http.Get("http://blarden-mcu/garage-gate/" + os.Getenv("INTERNAL_COMMUNICATION_KEY"))

        if err != nil {
            log.Fatalln(err)
        }

        _, err = ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalln(err)
        }

        log.Println(resp.Status)

        if resp.Status == "200 OK" {
            c.JSON(200, gin.H{
                "status": "open",
            })

            return
        }

        c.AbortWithStatus(500)
    })

    r.POST("/outside-gate", api2apiValidationMiddleware(), func(c *gin.Context) {
        resp, err := http.Get("http://blarden-mcu/outside-gate/" + os.Getenv("INTERNAL_COMMUNICATION_KEY"))

        if err != nil {
            log.Fatalln(err)
        }

        _, err = ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalln(err)
        }

        log.Println(resp.Status)

        if resp.Status == "200 OK" {
            c.JSON(200, gin.H{
                "status": "open",
            })

            return
        }

        c.AbortWithStatus(500)
    })

    r.Run()
}