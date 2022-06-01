package get_users

import (
	"fmt"
	"os/exec"
	"strings"
)

var global_ignore = []string{
	"daemon",
	"nobody",
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func user_filter(in_string_slice []string) []string {
	var output_slice []string
	for i := 0; i < len(in_string_slice); i++ {
		if (!strings.HasPrefix(string(in_string_slice[i]), "_")) &&
			(!contains(global_ignore, in_string_slice[i])) &&
			in_string_slice[i] != "" {
			output_slice = append(output_slice, strings.ReplaceAll(in_string_slice[i], " ", ""))
		}
	}
	return output_slice
}

func exec_shell(process_path string, kwargs []string) []string {
	out, err := exec.Command(process_path, kwargs...).Output()
	if err != nil {
		fmt.Println(err)
	}
	out_slice := strings.Split(string(out), "\n")
	// cleaned_output := user_filter(out_slice)
	// fmt.Println(cleaned_output)
	return user_filter(out_slice)
}

func get_process_path(process_name string) string {
	out, err := exec.Command("which", process_name).Output()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSuffix(string(out[:]), "\n")
}

func Get_Users() []string {
	cmd_args := []string{".", "-ls", "/Users"}
	return exec_shell(get_process_path("dscl"), cmd_args)
}
