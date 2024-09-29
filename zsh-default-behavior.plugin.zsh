echo "bruh"

#
# WARNING : !!! 
# IF YOU WANT TO BUILF AND EXEC WILL BE BETTER 
#
    # just build  
    # change the 
        # eval (go run $ZSH_CUSTOM/plugins/zsh-default-behavior $cmd ) 
        # to
        # eval ($ZSH_CUSTOM/plugins/zsh-default-behavior/<your_exec> $cmd ) 
#




# Hook that runs before a command is executed
preexec () {
    LAST_COMMAND="$1"
    eval "$($ZSH_CUSTOM/plugins/zsh-default-behavior/main $LAST_COMMAND)"
    kill -INT $$
    return  1 
}
# Hook that runs after the command is executed
precmd () {
    local exit_code=$?
}

autoload -Uz add-zsh-hook

m 

