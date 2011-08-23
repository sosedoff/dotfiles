# Terminal options
#
if [[ `whoami` == "root" ]]; then
  PS1="# "
else
  PS1="$ "
fi

# Shell aliases
#
alias ls='ls -G'
alias la='ls -la'
alias '..'='cd .. && echo "--> now at: $(pwd)"'
alias copyssh='cat ~/.ssh/id_rsa.pub | pbcopy && echo "Public key was copied into the clipboard."'

# RVM Stuff
#
[[ -s "$HOME/.rvm/scripts/rvm" ]] && . "$HOME/.rvm/scripts/rvm"

alias 186='rvm use 1.8.6'
alias 187='rvm use 1.8.7'
alias 192='rvm use 1.9.2'
alias ree='rvm use ree'

# Rails stuff
#
alias rs="rails server thin"
alias rc="rails console"