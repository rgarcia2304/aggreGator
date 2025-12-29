package config 

import(
	"errors"
	"os"
	"path/filepath"
	"encoding/json"

)
type Config struct{
	Url string `json: "db_url"`
	Username string `json: "current_user_name"`
}

func getConfigFilePath() (string, error){
	//read the JSON File from the home directory 
	homePath := os.UserHomeDir()
	//combine and clean this with the location of the file
	fileLocation, err := homepath.Ext(".gatorconfig.json")
	
	if err != nil{
		return "", errors.New("Location Was Not Created Properly")
	}
	
	return fileLocation, nil
}

func (cfg *Config) Write(error){
	configFilePath := getConfigFilePath()
	f, err := os.Create(configFilePath)

	if err != nil{
		return err
	}

	defer f.Close()

	r, err := json.Marshal(cfg)
	
	if err != nil{
		return err
	}

	_, err = io.Copy(f, r)
	return err
}

func Read() (*Config, error){
	configFilePath, err := getConfigFilePath()
	
	if err != nil{
		return err
	}

	f, err := os.Open(configFilePath)
	
	if err != nil{
		return errors.New("There was an error loading the file")
	}	

	defer f.Close()

	//Now unmarshal the data into the config struct 
	cfg := Config{} 
	
	if err := json.Unmarshal(f, &cfg); err != nil{
		return Config{}, errors.New("Could not Unmarshall Data into Config Struct")
	}

	return &cfg, nil
}

func (cfg *Config) SetUser(string user) error{
	cfg.Username = user
	_, err := cfg.Write()
	return err
}


