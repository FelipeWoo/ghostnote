# ghostnote

Minimal CLI for invisible notes.

Create markdown notes in your current folder (hidden), while indexing them in a central vault with git history.

## Features
- create, read, list, edit, delete notes
- YAML metadata (created_at, updated_at, tags, links)
- local `.ghostnotes/` (no context pollution)
- central vault index
- optional local git tracking

## Usage
```bash
ghostnote init
ghostnote new "fix nginx issue"
ghostnote list
ghostnote read <id>
ghostnote edit <id>
ghostnote delete <id>
```

## Philosophy

>Notes should exist without interrupting your workflow.