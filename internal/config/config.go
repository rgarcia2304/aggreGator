package config 

import(
	"errors"
	"os"
	"path/filepath"
	"encoding/json"
	"io"
	"bytes"

)
type Config struct{
	Url string `json: "db_url"`
	Username string `json: "current_user_name"`
}

func getConfigFilePath() (string, error){
	//read the JSON File from the home directory 
	homePath, err := os.UserHomeDir()
	if err != nil{
		return "", errors.New("There is an issue getting the homedirectory")
	}

	//combine and clean this with the location of the file
	fileLocation := homePath + "/.gatorconfig.json"
	cleanedPath := filepath.Clean(fileLocation)
	
	if err != nil{
		return "", errors.New("Location Was Not Created Properly")
	}
	
	return cleanedPath, nil
}

func (cfg *Config) Write() (error){
	configFilePath, err := getConfigFilePath()
	if err != nil{
		return err
	}

	f, err := os.Create(configFilePath)

	if err != nil{
		return err
	}

	defer f.Close()

	r, err := json.Marshal(cfg)
	
	if err != nil{
		return err
	}

	_, err = io.Copy(f, bytes.NewReader(r))
	return err
}

func Read() (Config, error){
	configFilePath, err := getConfigFilePath()
	
	if err != nil{
		return Config{}, err
	}

	f, err := os.Open(configFilePath)
	
	if err != nil{
		return Config{}, errors.New("There was an error loading the file")
	}	

	defer f.Close()
	
	byteValue, err := io.ReadAll(f)
	if err != nil{
		return Config{}, errors.New("There was an issue reading the file into bytes")
	}

	//Now unmarshal the data into the config struct 
	cfg := Config{} 
	
	if err := json.Unmarshal(byteValue, &cfg); err != nil{
		return Config{}, errors.New("Could not Unmarshall Data into Config Struct")
	}

	return cfg, nil
}

func (cfg *Config) SetUser(user string) error{
	cfg.Username = user
	err := cfg.Write()
	if err != nil{
		return errors.New("There has been an issue writing")
	}

	return err
}


