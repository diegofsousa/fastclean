package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var op string
	var args []string
	binary, lookErr := exec.LookPath("sudo")

	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Printf("Deseja efetuar limpeza 'autoremove'?\n[y] para Sim [n] para Não: ")
	fmt.Scan(&op)

	if op == "Y" || op == "y" {
		fmt.Println("Efetuando limpeza...")
		args = []string{"sudo", "apt-get", "autoremove"}
		env := os.Environ()
		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			panic(execErr)
		}
	} else {
		fmt.Println("Limpeza cancelada")
	}
}
