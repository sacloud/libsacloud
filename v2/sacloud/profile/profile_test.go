// Copyright 2016-2021 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !windows
// +build !windows

package profile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func init() {
	os.Setenv(DirectoryNameEnv, "/tmp/.usacloud") // nolint
}

func Test_baseDir(t *testing.T) {
	os.Unsetenv(DirectoryNameEnv) // nolint
	defer func() {
		os.Setenv(DirectoryNameEnv, "/tmp/.usacloud") // nolint
	}()

	t.Run(fmt.Sprintf("without %s env", DirectoryNameEnv), func(t *testing.T) {
		homeDir, err := homedir.Dir()
		require.NoError(t, err)

		baseDir, err := baseDir()
		require.NoError(t, err)

		require.EqualValues(t, homeDir, baseDir)
	})

	t.Run(fmt.Sprintf("with %s env", DirectoryNameEnv), func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			testDir := "/test"
			err := os.Setenv(DirectoryNameEnv, testDir)
			require.NoError(t, err)

			baseDir, err := baseDir()
			require.NoError(t, err)

			require.EqualValues(t, testDir, baseDir)
		})

		t.Run("redundant path", func(t *testing.T) {
			testDir := "/test/../test1/../test"
			expect := "/test"

			err := os.Setenv(DirectoryNameEnv, testDir)
			require.NoError(t, err)

			baseDir, err := baseDir()
			require.NoError(t, err)

			require.EqualValues(t, expect, baseDir)
		})
		t.Run("Invalid env var", func(t *testing.T) {
			// - with filepath.ListSeparator
			testDir := "/test" + string([]rune{filepath.ListSeparator}) + "test"
			err := os.Setenv(DirectoryNameEnv, testDir)
			require.NoError(t, err)

			_, err = baseDir()
			require.Error(t, err)
		})
	})
}

func Test_ConfigFilePath(t *testing.T) {
	dir := os.Getenv(DirectoryNameEnv)

	expects := []struct {
		profileName string
		filePath    string
	}{
		{
			profileName: "default",
			filePath:    "/.usacloud/default/config.json",
		},
		{
			profileName: "test1",
			filePath:    "/.usacloud/test1/config.json",
		},
		{
			profileName: "",
			filePath:    "/.usacloud/default/config.json",
		},
	}

	t.Run("Valid profiles", func(t *testing.T) {
		for _, expect := range expects {
			path, err := ConfigFilePath(expect.profileName)
			require.NoError(t, err)

			p := strings.Replace(path, dir, "", 1)
			require.EqualValues(t, expect.filePath, p)
		}
	})

	t.Run("Invalid profiles", func(t *testing.T) {
		// - with filepath.Separator
		_, err := ConfigFilePath("test" + string([]rune{filepath.Separator}) + "test")
		require.Error(t, err)
		// - with filepath.ListSeparator
		_, err = ConfigFilePath("test" + string([]rune{filepath.ListSeparator}) + "test")
		require.Error(t, err)
	})
}

type loadConfigExpects struct {
	profileName string
	isValid     bool
	body        string
}

func testTargetProfiles() []loadConfigExpects {
	return []loadConfigExpects{
		{
			profileName: "default",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "default", "default"),
		},
		{
			profileName: "for-usacloud-unit-test1",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "for-usacloud-unit-test1", "for-usacloud-unit-test1"),
		},
		{
			profileName: "for-usacloud-unit-test2",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "for-usacloud-unit-test2", "for-usacloud-unit-test2"),
		},
		{
			profileName: " for-usacloud-unit-test3\n\n",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "for-usacloud-unit-test3", "for-usacloud-unit-test3"),
		},
		{
			profileName: "invalid-json",
			isValid:     false,
			body:        "{",
		},
		{
			profileName: "empty-body",
			isValid:     false,
			body:        "",
		},
	}
}

