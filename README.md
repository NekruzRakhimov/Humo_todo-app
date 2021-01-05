# Simple todo-app based on rest-api
This todo-app has pure architecture. Because of having dependency injection it is easy to add new expansions. This todo-app is divided into 3 parts and every of them are responsible for different cases.
> **Handlers** - this part of app is responsible for listening  and serving HTTP requests

> **Service** - this part of app is responsible for releasing logical part 

> **Repository** - this part of app is responsible for working with DB

Also this app has principle of auto migration. Following directories serve for: 
> ddl - creating tables
> dml - inserting some modifications
> dbinit - functions for initializing and dropping tables

Database used in this app is - SQLite

I'm glad to present you my app!
>>        Nekruz Rakhimov - Humo Academy