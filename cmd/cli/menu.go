package cli

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
)

func Menu(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			prompt := promptui.Select{
				Label: "Choose action",
				Items: []string{"Create note", "List notes", "Delete note"},
			}

			_, result, err := prompt.Run()
			if err != nil {
				return err
			}

			fmt.Println("Selected:", result)
			return nil

		}
	}

}