func Test_Load(t *testing.T) {
	defer initConfigFiles()()

	t.Run("Valid profiles", func(t *testing.T) {
		for _, prof := range testTargetProfiles() {
			conf := &ConfigValue{}
			err := Load(prof.profileName, conf)
			if prof.isValid {
				require.NoError(t, err)
				pname := cleanupProfileName(prof.profileName)
				require.EqualValues(t, pname, conf.AccessToken)
				require.EqualValues(t, pname, conf.AccessTokenSecret)
			} else {
				require.Error(t, err)
			}
		}
	})

	t.Run("Not exists profile", func(t *testing.T) {
		// not exists profile
		conf := &ConfigValue{}
		require.Error(t, Load("not-exists-profile-name", conf))
	})

	t.Run("Invalid profile names", func(t *testing.T) {
		conf := &ConfigValue{}

		// - with filepath.Separator
		require.Error(t, Load("test"+string([]rune{filepath.Separator})+"test", conf))

		// - with filepath.ListSeparator
		require.Error(t, Load("test"+string([]rune{filepath.ListSeparator})+"test", conf))
	})
}

func Test_LoadWithExtendedConfig(t *testing.T) {
	initFunc := func() func() {
		cleanupAllProfiles()
		return func() {
			cleanupAllProfiles()
		}
	}
	defer initFunc()()

	val := &extendedConfigValue{
		ConfigValue: ConfigValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
		},
		Added: "added",
	}

	err := Save("test-profile-extended-config", val)
	require.NoError(t, err)

	extended := &extendedConfigValue{}
	err = Load("test-profile-extended-config", extended)
	require.NoError(t, err)
	require.Equal(t, "added", extended.Added)

	// load as ConfigValue
	configValue := &ConfigValue{}
	err = Load("test-profile-extended-config", configValue)
	require.NoError(t, err)
	require.Equal(t, extended.AccessToken, configValue.AccessToken)
	require.Equal(t, extended.AccessTokenSecret, configValue.AccessTokenSecret)
}

func initConfigFiles() func() {
	p := "/tmp/.usacloud"
	os.MkdirAll(p, 0700)           // nolint
	os.Setenv(DirectoryNameEnv, p) // nolint

	for _, prof := range testTargetProfiles() {
		p, _ := ConfigFilePath(prof.profileName)
		if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			panic(err)
		}
		if err := os.WriteFile(p, []byte(prof.body), 0600); err != nil {
			panic(err)
		}
	}

	return func() {
		for _, prof := range testTargetProfiles() {
			p, _ := ConfigFilePath(prof.profileName)
			os.RemoveAll(filepath.Dir(p)) // nolint
		}
	}
}

var confTemplate = `
{
        "AccessToken": "%s",
        "AccessTokenSecret": "%s"
}`

func Test_CurrentName(t *testing.T) {
	homeDir, err := baseDir()
	require.NoError(t, err)

	configDir := filepath.Join(homeDir, configDirName)
	profNameFile := filepath.Join(homeDir, configDirName, currentFileName)

	os.Mkdir(configDir, 0755) // nolint
	os.Remove(profNameFile)   // nolint
	defer func() {
		os.Remove(profNameFile) // nolint
	}()

	t.Run("Should use default", func(t *testing.T) {
		n, err := CurrentName()
		require.NoError(t, err)
		require.Equal(t, "default", n)
	})

	t.Run("Should use profile file", func(t *testing.T) {
		// create profile name
		if err := os.WriteFile(profNameFile, []byte("usacloud-unit-test1"), 0600); err != nil {
			panic(err)
		}
		n, err := CurrentName()
		require.NoError(t, err)
		require.Equal(t, "usacloud-unit-test1", n)
	})

	t.Run("Invalid name in profile file", func(t *testing.T) {
		// - with filepath.Separator
		if err := os.WriteFile(profNameFile, []byte("test"+string([]rune{filepath.Separator})+"test"), 0600); err != nil {
			panic(err)
		}
		_, err := CurrentName()
		require.Error(t, err)

		// - with filepath.ListSeparator
		if err := os.WriteFile(profNameFile, []byte("test"+string([]rune{filepath.ListSeparator})+"test"), 0600); err != nil {
			panic(err)
		}
		_, err = CurrentName()
		require.Error(t, err)
	})
}

