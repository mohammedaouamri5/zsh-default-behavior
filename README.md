
---

# BRUH! This Is The Stupidest Thing I've Ever Done In My Life  
---
**NOTE**: There is an executable in the tool.

---

# Overview
This program functions as a command mapper, allowing you to create *programmable aliases*. \
It enables you to customize how commands are executed based pattern in them.


# Features
Here are some examples of what you can configure the program to do:

- **Git Repositories**: If it points to a Git repository, it will clone it. *That's the main motivation!*
- **Images**: For image files, it will run `chafa` to display the image directly in your terminal.
- **Directories**: If the input is a directory, the program will navigate to that location.
- **Websites**: If the command is a website, it will open your browser to that page.

The possibilities are endless! You can configure it for various commands and even create new commands or scripts. \
If you’re a wizard, you can generate scripts on the fly.




## Installation   
### With Oh-my-zsh 

After running:
```zsh
git clone https://github.com/mohammedaouamri5/zsh-default-behavior.git
# You might need sudo
cp -rfv zsh-default-behavior $ZSH_CUSTOM/plugins/     
```

Then in `~/.zshrc`, add the plugin:
```zsh
plugins=(
  zsh-default-behavior
) 
```

### Without Oh-my-zsh  
You need to make some tweaks. Clone it wherever you want, just remember where. You probably want to put it in `~/.config`:
```zsh
git clone https://github.com/mohammedaouamri5/zsh-default-behavior.git
```

Then go to `zsh-default-behavior.plugin.zsh` and change this block of code:

**FROM:**
```zsh
preexec () {
    LAST_COMMAND="$1"
    eval "$($HOME/DEV/tools/zsh-default-behavior/main $LAST_COMMAND)"
    kill -INT $$
    return 1 
}
```

**TO:**
```zsh
preexec () {
    LAST_COMMAND="$1"
    eval "$(~/.config/zsh-default-behavior/main $LAST_COMMAND)"
    kill -INT $$
    return 1 
}
```

Finally, source the plugin in `~/.zshrc`:
```zsh
source ~/.config/zsh-default-behavior/zsh-default-behavior.plugin.zsh
```

---

## Configuration 

1. Go to `Config.go`
2. Add one of the following:
    - `func (Start) MyFunc(command string) {}`  
    - `func (Start) MyFunc(command string, exit_code string) {}`
3. Make sure that you call `Run()`.
4. Compile the program : — yeah, I know, I’m one of those people who make you compile a program just to use a simple config.
5. 
```zsh
go build -o main .
```

---

## What's Going On Here? 
***So...*** It's crazy, I know.  
Yes, it's stupid, kind of efficient, unsafe, and doesn't even do the main task yet.

***BUT 🍑***  
- I really need something to handle this.
- I want this tool to enhance terminal workflows.
- It's fun to use GoLang and spawn goroutines like crazy.

**Real talk:**  
I've never built an Oh-my-zsh plugin or a Go concurrency tool before. I really hope to improve on it as I go.

---

# Development  
## TODO  
- [x] Get the command.
- [x] Run your own command.
- [x] Block the original command.
- [ ] Make it use the directory of the command instead of the config.
---

