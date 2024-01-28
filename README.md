# Pig Is Generator

## 1. Instalation
   Clone or download this repository:
    ```bash
    git clone https://github.com/bokiledobri/pig.git
    ```
   From the project folder run "go install". The example below will install it in the /$HOME/local/bin:

    ```bash
    cd pig
    go install -o=/$HOME/.local/bin/pig ./cmd/cli
    ```
    
   Next command will install it in the /usr/local/bin:

    ```bash
    sudo make install
    ```

   Feel free to open [Makefile](https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects) in your favorite editor and edit make commands

## 2. Usage
### 1. Generate a project
   Go to your projects directory, and run "pig project [module name] --type=[type of the application, cli is default].
   Currently supported types are web and cli. Next example will create a new folder named cow in the current directory, create an awesome [Makefile](https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects),
   .gitignore file, setup a basic web server with a home handler, init a git repository if git binary is available and init a module named github.com/bokiledobri/cow
   if go binary is found
    ```bash
     pig project github.com/bokiledobri/cow --type=web
     ```