func Test_SetCurrentName(t *testing.T) {
	homeDir, err := baseDir()
	require.NoError(t, err)

	configDir := filepath.Join(homeDir, configDirName)
	profNameFile := filepath.Join(homeDir, configDirName, currentFileName)

	os.Mkdir(configDir, 0755) // nolint
	os.Remove(profNameFile)   // nolint
	defer func() {
		os.Remove(profNameFile) // nolint
	}()

	t.Run("Default profile", func(t *testing.T) {
		// profile dir isnot exists
		configFilePath, err := ConfigFilePath("default")
		require.NoError(t, err)
		profileDirExists := false
		if _, err := os.Stat(configFilePath); err == nil {
			profileDirExists = true
		}
		require.False(t, profileDirExists)

		err = SetCurrentName("default")
		require.NoError(t, err)

		data, err := os.ReadFile(profNameFile)
		require.NoError(t, err)
		require.Equal(t, "default", string(data))
	})

	t.Run("Exists profile", func(t *testing.T) {
		defer initConfigFiles()()

		err = SetCurrentName("for-usacloud-unit-test1")
		require.NoError(t, err)

		data, err := os.ReadFile(profNameFile)
		require.NoError(t, err)
		require.Equal(t, "for-usacloud-unit-test1", string(data))
	})

	t.Run("Not exists profile", func(t *testing.T) {
		defer initConfigFiles()()

		err := SetCurrentName("for-usacloud-unit-test1")
		require.NoError(t, err)

		err = SetCurrentName("not-exists")
		require.Error(t, err)

		data, err := os.ReadFile(profNameFile)
		require.NoError(t, err)
		require.Equal(t, "for-usacloud-unit-test1", string(data))
	})

	t.Run("Invalid name ", func(t *testing.T) {
		// - with filepath.Separator
		err := SetCurrentName("test" + string([]rune{filepath.Separator}) + "test")
		require.Error(t, err)

		// - with filepath.ListSeparator
		err = SetCurrentName("test" + string([]rune{filepath.ListSeparator}) + "test")
		require.Error(t, err)
	})
}

type extendedConfigValue struct {
	ConfigValue
	Added string
}

func Test_Save(t *testing.T) {
	testProfileName := "for-usacloud-unit-test1"
	cleanupProfile(testProfileName)
	defer cleanupProfile(testProfileName)

	fileExists := func(path string) bool {
		_, err := os.Stat(path)
		return err == nil
	}

	t.Run("Valid profile", func(t *testing.T) {
		defer cleanupProfile(testProfileName)

		val := &ConfigValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
		}

		err := Save(testProfileName, val)
		require.NoError(t, err)

		path, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)

		require.True(t, fileExists(filepath.Dir(path)))
		require.True(t, fileExists(path))
	})
	t.Run("Extended config value", func(t *testing.T) {
		defer cleanupProfile(testProfileName)

		val := &extendedConfigValue{
			ConfigValue: ConfigValue{
				AccessToken:       "test-token",
				AccessTokenSecret: "test-secret",
			},
			Added: "added",
		}

		err := Save(testProfileName, val)
		require.NoError(t, err)

		path, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)

		require.True(t, fileExists(filepath.Dir(path)))
		require.True(t, fileExists(path))
	})
	t.Run("Merge extended config value when saving config value", func(t *testing.T) {
		defer cleanupProfile(testProfileName)

		val := &extendedConfigValue{
			ConfigValue: ConfigValue{
				AccessToken:       "test-token",
				AccessTokenSecret: "test-secret",
			},
			Added: "added",
		}
		err := Save(testProfileName, val)
		require.NoError(t, err)

		val2 := &ConfigValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
		}
		err = Save(testProfileName, val2)
		require.NoError(t, err)

		targetFile, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)

		data, err := os.ReadFile(targetFile)
		require.NoError(t, err)

		var mapData map[string]interface{}
		err = json.Unmarshal(data, &mapData)
		require.NoError(t, err)

		// check extended fields
		got, ok := mapData["Added"]
		require.True(t, ok)
		require.Equal(t, "added", got.(string))
	})
	t.Run("Invalid profile name", func(t *testing.T) {
		defer cleanupProfile(testProfileName)

		val := &ConfigValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
		}
		// - with filepath.Separator
		profileName := "test" + string([]rune{filepath.Separator}) + "test"

		err := Save(profileName, val)
		require.Error(t, err)

		path, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)

		require.False(t, fileExists(path))

		// - with filepath.ListSeparator
		profileName = "test" + string([]rune{filepath.ListSeparator}) + "test"
		err = Save(profileName, val)
		require.Error(t, err)

		path, err = ConfigFilePath(testProfileName)
		require.NoError(t, err)

		require.False(t, fileExists(path))
	})
}

