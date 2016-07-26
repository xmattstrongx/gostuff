# Path to your oh-my-zsh installation.
export ZSH=/Users/matthewst/.oh-my-zsh

# Set name of the theme to load.
# Look in ~/.oh-my-zsh/themes/
# Optionally, if you set this to "random", it'll load a random theme each
# time that oh-my-zsh is loaded.
ZSH_THEME="agnoster"

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion. Case
# sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# The optional three formats: "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load? (plugins can be found in ~/.oh-my-zsh/plugins/*)
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(git chucknorris docker jsontools knife npm osx rvm sudo wd)

# User configuration

# export PATH="/Users/matthewst/.gvm/pkgsets/go1.5.3/global/bin:/Users/matthewst/.gvm/gos/go1.5.3/bin:/Users/matthewst/.gvm/pkgsets/go1.5.3/global/overlay/bin:/Users/matthewst/.gvm/bin:/Users/matthewst/.gvm/bin:/Users/matthewst/.rbenv/shims:/Users/matthewst/.rvm/gems/ruby-2.2.3/bin:/Users/matthewst/.rvm/gems/ruby-2.2.3@global/bin:/Users/matthewst/.rvm/rubies/ruby-2.2.3/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/Users/matthewst/.rbenv/shims:/Users/matthewst/.rvm/gems/ruby-2.2.3/bin:/Users/matthewst/.rvm/gems/ruby-2.2.3@global/bin:/Users/matthewst/.rvm/rubies/ruby-2.2.3/bin:/Users/matthewst/.rvm/bin:/usr/local/sbin:/Users/matthewst/Tools/GO/bin:/Users/matthewst/go/bin:/Users/matthewst/.rvm/bin:/usr/local/sbin:/Users/matthewst/Tools/GO/bin:/Users/matthewst/go/bin"
# export MANPATH="/usr/local/man:$MANPATH"

source $ZSH/oh-my-zsh.sh

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
if [[ -n $SSH_CONNECTION ]]; then
 export EDITOR='vim'
else
 export EDITOR='vim'
fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# ssh
# export SSH_KEY_PATH="~/.ssh/dsa_id"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
alias zshconfig="$EDITOR ~/.zshrc"
alias ohmyzsh="$EDITOR ~/.oh-my-zsh"
alias ..='cd ..'            # Go up one directory
alias ...='cd ../..'        # Go up two directories
alias ....='cd ../../..'    # And for good measure
alias l='ls -lahp'   # Long view, show hidden
alias la='ls -AFp'   # Compact view, show hidden
alias ll='ls -lFhp'  # Long view, no hidden

alias rake='noglob rake'

alias nlist='nova list'
alias ncount='nova list | grep net | wc -l'






###MATTS CUSTOM STUFF###
########SOURCE SCRIPTS##########
source ~/Desktop/identity.sh
######END SOURCE SCRIPTS########

########RVM EXPORTS######
#[[ -s "$HOME/.rvm/scripts/rvm" ]] &&
# eval "$(rbenv init - zsh)"
#
# eval export PATH=$PATH:"/Users/matthewst/.rbenv/shims:${PATH}"
# #source "/usr/local/Cellar/rbenv/0.4.0/libexec/../completions/rbenv.bash"
# rbenv rehash 2>/dev/null
# rbenv() {
#   typeset command
#   command="$1"
#   if [ "$#" -gt 0 ]; then
#     shift
#   fi
#
#   case "$command" in
#   rehash|shell)
#     eval `rbenv "sh-$command" "$@"`;;
#   *)
#     command rbenv "$command" "$@";;
#   esac
# }

export PATH="$PATH:$HOME/.rvm/bin" # Add RVM to PATH for scriptin
eval "$(rbenv init - zsh)"
eval export PATH="/Users/matthewst/.rbenv/shims:${PATH}"
source "/usr/local/Cellar/rbenv/0.4.0/libexec/../completions/rbenv.bash"
rbenv rehash 2>/dev/null
rbenv() {
  typeset command
  command="$1"
  if [ "$#" -gt 0 ]; then
    shift
  fi

  case "$command" in
  rehash|shell)
    eval `rbenv "sh-$command" "$@"`;;
  *)
    command rbenv "$command" "$@";;
  esac
}

