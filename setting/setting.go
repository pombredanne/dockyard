/*
Copyright 2015 The ContainerOps Authors All rights reserved.

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

package setting

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/config"
	"github.com/fernet/fernet-go"
)

var (
	//@Global Config

	//AppName should be "dockyard"
	AppName string
	//Usage is short description
	Usage   string
	Version string
	Author  string
	Email   string

	//@Basic Runtime Config

	RunMode        string
	ListenMode     string
	HTTPSCertFile  string
	HTTPSKeyFile   string
	LogPath        string
	LogLevel       string
	DatabaseDriver string
	DatabaseURI    string
	Domains        string

	//@Docker V1 Config

	DockerStandalone      string
	DockerRegistryVersion string
	DockerV1Storage       string

	//@Docker V2 Config

	DockerDistributionVersion string
	DockerV2Storage           string

	//@Appc Config

	AppcStorage string

	//@UpdateService

	KeyManager string
	Storage    string

	//@ScanContent
	//32-bit URL-safe base64 key used to encrypt id in database

	ScanKey string
)

//
func init() {
	conf, err := getConfig()
	if err == nil {
		err = setGlobalConfig(conf)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getConfig() (conf config.Configer, err error) {
	home := os.Getenv("HOME")
	if home != "" {
		homePath := filepath.Join(home, ".dockyard", "containerops.conf")
		conf, err = config.NewConfig("ini", homePath)
	}

	if err != nil {
		conf, err = config.NewConfig("ini", "conf/containerops.conf")
	}

	return
}

func LoadServerConfig() {
	conf, err := getConfig()
	if err == nil {
		err = setServerConfig(conf)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setGlobalConfig(conf config.Configer) error {
	//config globals
	if appname := conf.String("appname"); appname != "" {
		AppName = appname
	} else if appname == "" {
		return fmt.Errorf("AppName config value is null")
	}

	if usage := conf.String("usage"); usage != "" {
		Usage = usage
	} else if usage == "" {
		return fmt.Errorf("Usage config value is null")
	}

	if version := conf.String("version"); version != "" {
		Version = version
	} else if version == "" {
		return fmt.Errorf("Version config value is null")
	}

	if author := conf.String("author"); author != "" {
		Author = author
	} else if author == "" {
		return fmt.Errorf("Author config value is null")
	}

	if email := conf.String("email"); email != "" {
		Email = email
	} else if email == "" {
		return fmt.Errorf("Email config value is null")
	}

	return nil
}

func setServerConfig(conf config.Configer) error {
	//config runtime
	if runmode := conf.String("runmode"); runmode != "" {
		RunMode = runmode
	} else if runmode == "" {
		return fmt.Errorf("RunMode config value is null")
	}

	if listenmode := conf.String("listenmode"); listenmode != "" {
		ListenMode = listenmode
	} else if listenmode == "" {
		return fmt.Errorf("ListenMode config value is null")
	}

	if httpscertfile := conf.String("httpscertfile"); httpscertfile != "" {
		HTTPSCertFile = httpscertfile
	} else if httpscertfile == "" {
		return fmt.Errorf("HttpsCertFile config value is null")
	}

	if httpskeyfile := conf.String("httpskeyfile"); httpskeyfile != "" {
		HTTPSKeyFile = httpskeyfile
	} else if httpskeyfile == "" {
		return fmt.Errorf("HttpsKeyFile config value is null")
	}

	if logpath := conf.String("log::filepath"); logpath != "" {
		LogPath = logpath
	} else if logpath == "" {
		return fmt.Errorf("LogPath config value is null")
	}

	if loglevel := conf.String("log::level"); loglevel != "" {
		LogLevel = loglevel
	} else if loglevel == "" {
		return fmt.Errorf("LogLevel config value is null")
	}

	if databasedriver := conf.String("database::driver"); databasedriver != "" {
		DatabaseDriver = databasedriver
	} else if databasedriver == "" {
		return fmt.Errorf("Database Driver config value is null")
	}

	if databaseuri := conf.String("database::uri"); databaseuri != "" {
		DatabaseURI = databaseuri
	} else if databaseuri == "" {
		return fmt.Errorf("Database URI config vaule is null")
	}

	// Deployment domain could be empty
	if domains := conf.String("deployment::domains"); domains != "" {
		Domains = domains
	} else if domains == "" {
		return fmt.Errorf("Deployment domains config vaule is null")
	}

	//TODO: Add a config option for provide Docker Registry V1.
	//TODO: Link @middle/header/setRespHeaders, @handler/dockerv1/-functions.
	if standalone := conf.String("dockerv1::standalone"); standalone != "" {
		DockerStandalone = standalone
	} else if standalone == "" {
		return fmt.Errorf("DockerV1 standalone value is null")
	}

	if registry := conf.String("dockerv1::version"); registry != "" {
		DockerRegistryVersion = registry
	} else if registry == "" {
		return fmt.Errorf("DockerV1 Registry Version value is null")
	}

	if storage := conf.String("dockerv1::storage"); storage != "" {
		DockerV1Storage = storage
	} else if storage == "" {
		return fmt.Errorf("DockerV1 Storage value is null")
	}

	if distribution := conf.String("dockerv2::distribution"); distribution != "" {
		DockerDistributionVersion = distribution
	} else if distribution == "" {
		return fmt.Errorf("DockerV2 Distribution Version value is null")
	}

	if storage := conf.String("dockerv2::storage"); storage != "" {
		DockerV2Storage = storage
	} else if storage == "" {
		return fmt.Errorf("DockerV2 Storage value is null")
	}

	if storage := conf.String("appc::storage"); storage != "" {
		AppcStorage = storage
	} else if storage == "" {
		return fmt.Errorf("Appc Storage value is null")
	}

	//config update service
	if uskeymanager := conf.String("updateserver::keymanager"); uskeymanager != "" {
		KeyManager = uskeymanager
	} else if uskeymanager == "" {
		return fmt.Errorf("Update Server Key manager config value is null")
	}

	if usstorage := conf.String("updateserver::storage"); usstorage != "" {
		Storage = usstorage
	} else if usstorage == "" {
		return fmt.Errorf("Update Server Storage config value is null")
	}

	//scan content
	if scanKey := conf.String("scancontent:scanKey"); scanKey != "" {
		ScanKey = scanKey
	} else if scanKey == "" {
		//auto-generate the scan key if it is empty
		//WARNING: if the dockyard server restarts, user received data like callbackID will be useless.
		var key fernet.Key
		key.Generate()
		ScanKey = string(key.Encode())
	}

	return nil
}
