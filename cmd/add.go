package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Add(letterPair, word string) {
	file, err := os.OpenFile(Filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening csv file %v\n", err)
		return
	}

	csvFile := csv.NewReader(file)
	data, err := csvFile.ReadAll()
	if err != nil {
		fmt.Printf("Error reading csv file %v\n", err)
		return
	}
	file.Close()
	
	firstLetter := letterPair[0]
	secondLetter := letterPair[1]

	index1 := int(firstLetter) - 97
	index2 := int(secondLetter) - 97

	data[index1][index2] = word

	file, err = os.Create(Filename)
	if err != nil {
		fmt.Printf("Error creating csv file %v\n", err)
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(data)
	if err != nil {
		fmt.Printf("Error writing to csv file %v\n", err)
		return
	}

	csvWriter.Flush()
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"add", "a"},
	Short:   "add a letter pair",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		pair := strings.ToLower(args[0])
		word := args[1]
		if len(pair) != 2 {
			fmt.Println("Letter pair provided is not two letters")
			return
		}
		Add(pair, word)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
