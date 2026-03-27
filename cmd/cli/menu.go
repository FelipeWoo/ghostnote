package cli

import (
	"context"
	"fmt"

	"ghostnote/internal/logger"
	t "ghostnote/internal/transport/cli"

	"github.com/manifoldco/promptui"
)

func Menu(ctx context.Context, noteCommand *t.NoteCommand) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			shouldExit, err := mainMenu(noteCommand)
			if err != nil {
				return err
			}
			if shouldExit {
				return nil
			}
		}
	}
}

func mainMenu(noteCommand *t.NoteCommand) (bool, error) {
	prompt := promptui.Select{
		Label: "GhostNote Operations",
		Items: []string{"Create note", "List notes", "Delete note", "Exit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	switch result {
	case "Create note":
		return false, createMenu(noteCommand)
	case "List notes":
		return false, listMenu(noteCommand)
	case "Delete note":
		return false, deleteMenu(noteCommand)
	case "Exit":
		fmt.Println("Bye")
		return true, nil
	default:
		return false, nil
	}
}

func createMenu(noteCommand *t.NoteCommand) error {
	titlePrompt := promptui.Prompt{
		Label: "Note title",
	}
	title, err := titlePrompt.Run()
	if err != nil {
		return err
	}

	tagsPrompt := promptui.Prompt{
		Label: "Note tags, separated by comma",
	}
	tags, err := tagsPrompt.Run()
	if err != nil {
		return err
	}

	linksPrompt := promptui.Prompt{
		Label: "Note links, separated by comma",
	}
	links, err := linksPrompt.Run()
	if err != nil {
		return err
	}

	payloadPrompt := promptui.Prompt{
		Label: "Note payload",
	}
	payload, err := payloadPrompt.Run()
	if err != nil {
		return err
	}

	title, err = normalize(title)
	if err != nil {
		return err
	}

	payload = cleanPayload(payload)

	tagList := cleanCSV(tags)
	linkList := cleanCSV(links)

	err = noteCommand.Create(title, payload, tagList, linkList)
	if err != nil {
		return err
	}

	logger.Info("Note created: %s", title)
	return nil
}

func listMenu(noteCommand *t.NoteCommand) error {
	logger.Info("List notes not implemented yet")
	return nil
}

func deleteMenu(noteCommand *t.NoteCommand) error {
	logger.Info("Delete note not implemented yet")
	return nil
}