func Test_Remove(t *testing.T) {
	testProfileName := "for-usacloud-unit-test1"

	initFunc := func() func() {
		val := &ConfigValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
		}

		err := Save(testProfileName, val)
		if err != nil {
			panic(err)
		}

		err = SetCurrentName(testProfileName)
		if err != nil {
			panic(err)
		}

		return func() {
			cleanupProfile(testProfileName)
		}
	}
	fileExists := func(path string) bool {
		_, err := os.Stat(path)
		return err == nil
	}
	t.Run("Profile exists", func(t *testing.T) {
		defer initFunc()()
		err := Remove(testProfileName)

		require.NoError(t, err)

		path, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)

		require.False(t, fileExists(filepath.Dir(path)))
		require.False(t, fileExists(path))

		current, err := CurrentName()
		require.NoError(t, err)
		require.EqualValues(t, DefaultProfileName, current)
	})
	t.Run("Profile exists with other file", func(t *testing.T) {
		defer initFunc()()

		// create file in ProfileDir
		path, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)

		testOtherFile := filepath.Join(filepath.Dir(path), "test")
		err = os.WriteFile(testOtherFile, []byte{}, 0600)
		if err != nil {
			panic(err)
		}

		err = Remove(testProfileName)

		require.NoError(t, err)

		require.True(t, fileExists(filepath.Dir(path)))
		require.True(t, fileExists(testOtherFile))
		require.False(t, fileExists(path))
	})
	t.Run("Profile not exists", func(t *testing.T) {
		defer initFunc()()
		err := Remove("NotExistsProfileName")

		require.Error(t, err)

		path, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)

		require.True(t, fileExists(filepath.Dir(path)))
		require.True(t, fileExists(path))

		current, err := CurrentName()
		require.NoError(t, err)
		require.EqualValues(t, testProfileName, current)
	})
	t.Run("Invalid profile name", func(t *testing.T) {
		defer initFunc()()

		// - with filepath.Separator
		profileName := "test" + string([]rune{filepath.Separator}) + "test"

		err := Remove(profileName)
		require.Error(t, err)

		path, err := ConfigFilePath(testProfileName)
		require.NoError(t, err)
		require.True(t, fileExists(filepath.Dir(path)))
		require.True(t, fileExists(path))

		current, err := CurrentName()
		require.NoError(t, err)
		require.EqualValues(t, testProfileName, current)

		// - with filepath.ListSeparator
		profileName = "test" + string([]rune{filepath.ListSeparator}) + "test"
		err = Remove(profileName)
		require.Error(t, err)

		path, err = ConfigFilePath(testProfileName)
		require.NoError(t, err)
		require.True(t, fileExists(filepath.Dir(path)))
		require.True(t, fileExists(path))

		current, err = CurrentName()
		require.NoError(t, err)
		require.EqualValues(t, testProfileName, current)
	})
}

func cleanupProfile(profile string) {
	path, err := ConfigFilePath(profile)
	if err != nil {
		panic(err)
	}
	os.RemoveAll(filepath.Dir(path))
}

// cleanupAllProfiles remove all entries under profile-base-dir(includes "current" file)
func cleanupAllProfiles() {
	dir, err := baseDir()
	if err != nil {
		panic(err)
	}

	// dir is exists?
	configDirPath := filepath.Join(dir, configDirName)
	if _, err := os.Stat(configDirPath); err == nil {
		err := os.RemoveAll(configDirPath)
		if err != nil {
			panic(err)
		}
	}
}

