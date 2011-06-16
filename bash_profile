# Shell aliases
alias ls='ls -G'
alias la='ls -la'
alias '..'='cd ..'

# Copy pubkey
alias copyssh='cat ~/.ssh/id_rsa.pub | pbcopy'

# RVM Stuff

[[ -s "$HOME/.rvm/scripts/rvm" ]] && . "$HOME/.rvm/scripts/rvm"

alias 186='rvm use 1.8.6'
alias 187='rvm use 1.8.7'
alias ree='rvm use ree'
alias 192='rvm use 1.9.2'

# Rails stuff

alias rs='rails server thin'