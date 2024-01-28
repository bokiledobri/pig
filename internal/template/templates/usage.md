
   Go to your projects directory, and run "pig project [module name] --type=[type of the application, cli is default].
   Currently supported types are web and cli. Next example will create a new folder named cow in the current directory, create an awesome Makefile,
   .gitignore file, setup a basic web server with a home handler, init a git repository if git binary is available and init a module named github.com/bokiledobri/cow
   if go binary is found

   pig project github.com/bokiledobri/cow --type=web
   
