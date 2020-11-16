/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package apiserver

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/check"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
	"github.com/Tencent/bk-bcs/bcs-common/common/encrypt"
	"github.com/Tencent/bk-bcs/bcs-common/common/http/httpserver"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/app/options"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/actions"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/drivers"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/drivers/mongo"
	storageErr "github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/errors"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/store"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-storage/storage/store/zookeeper"
)

const (
	configKeySep     = "/"
	mongodbConfigKey = "mongodb"
	zkConfigKey      = "zk"
)

// APIResource api resource object
type APIResource struct {
	Conf      *options.StorageOptions
	ActionsV1 []*httpserver.Action
	storeMap  map[string]store.Store
	dbMap     map[string]drivers.DB
}

var api = APIResource{}

// GetAPIResource that loads config from config-file and handlers from api-actions
func GetAPIResource() *APIResource {
	return &api
}

// SetConfig Set storageConfig to APIResource
func (a *APIResource) SetConfig(op *options.StorageOptions) {
	a.Conf = op
	// parse config-map from file
	dbConfig := a.ParseDBConfig()
	blog.Infof("Begin to parse databases.")

	a.storeMap = make(map[string]store.Store)
	a.dbMap = make(map[string]drivers.DB)
	for _, key := range dbConfig.KeyList {
		if _, ok := a.storeMap[key]; ok {
			blog.Warnf("Store config duplicated: %s", key)
			continue
		}
		if _, ok := a.dbMap[key]; ok {
			blog.Warnf("Database config duplicated: %s", key)
			continue
		}
		s := strings.Split(key, configKeySep)
		if len(s) != 2 {
			blog.Errorf("Database config invalid: %s | Format like mongodb/dynamic.", key)
			continue
		}
		var err error
		switch s[0] {
		case mongodbConfigKey:
			if err = a.parseMongodb(key, dbConfig); err != nil {
				SetUnhealthy(mongodbConfigKey, err.Error())
			}
		case zkConfigKey:
			if err = a.parseZk(key, dbConfig); err != nil {
				SetUnhealthy(zkConfigKey, err.Error())
			}
		default:
			err = storageErr.DatabaseConfigUnknown
			blog.Errorf("%v: %s", err, key)
			SetUnhealthy("unknown_config", fmt.Sprintf("%v: %s", err, key))
		}
		if err != nil {
			check.Occur(err)
		}
	}
	blog.Infof("Databases parsing completed.")
}

// InitActions init actions
func (a *APIResource) InitActions() {
	a.ActionsV1 = append(a.ActionsV1, actions.GetApiV1Action()...)
}

// ParseDBConfig parse db config
func (a *APIResource) ParseDBConfig() (dbConf *conf.Config) {
	dbConf = new(conf.Config)
	if _, err := os.Stat(a.Conf.DBConfig); !os.IsNotExist(err) {
		blog.Infof("Parsing config file: %s", a.Conf.DBConfig)
		dbConf.InitConfig(a.Conf.DBConfig)
	} else {
		blog.Errorf("Config file not exists: %s", a.Conf.DBConfig)
	}
	return
}

// GetDBClient get db client by key
func (a *APIResource) GetDBClient(key string) drivers.DB {
	return a.dbMap[key]
}

// GetStoreClient get store client by keys
func (a *APIResource) GetStoreClient(key string) store.Store {
	return a.storeMap[key]
}

func (a *APIResource) parseMongodb(key string, dbConf *conf.Config) error {
	address := dbConf.Read(key, "Addr")
	timeoutRaw := dbConf.Read(key, "ConnectTimeout")
	timeout, err := strconv.Atoi(timeoutRaw)
	if err != nil {
		return err
	}
	database := dbConf.Read(key, "Database")
	username := dbConf.Read(key, "Username")
	password := dbConf.Read(key, "Password")
	maxPoolSizeRaw := dbConf.Read(key, "MaxPoolSize")
	maxPoolSize, err := strconv.Atoi(maxPoolSizeRaw)
	if err != nil {
		return err
	}
	minPoolSizeRaw := dbConf.Read(key, "MinPoolSize")
	minPoolSize, err := strconv.Atoi(minPoolSizeRaw)
	if err != nil {
		return err
	}

	if password != "" {
		realPwd, _ := encrypt.DesDecryptFromBase([]byte(password))
		password = string(realPwd)
	}

	mongoOptions := &mongo.Options{
		Hosts:                 strings.Split(address, ","),
		ConnectTimeoutSeconds: timeout,
		Database:              database,
		Username:              username,
		Password:              password,
		MaxPoolSize:           uint64(maxPoolSize),
		MinPoolSize:           uint64(minPoolSize),
	}

	mongoDB, err := mongo.NewDB(mongoOptions)
	if err != nil {
		blog.Errorf("create mongo db with %s failed, err %s", key, err.Error())
		return fmt.Errorf("create mongo db with %s failed, err %s", key, err.Error())
	}
	err = mongoDB.Ping()
	if err != nil {
		blog.Errorf("ping mongo db failed, err %s", err.Error())
		return fmt.Errorf("ping mongo db failed, err %s", err.Error())
	}
	a.dbMap[key] = mongoDB
	blog.Infof("init mongo db with key %s successfully", key)
	return nil
}

func (a *APIResource) parseZk(key string, dbConf *conf.Config) error {
	address := dbConf.Read(key, "Addr")
	timeoutRaw := dbConf.Read(key, "ConnectTimeout")
	timeout, _ := strconv.Atoi(timeoutRaw)
	database := dbConf.Read(key, "Database")
	username := dbConf.Read(key, "Username")
	password := dbConf.Read(key, "Password")

	if password != "" {
		realPwd, _ := encrypt.DesDecryptFromBase([]byte(password))
		password = string(realPwd)
	}

	zkOpt := &zookeeper.Options{
		Addrs:                 strings.Split(address, ","),
		ConnectTimeoutSeconds: timeout,
		Database:              database,
		Username:              username,
		Password:              password,
	}

	zkStore, err := zookeeper.NewStore(zkOpt)
	if err != nil {
		return err
	}
	a.storeMap[key] = zkStore

	blog.Infof("init zookeeper with key %s successfully", key)
	return nil
}