func TestList(t *testing.T) {
	initFunc := func() func() {
		cleanupAllProfiles()
		return func() {
			cleanupAllProfiles()
		}
	}
	defer initFunc()()

	t.Run("Default only", func(t *testing.T) {
		defer initFunc()()

		profiles, err := List()
		require.NoError(t, err)
		require.Len(t, profiles, 1)
		require.EqualValues(t, DefaultProfileName, profiles[0])
	})
	t.Run("Multiple profile", func(t *testing.T) {
		defer initFunc()()

		// create profile
		testProfileNames := []string{"test2", "test1"}
		val := &ConfigValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
		}
		for _, n := range testProfileNames {
			err := Save(n, val)
			if err != nil {
				panic(err)
			}
		}

		profiles, err := List()
		require.NoError(t, err)
		require.Len(t, profiles, 3) // default + test1 + test2
		require.EqualValues(t, DefaultProfileName, profiles[0])
		require.EqualValues(t, "test1", profiles[1]) // sorted by name(except "default")
		require.EqualValues(t, "test2", profiles[2])
	})
	t.Run("With invalid profile", func(t *testing.T) {
		defer initFunc()()
		defer initConfigFiles()()

		profiles, err := List()
		require.NoError(t, err)

		targets := testTargetProfiles()
		validCount := 0
		for _, p := range targets {
			if p.isValid {
				validCount++
			}
		}

		require.Len(t, profiles, validCount)
	})
	t.Run("With empty dir", func(t *testing.T) {
		defer initFunc()()

		// create profile
		testProfileNames := []string{"test2", "test1"}
		val := &ConfigValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
		}
		for _, n := range testProfileNames {
			err := Save(n, val)
			if err != nil {
				panic(err)
			}
		}

		// create empty dir
		dir, _ := baseDir()
		err := os.MkdirAll(filepath.Join(dir, configDirName, "test3"), 0755)
		if err != nil {
			panic(err)
		}

		profiles, err := List()
		require.NoError(t, err)
		require.Len(t, profiles, 3) // default + test1 + test2 ( without test3 )
		require.EqualValues(t, DefaultProfileName, profiles[0])
		require.EqualValues(t, "test1", profiles[1]) // sorted by name(except "default")
		require.EqualValues(t, "test2", profiles[2])
	})
}

func TestConfigValue_TraceMode(t *testing.T) {
	cases := []struct {
		in                *ConfigValue
		expectHTTPEnabled bool
		expectAPIEnabled  bool
	}{
		{
			in:                &ConfigValue{},
			expectHTTPEnabled: false,
			expectAPIEnabled:  false,
		},
		{
			in:                &ConfigValue{TraceMode: " "},
			expectHTTPEnabled: false,
			expectAPIEnabled:  false,
		},
		{
			in:                &ConfigValue{TraceMode: "api"},
			expectHTTPEnabled: false,
			expectAPIEnabled:  true,
		},
		{
			in:                &ConfigValue{TraceMode: "API"},
			expectHTTPEnabled: false,
			expectAPIEnabled:  true,
		},
		{
			in:                &ConfigValue{TraceMode: "aPi"},
			expectHTTPEnabled: false,
			expectAPIEnabled:  true,
		},
		{
			in:                &ConfigValue{TraceMode: "http"},
			expectHTTPEnabled: true,
			expectAPIEnabled:  false,
		},
		{
			in:                &ConfigValue{TraceMode: "HTTP"},
			expectHTTPEnabled: true,
			expectAPIEnabled:  false,
		},
		{
			in:                &ConfigValue{TraceMode: "HtTp"},
			expectHTTPEnabled: true,
			expectAPIEnabled:  false,
		},
		{
			in:                &ConfigValue{TraceMode: "1"},
			expectHTTPEnabled: true,
			expectAPIEnabled:  true,
		},
	}

	for _, tc := range cases {
		require.Equal(t, tc.expectHTTPEnabled, tc.in.EnableHTTPTrace(), tc.in)
		require.Equal(t, tc.expectAPIEnabled, tc.in.EnableAPITrace(), tc.in)
	}
}
