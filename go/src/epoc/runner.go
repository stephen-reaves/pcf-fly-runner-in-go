package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

func main() {

	fmt.Println("fly -t phoenix-staging sp -p install-product -c pipeline.yml -l params.yml -l common.yml")

	//opens baselines.yml

	b, err := ioutil.ReadFile("baselines.yml") // just pass the file name for baselines
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'

	s, err := fmt.Scanln(os.Stdin)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(s)

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	//prompt
	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

}
