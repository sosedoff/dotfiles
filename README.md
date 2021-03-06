# Dotfiles

Install dotfiles and binfiles:

```
rake install
```

Or separately:

```
rake dotfiles
rake binfiles
```

## Sublime Text 3

Few tasks to manage Sublime Text settings:

```
rake sublime:settings  # Install settings
rake sublime:themes    # Install custom themes
```

## Scripts and Tools

General purpose:

- `addr`          - Get current IP address of the machine
- `dmg`           - Quickly install DMG package
- `ramfs`         - Manage RAMFS volumes (OSX specific)
- `gem-version`   - Quickly check the latest gem version
- `genpassword`   - Generate a random password
- `static_server` - Start a static content server in current directory

Git extras:

- `feature`          - Start a new git branch or checkout existing.
- `master`           - Switch to master branch
- `develop`          - Switch to development branch
- `git track`        - Setup remote git branch tracking
- `git tarball`      - Package git HEAD into tar.gz archive
- `git branches`     - List all git branches and their diff stats with current branch
- `git merge-repo`   - Import another repos to a subdirectory of the current repo
- `git clean-merged` - Remove any local branches that were merged into master