# [[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*
#####END RVM EXPORTS#####

########GO EXPORTS######
#Set which version of go to use and export go project to GOPATH
source "/Users/matthewst/.gvm/scripts/gvm"
export GOPATH=~/go

#Set up go to build linux executables
export GOXCOMP_OS=linux
export GOXCOMP_ARCH=amd64

# enable PATH to find rabbitmqctl executable
export PATH=$PATH:/usr/local/sbin

# enable PATH to find mockery executable
export PATH=$PATH:~/Tools/GO/bin

# enable PATH to GORE (GO REPL)
export PATH=$PATH:$GOPATH/bin
#####END GO EXPORTS#####

####DOCKER ALIASES AND ETC##########
# If you think these program names are too long,
# especially useful for Compose and Machine.
alias dk='docker'
alias dcom='docker-compose'

# Start your local Docker VM.
# Only really applicable on OS X.
#alias dockstart='docker-machine start default'

#remove all untagged containers
alias dkpurge='docker rmi -f "$(docker images | grep "^<none>" | awk "{print $3}")"'

# This stops and removes all of the containers
# in a Docker Compose group.
# Useful for when you just want to start over fresh.
alias dcomtrash='docker-compose kill && docker-compose rm -f'

# If you always use your local machine, this will be
# very convenient. If you use different Docker machines,
# check out the whale function below.
# To get a better understanding of what this command
# does, see the guide on Docker Environment Variables.
# Basically, this command sets up your terminal to
# connect to your local Docker VM.
# If you're on Linux, this won't be as applicable to
# to you. You can just set your DOCKER_HOST env var
# directly in your bashrc/bash_profile as explained
# in the Linux setup instructions.
#eval "$(docker-machine env default)";

# If you use different Docker machines, such as
# your local one and remote ones, you might find it
# useful to switch between them quickly.
# When you execute whale with no arguments, it will
# set up your shell to connect to your default local
# machine. If you pass it an argument, it will set up
# your terminal to connect to the machine with that name.
#function whale {
#	if [ -z "$1" ]; then
#		eval "$(docker-machine env default)";
#	else
#		eval "$(docker-machine env $1)";
#	fi
#}
######END DOCKER ALIASES########

#########FUNCTIONS##########
#ssh shorthand function
function sshin {
  ssh -i ~/Desktop/pem/identity-v16-dev.pem -o StrictHostKeyChecking=no ubuntu@$@
}

# set DED GOPATH
ded-gopath() {
export GOPATH=~/ded/Product
}

# set identity GOPATH
identity-gopath() {
export GOPATH=~/identity/Product
}

# set entitlements GOPATH
entitlements-gopath() {
export GOPATH=~/src/ultimatesoftware.com/entitlements
}

# set entitlements GOPATH
notifications-gopath() {
export GOPATH=~/src/ultimatesoftware.com/notifications
}

resty() {
source ~/resty
}

set-ded-local() {
  CURRENTIPADDRESS=$(ifconfig | grep 'inet ' | grep -v 127.0.0.1 | cut -d\  -f2 | tail -1)
  export artifactory_url="http://${CURRENTIPADDRESS}:8000/"
  export version="0.0.1"
# export artifactory_url="http://10.50.214.31:8000/"
  export artifact_dir="../../artifacts/ded/0.0.1/"
  export GOXCOMP_OS=linux
  export GOXCOMP_ARCH=amd64
}

set-ginkgo() {
#Ginkgo needs the vendor path for outside executables and the internal project Product path
export GOPATH=/Users/matthewst/go/src/ultimatesoftware.com/identity/Product:/Users/matthewst/go/src/ultimatesoftware.com/identity/Product/src/ultimatesoftware.com/accountstore/_vendor:/Users/matthewst/go
}
########END FUNCTIONS#########

export PATH="$PATH:$HOME/.rvm/bin" # Add RVM to PATH for scripting

chuck | cowsay
