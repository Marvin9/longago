
<img src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/illustrations/svg/19.svg" height="100" />

## atlan-collect-assessment

> Go Start - Stop - Pause - Resume Looooooong upload tasks.


<br/>

### Quick start

Setup ```.env``` file
```
UPLOAD_STORAGE=/full/path/where/you/want/to/store/uploaded/files

#RECOMMENDED_PATH: ${full_path_of_current_directory}/tmp/uploads
```

- Run using docker

    ```make docker-compose-up```

    OR

    ```make storage && docker-compose up```

- Run manually

    ```make check```

    OR

    ```make storage && go run main.go```

The API is now available on your host at http://localhost:8000.

[API Documentation.](/docs/API.md)