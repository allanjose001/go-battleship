package service

import (
	"encoding/json"
	"os"
)

type Profile struct {
	Username string;
	TotalScore int;
	HighestScore int;
	GamesPlayed int;
	MedalsEarned int;
}

const defaultPath string = "internal/service/saves/profiles.json";

func SaveProfile(profile Profile) error {
    if _, err := os.Stat(defaultPath); err == nil { // arquivo existe, atualiza Json
        return UpdateProfile(profile)

    } else if !os.IsNotExist(err) { 
        return err
    }

    // arquivo não existe, cria um com o profile
    profiles := []Profile{profile}

    data, err := json.MarshalIndent(profiles, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(defaultPath, data, 0644)
}

func UpdateProfile(profile Profile) error {
    profiles, err := LoadProfiles()
    if err != nil {
        return err
    }

    updated := false
    for i, p := range profiles {
        if p.Username == profile.Username {
            profiles[i] = profile
            updated = true
            break
        }
    }

    if !updated {
        profiles = append(profiles, profile)
    }

    data, err := json.MarshalIndent(profiles, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(defaultPath, data, 0644)
}



func LoadProfiles() ([]Profile, error) {
	
	data, err := os.ReadFile(defaultPath)
    if err != nil {
		if os.IsNotExist(err) {
			return []Profile{}, nil // arquivo ainda não existe
        }
        return nil, err
    }

    var profiles []Profile
    err = json.Unmarshal(data, &profiles)
    if err != nil {
		return nil, err
    }
	
    return profiles, nil
}

