package sincedb

import (
    "os"
    "fmt"
    "path/filepath"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"

)

const Namespace = "plugin.client.sincedb"

type sinceDB struct {
    log log.Log
}

func (r *sinceDB) Init() error {
	// The registry file is opened in the data path
	r.registryFile = paths.Resolve(paths.Data, r.registryFile)

	// Create directory if it does not already exist.
	registryPath := filepath.Dir(r.registryFile)
	err := os.MkdirAll(registryPath, 0750)
	if err != nil {
		return fmt.Errorf("Failed to created registry file dir %s: %v", registryPath, err)
	}

	// Check if files exists
	fileInfo, err := os.Lstat(r.registryFile)
	if os.IsNotExist(err) {
		logp.Info("No registry file found under: %s. Creating a new registry file.", r.registryFile)
		// No registry exists yet, write empty state to check if registry can be written
		return r.writeRegistry()
	}
	if err != nil {
		return err
	}

	// Check if regular file, no dir, no symlink
	if !fileInfo.Mode().IsRegular() {
		// Special error message for directory
		if fileInfo.IsDir() {
			return fmt.Errorf("Registry file path must be a file. %s is a directory.", r.registryFile)
		}
		return fmt.Errorf("Registry file path is not a regular file: %s", r.registryFile)
	}

	logp.Info("Registry file set to: %s", r.registryFile)

	return nil
}

func (r *sinceDB) GetStates() []file.State {
	return r.states.GetStates()
}

func (r *sinceDB) load() error {
	f, err := os.Open(r.registryFile)
	if err != nil {
		return err
	}

	defer f.Close()

	logp.Info("Loading registrar data from %s", r.registryFile)

	decoder := json.NewDecoder(f)
	states := []file.State{}
	err = decoder.Decode(&states)
	if err != nil {
		return fmt.Errorf("Error decoding states: %s", err)
	}

	states = resetStates(states)
	r.states.SetStates(states)
	logp.Info("States Loaded from registrar: %+v", len(states))

	return nil
}

// resetStates sets all states to finished and disable TTL on restart
// For all states covered by a prospector, TTL will be overwritten with the prospector value
func resetStates(states []file.State) []file.State {
	for key, state := range states {
		state.Finished = true
		// Set ttl to -2 to easily spot which states are not managed by a prospector
		state.TTL = -2
		states[key] = state
	}
	return states
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    sincedb := &sinceDB{
        log: l,
    }

    sincedb.Init()

    if err := sincedb.load(); err != nil {
    	return nil, fmt.Errorf("Error loading state: %v", err)
	}

    return sincedb, nil
}

func (r *sinceDB) Sender(e event.Event) int {
    return state.Ok
}

func (r *sinceDB) Add() int {
    return state.Ok
}

func (r *sinceDB) Find() types.Object {
    return []byte("this is find")
}

func (r *sinceDB) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
