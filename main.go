/*
Copyright 2023 Christos Triantafyllidis <christos.triantafyllidis@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/carenaggio/hermes/ent"
	"github.com/carenaggio/hermes/ent/system"
	"github.com/carenaggio/libs/crypt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var privateKey crypt.PrivateKey
var config HermesConfig

type LoginMessage struct {
	SystemId  uuid.UUID `json:"system_id"`
	PublicKey []byte    `json:"public_key"`
	Timestamp int64     `json:"timestamp"`
}

func httpApiPublicKey(c *gin.Context) {
	var publicKey crypt.PublicKey
	publicKey.Init(privateKey.PublicKey())
	resp := make(map[string][]byte)
	resp["public_key"] = privateKey.PublicKey()
	c.JSON(http.StatusOK, resp)
}

func httpHealthCheck(c *gin.Context) {
	c.Writer.Write([]byte("OK"))
}

func httpApiLogin(c *gin.Context) {
	var loginPayload crypt.SignedMessage
	var login_msg LoginMessage
	var public_key crypt.PublicKey

	if err := c.BindJSON(&loginPayload); err != nil {
		return
	}

	err := json.Unmarshal(loginPayload.Message, &login_msg)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unable to parse login message"})
		return
	}

	public_key.Init(login_msg.PublicKey)

	if !public_key.Verify(loginPayload) {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Signature verification failed"})
	}

	dbClient, err := ent.Open(config.DataBase.Driver, config.DataBase.DSN)
	if err != nil {
		log.Fatalf("failed opening DB connection: %v", err)
	}
	defer dbClient.Close()

	system, err := dbClient.System.Query().Where(system.SystemIDEQ(login_msg.SystemId)).Only(c.Request.Context())
	if ent.IsNotFound(err) {
		system, err = dbClient.System.Create().SetApproved(false).SetSystemID(login_msg.SystemId).SetPublicKey(login_msg.PublicKey).SetLastLogin(0).Save(c.Request.Context())
	}
	if err != nil {
		log.Fatalf("DB query error: %v\n", err)
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Backend failure"})
		return
	}

	if system.LastLogin >= login_msg.Timestamp {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Stale login, try again"})
		return
	}

	if !system.Approved {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "System is not approved"})
		return
	}

	system.Update().SetLastLogin(time.Now().UnixMilli())
	c.JSON(http.StatusOK, gin.H{"login": "OK", "system": system.SystemID})
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	var err error

	configFlag := flag.String("config", getEnv("HERMES_CONFIG", ""), "The hermes configuration file")
	keyfileFlag := flag.String("keyfile", getEnv("HERMES_KEYFILE", ""), "Private key used for message decryption and signature")
	flag.Parse()

	configfile := *configFlag
	if configfile == "" {
		log.Fatal("Configuration file wasn't specified.")
	}
	config, err = load_config(configfile)
	if err != nil {
		panic(err)
	}

	keyfile := *keyfileFlag
	if keyfile == "" {
		if config.KeyFile != "" {
			keyfile = config.KeyFile
		} else {
			log.Fatal("No keyfile was specified.")
		}
	}

	privateKey.Init(keyfile)

	log.Println("Waiting for DB to be become available")
	for {
		db, err := sql.Open(config.DataBase.Driver, config.DataBase.DSN)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = db.Ping()
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
		continue

	}

	dbClient, err := ent.Open(config.DataBase.Driver, config.DataBase.DSN)
	if err != nil {
		panic(err)
	}

	defer dbClient.Close()

	dbClient.Schema.Create(context.Background())

	r := gin.Default()
	r.GET("/health-check", httpHealthCheck)
	r.GET("/public_key", httpApiPublicKey)
	r.POST("/login", httpApiLogin)
	r.Run()
}
