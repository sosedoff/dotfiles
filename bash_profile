export PATH=$PATH:$HOME/bin

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

alias gs='git st'
alias gd='git diff'
alias gl='git log'
alias gp='git push'

# RVM Stuff
#
[[ -s "$HOME/.rvm/scripts/rvm" ]] && . "$HOME/.rvm/scripts/rvm"

alias 186='rvm use 1.8.6'
alias 187='rvm use 1.8.7'
alias 192='rvm use 1.9.2-p290'
alias 193='rvm use 1.9.2-rc1'
alias ree='rvm use ree'

# Some other dev-related aliases
#
alias b="bundle"
alias be="bundle exec"
alias bi="bundle install"
alias bu="bundle update"

alias r='rake'

# Rails stuff
#
alias rs="rails server thin"
alias rc="rails console"

# Other stuff
alias "please"="sudo"
