package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Get(letterPair string) {
	file, err := os.OpenFile(getFileName(), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening csv file: %v\n", err)
		return
	}

	csvFile := csv.NewReader(file)
	data, err := csvFile.ReadAll()
	if err != nil {
		fmt.Printf("Error reading csv file: %v\n", err)
		return
	}
	file.Close()

	firstLetter := letterPair[0]
	secondLetter := letterPair[1]

	index1 := int(firstLetter) - 96
	index2 := int(secondLetter) - 96

	fmt.Println(data[index1][index2])
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "get a letter pair",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pair := strings.ToLower(args[0])
		if len(pair) != 2 {
			fmt.Println("Letter pair provided is not two letters")
			return
		}
		Get(pair)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
